package wire

import (
	"github.com/google/wire"
	smh_security_providers "github.com/solo-io/service-mesh-hub/pkg/api/security.smh.solo.io/v1alpha1/providers"
	networking_multicluster "github.com/solo-io/service-mesh-hub/pkg/mesh-networking/compute-target"
	controller_factories "github.com/solo-io/service-mesh-hub/pkg/mesh-networking/compute-target/controllers"
)

var (
	ClientFactoryProviderSet = wire.NewSet(
		smh_security_providers.VirtualMeshCertificateSigningRequestClientFactoryProvider,
		NewClientFactories,
	)

	ControllerFactoryProviderSet = wire.NewSet(
		controller_factories.NewVirtualMeshCSRControllerFactory,
		NewControllerFactories,
	)
)

func NewClientFactories(
	VirtualMeshCertificateSigningRequestClientFactory smh_security_providers.VirtualMeshCertificateSigningRequestClientFactory,
) *networking_multicluster.ClientFactories {
	return &networking_multicluster.ClientFactories{
		VirtualMeshCSRClientFactory: VirtualMeshCertificateSigningRequestClientFactory,
	}
}

func NewControllerFactories(
	CSRControllerFactory controller_factories.VirtualMeshCSRControllerFactory,
) *networking_multicluster.ControllerFactories {
	return &networking_multicluster.ControllerFactories{VirtualMeshCSRControllerFactory: CSRControllerFactory}
}
