// Code generated by protoc-gen-solo-kit. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	encryption_istio_io "github.com/solo-io/supergloo/pkg/api/external/istio/encryption/v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
)

var _ = Describe("InstallEventLoop", func() {
	var (
		namespace string
		emitter   InstallEmitter
		err       error
	)

	BeforeEach(func() {

		installClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		installClient, err := NewInstallClient(installClientFactory)
		Expect(err).NotTo(HaveOccurred())

		istioCacertsSecretClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		istioCacertsSecretClient, err := encryption_istio_io.NewIstioCacertsSecretClient(istioCacertsSecretClientFactory)
		Expect(err).NotTo(HaveOccurred())

		emitter = NewInstallEmitter(installClient, istioCacertsSecretClient)
	})
	It("runs sync function on a new snapshot", func() {
		_, err = emitter.Install().Write(NewInstall(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		_, err = emitter.IstioCacertsSecret().Write(encryption_istio_io.NewIstioCacertsSecret(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		sync := &mockInstallSyncer{}
		el := NewInstallEventLoop(emitter, sync)
		_, err := el.Run([]string{namespace}, clients.WatchOpts{})
		Expect(err).NotTo(HaveOccurred())
		Eventually(func() bool { return sync.synced }, time.Second).Should(BeTrue())
	})
})

type mockInstallSyncer struct {
	synced bool
}

func (s *mockInstallSyncer) Sync(ctx context.Context, snap *InstallSnapshot) error {
	s.synced = true
	return nil
}
