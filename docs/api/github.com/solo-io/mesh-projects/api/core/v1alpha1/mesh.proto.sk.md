
---
title: "mesh.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `core.zephyr.solo.io` 
#### Types:


- [MeshSpec](#meshspec)
- [MeshInstallation](#meshinstallation)
- [IstioMesh](#istiomesh)
- [AwsAppMesh](#awsappmesh)
- [LinkerdMesh](#linkerdmesh)
- [ConsulConnectMesh](#consulconnectmesh)
- [MeshGroupSpec](#meshgroupspec)
  



##### Source File: [github.com/solo-io/mesh-projects/api/core/v1alpha1/mesh.proto](https://github.com/solo-io/mesh-projects/blob/master/api/core/v1alpha1/mesh.proto)





---
### MeshSpec

 
Meshes represent a currently registered service mesh.

```yaml
"istio": .core.zephyr.solo.io.IstioMesh
"awsAppMesh": .core.zephyr.solo.io.AwsAppMesh
"linkerd": .core.zephyr.solo.io.LinkerdMesh
"consulConnect": .core.zephyr.solo.io.ConsulConnectMesh
"cluster": .core.zephyr.solo.io.ResourceRef
"services": []core.zephyr.solo.io.ResourceRef
"pods": []core.zephyr.solo.io.ResourceRef

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `istio` | [.core.zephyr.solo.io.IstioMesh](../mesh.proto.sk/#istiomesh) |  Only one of `istio`, `awsAppMesh`, or `consulConnect` can be set. |  |
| `awsAppMesh` | [.core.zephyr.solo.io.AwsAppMesh](../mesh.proto.sk/#awsappmesh) |  Only one of `awsAppMesh`, `istio`, or `consulConnect` can be set. |  |
| `linkerd` | [.core.zephyr.solo.io.LinkerdMesh](../mesh.proto.sk/#linkerdmesh) |  Only one of `linkerd`, `istio`, or `consulConnect` can be set. |  |
| `consulConnect` | [.core.zephyr.solo.io.ConsulConnectMesh](../mesh.proto.sk/#consulconnectmesh) |  Only one of `consulConnect`, `istio`, or `linkerd` can be set. |  |
| `cluster` | [.core.zephyr.solo.io.ResourceRef](../ref.proto.sk/#resourceref) |  |  |
| `services` | [[]core.zephyr.solo.io.ResourceRef](../ref.proto.sk/#resourceref) | List of services, these can either be kube services, or smh services. |  |
| `pods` | [[]core.zephyr.solo.io.ResourceRef](../ref.proto.sk/#resourceref) | pods also contain ServiceAccount Information, therefore we do not need to store it here. |  |




---
### MeshInstallation



```yaml
"installationNamespace": string
"version": string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `installationNamespace` | `string` | where the control plane has been installed. |  |
| `version` | `string` | version of the mesh which has been installed. |  |




---
### IstioMesh

 
Mesh object representing an installed Istio control plane

```yaml
"installation": .core.zephyr.solo.io.MeshInstallation

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `installation` | [.core.zephyr.solo.io.MeshInstallation](../mesh.proto.sk/#meshinstallation) |  |  |




---
### AwsAppMesh

 
Mesh object representing AWS App Mesh

```yaml
"installation": .core.zephyr.solo.io.MeshInstallation
"region": string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `installation` | [.core.zephyr.solo.io.MeshInstallation](../mesh.proto.sk/#meshinstallation) |  |  |
| `region` | `string` | The AWS region the AWS App Mesh control plane resources exist in. |  |




---
### LinkerdMesh

 
Mesh object representing an installed Linkerd control plane

```yaml
"installation": .core.zephyr.solo.io.MeshInstallation

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `installation` | [.core.zephyr.solo.io.MeshInstallation](../mesh.proto.sk/#meshinstallation) |  |  |




---
### ConsulConnectMesh



```yaml
"installation": .core.zephyr.solo.io.MeshInstallation

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `installation` | [.core.zephyr.solo.io.MeshInstallation](../mesh.proto.sk/#meshinstallation) |  |  |




---
### MeshGroupSpec



```yaml
"displayName": string
"meshes": []core.zephyr.solo.io.ResourceRef

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `displayName` | `string` | User-provided display name for the mesh group. |  |
| `meshes` | [[]core.zephyr.solo.io.ResourceRef](../ref.proto.sk/#resourceref) | The meshes contained in this group. |  |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->