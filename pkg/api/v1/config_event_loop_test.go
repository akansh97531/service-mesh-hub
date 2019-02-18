// Code generated by solo-kit. DO NOT EDIT.

// +build solokit

package v1

import (
	"context"
	"sync"
	"time"

	gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	istio_authentication_v1alpha1 "github.com/solo-io/supergloo/pkg/api/external/istio/authorization/v1alpha1"
	istio_networking_v1alpha3 "github.com/solo-io/supergloo/pkg/api/external/istio/networking/v1alpha3"
	core_kubernetes_io "github.com/solo-io/supergloo/pkg/api/external/kubernetes/core/v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
)

var _ = Describe("ConfigEventLoop", func() {
	var (
		namespace string
		emitter   ConfigEmitter
		err       error
	)

	BeforeEach(func() {

		meshClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		meshClient, err := NewMeshClient(meshClientFactory)
		Expect(err).NotTo(HaveOccurred())

		meshGroupClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		meshGroupClient, err := NewMeshGroupClient(meshGroupClientFactory)
		Expect(err).NotTo(HaveOccurred())

		routingRuleClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		routingRuleClient, err := NewRoutingRuleClient(routingRuleClientFactory)
		Expect(err).NotTo(HaveOccurred())

		securityRuleClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		securityRuleClient, err := NewSecurityRuleClient(securityRuleClientFactory)
		Expect(err).NotTo(HaveOccurred())

		tlsSecretClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		tlsSecretClient, err := NewTlsSecretClient(tlsSecretClientFactory)
		Expect(err).NotTo(HaveOccurred())

		upstreamClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		upstreamClient, err := gloo_solo_io.NewUpstreamClient(upstreamClientFactory)
		Expect(err).NotTo(HaveOccurred())

		podClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		podClient, err := core_kubernetes_io.NewPodClient(podClientFactory)
		Expect(err).NotTo(HaveOccurred())

		destinationRuleClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		destinationRuleClient, err := istio_networking_v1alpha3.NewDestinationRuleClient(destinationRuleClientFactory)
		Expect(err).NotTo(HaveOccurred())

		virtualServiceClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		virtualServiceClient, err := istio_networking_v1alpha3.NewVirtualServiceClient(virtualServiceClientFactory)
		Expect(err).NotTo(HaveOccurred())

		meshPolicyClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		meshPolicyClient, err := istio_authentication_v1alpha1.NewMeshPolicyClient(meshPolicyClientFactory)
		Expect(err).NotTo(HaveOccurred())

		emitter = NewConfigEmitter(meshClient, meshGroupClient, routingRuleClient, securityRuleClient, tlsSecretClient, upstreamClient, podClient, destinationRuleClient, virtualServiceClient, meshPolicyClient)
	})
	It("runs sync function on a new snapshot", func() {
		_, err = emitter.Mesh().Write(NewMesh(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		_, err = emitter.MeshGroup().Write(NewMeshGroup(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		_, err = emitter.RoutingRule().Write(NewRoutingRule(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		_, err = emitter.SecurityRule().Write(NewSecurityRule(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		_, err = emitter.TlsSecret().Write(NewTlsSecret(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		_, err = emitter.Upstream().Write(gloo_solo_io.NewUpstream(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		_, err = emitter.Pod().Write(core_kubernetes_io.NewPod(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		_, err = emitter.DestinationRule().Write(istio_networking_v1alpha3.NewDestinationRule(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		_, err = emitter.VirtualService().Write(istio_networking_v1alpha3.NewVirtualService(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		_, err = emitter.MeshPolicy().Write(istio_authentication_v1alpha1.NewMeshPolicy(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		sync := &mockConfigSyncer{}
		el := NewConfigEventLoop(emitter, sync)
		_, err := el.Run([]string{namespace}, clients.WatchOpts{})
		Expect(err).NotTo(HaveOccurred())
		Eventually(sync.Synced, 5*time.Second).Should(BeTrue())
	})
})

type mockConfigSyncer struct {
	synced bool
	mutex  sync.Mutex
}

func (s *mockConfigSyncer) Synced() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.synced
}

func (s *mockConfigSyncer) Sync(ctx context.Context, snap *ConfigSnapshot) error {
	s.mutex.Lock()
	s.synced = true
	s.mutex.Unlock()
	return nil
}
