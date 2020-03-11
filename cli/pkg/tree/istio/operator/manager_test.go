package operator_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/testutils"
	"github.com/solo-io/mesh-projects/cli/pkg/cliconstants"
	mock_kube "github.com/solo-io/mesh-projects/cli/pkg/common/kube/mocks"
	"github.com/solo-io/mesh-projects/cli/pkg/options"
	"github.com/solo-io/mesh-projects/cli/pkg/tree/istio/operator"
	mock_operator "github.com/solo-io/mesh-projects/cli/pkg/tree/istio/operator/mocks"
	mock_server "github.com/solo-io/mesh-projects/cli/pkg/tree/version/server/mocks"
	"github.com/solo-io/mesh-projects/pkg/common/docker"
	mock_docker "github.com/solo-io/mesh-projects/pkg/common/docker/mocks"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/resource"
)

var _ = Describe("Cluster Operations", func() {
	var (
		ctrl        *gomock.Controller
		testErr     = eris.New("test-err")
		clusterName = "cluster-name"
	)

	var buildInstallConfigWithDefaults = func(createNamespace, createCRD bool) *options.IstioInstallationConfig {
		return &options.IstioInstallationConfig{
			CreateNamespace:            createNamespace,
			CreateIstioControlPlaneCRD: createCRD,
			InstallNamespace:           cliconstants.DefaultIstioOperatorNamespace,
			IstioOperatorVersion:       cliconstants.DefaultIstioOperatorVersion,
		}
	}

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("Install", func() {
		It("can install after setting defaults", func() {
			kubeClient := mock_kube.NewMockUnstructuredKubeClient(ctrl)
			manifestBuilder := mock_operator.NewMockInstallerManifestBuilder(ctrl)
			deploymentClient := mock_server.NewMockDeploymentClient(ctrl)
			imageNameParser := mock_docker.NewMockImageNameParser(ctrl)

			installConfig := &options.IstioInstallationConfig{
				CreateNamespace:            true,
				CreateIstioControlPlaneCRD: true,
				InstallNamespace:           "", // purposely left empty
				IstioOperatorVersion:       "", // purposely left empty
			}

			installWithDefaults := *installConfig
			installWithDefaults.InstallNamespace = cliconstants.DefaultIstioOperatorNamespace
			installWithDefaults.IstioOperatorVersion = cliconstants.DefaultIstioOperatorVersion

			installManifest := "hoooo boy let's install istio"
			manifestBuilder.EXPECT().Build(&installWithDefaults).Return(installManifest, nil)

			resources := []*resource.Info{
				{Name: "resource1"}, {Name: "resource2"},
			}

			kubeClient.EXPECT().BuildResources(cliconstants.DefaultIstioOperatorNamespace, installManifest).Return(resources, nil)
			kubeClient.EXPECT().Create(cliconstants.DefaultIstioOperatorNamespace, resources).Return(resources, nil)

			operatorManager := operator.NewManager(
				kubeClient,
				manifestBuilder,
				deploymentClient,
				imageNameParser,
				installConfig,
			)

			err := operatorManager.Install()
			Expect(err).NotTo(HaveOccurred())
		})

		It("reports the correct error if the install manifest is unparseable", func() {
			kubeClient := mock_kube.NewMockUnstructuredKubeClient(ctrl)
			manifestBuilder := mock_operator.NewMockInstallerManifestBuilder(ctrl)
			deploymentClient := mock_server.NewMockDeploymentClient(ctrl)
			imageNameParser := mock_docker.NewMockImageNameParser(ctrl)

			installConfig := buildInstallConfigWithDefaults(true, true)

			operatorManager := operator.NewManager(
				kubeClient,
				manifestBuilder,
				deploymentClient,
				imageNameParser,
				installConfig,
			)

			manifestBuilder.EXPECT().Build(installConfig).Return("", testErr)

			err := operatorManager.Install()
			Expect(err).To(testutils.HaveInErrorChain(operator.FailedToGenerateInstallManifest(testErr)))
		})

		It("reports the correct error if the manifest can't be turned into k8s resources", func() {
			kubeClient := mock_kube.NewMockUnstructuredKubeClient(ctrl)
			manifestBuilder := mock_operator.NewMockInstallerManifestBuilder(ctrl)
			deploymentClient := mock_server.NewMockDeploymentClient(ctrl)
			imageNameParser := mock_docker.NewMockImageNameParser(ctrl)

			installConfig := buildInstallConfigWithDefaults(true, true)

			operatorManager := operator.NewManager(
				kubeClient,
				manifestBuilder,
				deploymentClient,
				imageNameParser,
				installConfig,
			)

			installManifest := "hoooo boy let's install istio"
			manifestBuilder.EXPECT().Build(installConfig).Return(installManifest, nil)

			kubeClient.EXPECT().BuildResources(cliconstants.DefaultIstioOperatorNamespace, installManifest).Return(nil, testErr)

			err := operatorManager.Install()
			Expect(err).To(testutils.HaveInErrorChain(operator.FailedToParseInstallManifest(testErr)))
		})

		It("cleans up the operator install if part of it fails", func() {
			kubeClient := mock_kube.NewMockUnstructuredKubeClient(ctrl)
			manifestBuilder := mock_operator.NewMockInstallerManifestBuilder(ctrl)
			deploymentClient := mock_server.NewMockDeploymentClient(ctrl)
			imageNameParser := mock_docker.NewMockImageNameParser(ctrl)

			installConfig := buildInstallConfigWithDefaults(false, false)

			installWithDefaults := *installConfig
			installWithDefaults.InstallNamespace = cliconstants.DefaultIstioOperatorNamespace
			installWithDefaults.IstioOperatorVersion = cliconstants.DefaultIstioOperatorVersion

			installManifest := "hoooo boy let's install istio"
			manifestBuilder.EXPECT().Build(&installWithDefaults).Return(installManifest, nil)

			resources := []*resource.Info{
				{Name: "resource1"}, {Name: "resource2"},
			}
			successfulResources := resources[:1]

			kubeClient.EXPECT().BuildResources(cliconstants.DefaultIstioOperatorNamespace, installManifest).Return(resources, nil)
			kubeClient.EXPECT().Create(cliconstants.DefaultIstioOperatorNamespace, resources).Return(successfulResources, testErr)
			kubeClient.EXPECT().Delete(cliconstants.DefaultIstioOperatorNamespace, successfulResources).Return(successfulResources, nil)

			operatorManager := operator.NewManager(
				kubeClient,
				manifestBuilder,
				deploymentClient,
				imageNameParser,
				installConfig,
			)

			err := operatorManager.Install()
			Expect(err).To(testutils.HaveInErrorChain(operator.FailedToInstallOperator(testErr)))
		})
	})

	Context("ValidateOperatorNamespace", func() {
		It("can report that an install is required because no operator was found", func() {
			kubeClient := mock_kube.NewMockUnstructuredKubeClient(ctrl)
			manifestBuilder := mock_operator.NewMockInstallerManifestBuilder(ctrl)
			deploymentClient := mock_server.NewMockDeploymentClient(ctrl)
			imageNameParser := mock_docker.NewMockImageNameParser(ctrl)
			installConfig := buildInstallConfigWithDefaults(false, false)

			operatorManager := operator.NewManager(
				kubeClient,
				manifestBuilder,
				deploymentClient,
				imageNameParser,
				installConfig,
			)

			deploymentClient.EXPECT().GetDeployments(cliconstants.DefaultIstioOperatorNamespace, "").Return(&appsv1.DeploymentList{}, nil)

			installNeeded, err := operatorManager.ValidateOperatorNamespace(clusterName)
			Expect(installNeeded).To(BeTrue())
			Expect(err).NotTo(HaveOccurred())
		})

		It("reports the proper error if the deployment list fails", func() {
			kubeClient := mock_kube.NewMockUnstructuredKubeClient(ctrl)
			manifestBuilder := mock_operator.NewMockInstallerManifestBuilder(ctrl)
			deploymentClient := mock_server.NewMockDeploymentClient(ctrl)
			imageNameParser := mock_docker.NewMockImageNameParser(ctrl)
			installConfig := buildInstallConfigWithDefaults(false, false)

			operatorManager := operator.NewManager(
				kubeClient,
				manifestBuilder,
				deploymentClient,
				imageNameParser,
				installConfig,
			)

			deploymentClient.EXPECT().GetDeployments(cliconstants.DefaultIstioOperatorNamespace, "").Return(nil, testErr)

			installNeeded, err := operatorManager.ValidateOperatorNamespace(clusterName)
			Expect(installNeeded).To(BeFalse())
			Expect(err).To(testutils.HaveInErrorChain(operator.FailedToCheckIfOperatorExists(testErr, clusterName, cliconstants.DefaultIstioOperatorNamespace, cliconstants.DefaultIstioOperatorVersion)))
		})

		It("reports the proper error if the operator already exists at a different version than what was requested", func() {
			kubeClient := mock_kube.NewMockUnstructuredKubeClient(ctrl)
			manifestBuilder := mock_operator.NewMockInstallerManifestBuilder(ctrl)
			deploymentClient := mock_server.NewMockDeploymentClient(ctrl)
			imageNameParser := mock_docker.NewMockImageNameParser(ctrl)
			installConfig := buildInstallConfigWithDefaults(false, false)

			operatorManager := operator.NewManager(
				kubeClient,
				manifestBuilder,
				deploymentClient,
				imageNameParser,
				installConfig,
			)

			deploymentClient.EXPECT().GetDeployments(cliconstants.DefaultIstioOperatorNamespace, "").Return(&appsv1.DeploymentList{
				Items: []appsv1.Deployment{{
					ObjectMeta: metav1.ObjectMeta{
						Name: cliconstants.DefaultIstioOperatorDeploymentName,
					},
					Spec: appsv1.DeploymentSpec{
						Template: corev1.PodTemplateSpec{
							Spec: corev1.PodSpec{
								Containers: []corev1.Container{{
									Image: "istio-operator-test-image:0.6.9",
								}},
							},
						},
					},
				}},
			}, nil)

			imageNameParser.EXPECT().Parse("istio-operator-test-image:0.6.9").Return(&docker.Image{Tag: "4.2.0"}, nil)

			installNeeded, err := operatorManager.ValidateOperatorNamespace(clusterName)
			Expect(installNeeded).To(BeFalse())
			Expect(err).To(testutils.HaveInErrorChain(operator.OperatorVersionMismatch(clusterName, installConfig.InstallNamespace, installConfig.IstioOperatorVersion, "4.2.0")))
		})

		It("reports that an installation is not needed if the operator is already installed at the requested version", func() {
			kubeClient := mock_kube.NewMockUnstructuredKubeClient(ctrl)
			manifestBuilder := mock_operator.NewMockInstallerManifestBuilder(ctrl)
			deploymentClient := mock_server.NewMockDeploymentClient(ctrl)
			imageNameParser := mock_docker.NewMockImageNameParser(ctrl)
			installConfig := buildInstallConfigWithDefaults(false, false)

			operatorManager := operator.NewManager(
				kubeClient,
				manifestBuilder,
				deploymentClient,
				imageNameParser,
				installConfig,
			)

			deploymentClient.EXPECT().GetDeployments(cliconstants.DefaultIstioOperatorNamespace, "").Return(&appsv1.DeploymentList{
				Items: []appsv1.Deployment{{
					ObjectMeta: metav1.ObjectMeta{
						Name: cliconstants.DefaultIstioOperatorDeploymentName,
					},
					Spec: appsv1.DeploymentSpec{
						Template: corev1.PodTemplateSpec{
							Spec: corev1.PodSpec{
								Containers: []corev1.Container{{
									Image: "istio-operator-test-image:0.6.9",
								}},
							},
						},
					},
				}},
			}, nil)

			imageNameParser.EXPECT().Parse("istio-operator-test-image:0.6.9").Return(&docker.Image{Tag: cliconstants.DefaultIstioOperatorVersion}, nil)

			installNeeded, err := operatorManager.ValidateOperatorNamespace(clusterName)
			Expect(installNeeded).To(BeFalse())
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
