// Code generated by hack/update-lifted. DO NOT EDIT.

// Package lifted contains the files lifted from other projects.
package lifted

/*
| lifted file              | source file                                                                                                                              | const/var/type/func                               | changed |
|--------------------------|------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------|---------|
| corehelpers.go           | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/apis/core/helper/helpers.go#L57-L61                                       | func IsQuotaHugePageResourceName                  | N       |
| corehelpers.go           | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/apis/core/helper/helpers.go#L212-L232                                     | var standardQuotaResources                        | N       |
| corehelpers.go           | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/apis/core/helper/helpers.go#L234-L238                                     | func IsStandardQuotaResourceName                  | N       |
| corehelpers.go           | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/apis/core/helper/helpers.go#L240-L261                                     | var standardResources                             | N       |
| corehelpers.go           | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/apis/core/helper/helpers.go#L263-L266                                     | func IsStandardResourceName                       | N       |
| corehelpers.go           | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/apis/core/helper/helpers.go#L268-L278                                     | var integerResources                              | N       |
| corehelpers.go           | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/apis/core/helper/helpers.go#L280-L283                                     | func IsIntegerResourceName                        | N       |
| corehelpers.go           | https://github.com/kubernetes/kubernetes/blob/release-1.27/pkg/apis/core/validation/validation.go                                        | type ValidateNameFunc                             | N       |
| corehelpers.go           | https://github.com/kubernetes/kubernetes/blob/release-1.27/pkg/apis/core/helper/helpers.go                                               | var standardFinalizers                            | N       |
| corehelpers.go           | https://github.com/kubernetes/kubernetes/blob/release-1.27/staging/src/k8s.io/apimachinery/pkg/api/validation/generic.go                 | func NameIsDNSSubdomain                           | N       |
| corev1helpers.go         | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/apis/core/v1/helper/helpers.go#L31-L46                                    | func IsExtendedResourceName                       | Y       |
| corev1helpers.go         | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/apis/core/v1/helper/helpers.go#L48-L51                                    | func IsPrefixedNativeResource                     | N       |
| corev1helpers.go         | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/apis/core/v1/helper/helpers.go#L54-L60                                    | func IsNativeResource                             | N       |
| corev1helpers.go         | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/apis/core/v1/helper/helpers.go#L62-L66                                    | func IsHugePageResourceName                       | N       |
| corev1helpers.go         | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/apis/core/v1/helper/helpers.go#L132-L135                                  | func IsAttachableVolumeResourceName               | N       |
| corevalidation.go        | https://github.com/kubernetes/kubernetes/blob/release-1.22/pkg/apis/core/validation/validation.go#L59                                    | const isNegativeErrorMsg                          | N       |
| corevalidation.go        | https://github.com/kubernetes/kubernetes/blob/release-1.22/pkg/apis/core/validation/validation.go#L60                                    | const isInvalidQuotaResource                      | N       |
| corevalidation.go        | https://github.com/kubernetes/kubernetes/blob/release-1.22/pkg/apis/core/validation/validation.go#L62                                    | const isNotIntegerErrorMsg                        | N       |
| corevalidation.go        | https://github.com/kubernetes/kubernetes/blob/release-1.22/pkg/apis/core/validation/validation.go#L225-L228                              | var ValidatePodName                               | N       |
| corevalidation.go        | https://github.com/kubernetes/kubernetes/blob/release-1.22/pkg/apis/core/validation/validation.go#L311-L318                              | func ValidateNonnegativeQuantity                  | N       |
| corevalidation.go        | https://github.com/kubernetes/kubernetes/blob/release-1.22/pkg/apis/core/validation/validation.go#L5036-L5054                            | func validateResourceName                         | Y       |
| corevalidation.go        | https://github.com/kubernetes/kubernetes/blob/release-1.22/pkg/apis/core/validation/validation.go#L5073-L5084                            | func ValidateResourceQuotaResourceName            | Y       |
| corevalidation.go        | https://github.com/kubernetes/kubernetes/blob/release-1.22/pkg/apis/core/validation/validation.go#L5651-L5661                            | func ValidateResourceQuantityValue                | Y       |
| deployment.go            | https://github.com/kubernetes/kubernetes/blob/release-1.22/pkg/controller/controller_utils.go#L1003-L1012                                | type ReplicaSetsByCreationTimestamp               | N       |
| deployment.go            | https://github.com/kubernetes/kubernetes/blob/release-1.22/pkg/controller/deployment/util/deployment_util.go#L569-L594                   | func ListReplicaSetsByDeployment                  | Y       |
| deployment.go            | https://github.com/kubernetes/kubernetes/blob/release-1.22/pkg/controller/deployment/util/deployment_util.go#L596-L628                   | func ListPodsByRS                                 | Y       |
| deployment.go            | https://github.com/kubernetes/kubernetes/blob/release-1.22/pkg/controller/deployment/util/deployment_util.go#L630-L642                   | func EqualIgnoreHash                              | N       |
| deployment.go            | https://github.com/kubernetes/kubernetes/blob/release-1.22/pkg/controller/deployment/util/deployment_util.go#L536-L544                   | func GetNewReplicaSet                             | Y       |
| deployment.go            | https://github.com/kubernetes/kubernetes/blob/release-1.22/pkg/controller/deployment/util/deployment_util.go#L644-L658                   | func FindNewReplicaSet                            | N       |
| deployment_test.go       | https://github.com/kubernetes/kubernetes/blob/release-1.22/pkg/controller/deployment/util/deployment_util_test.go#LL151C1-L186C2         | func generateDeployment                           | N       |
| deployment_test.go       | https://github.com/kubernetes/kubernetes/blob/release-1.22/pkg/controller/deployment/util/deployment_util_test.go#LL129C1-L145C2         | func generateRS                                   | Y       |
| deployment_test.go       | https://github.com/kubernetes/kubernetes/blob/release-1.22/pkg/controller/deployment/util/deployment_util_test.go#LL118C1-L127C2         | func newDControllerRef                            | N       |
| deployment_test.go       | https://github.com/kubernetes/kubernetes/blob/release-1.22/pkg/controller/deployment/util/deployment_util_test.go#L326                   | func generatePodTemplateSpec                      | N       |
| deployment_test.go       | https://github.com/kubernetes/kubernetes/blob/release-1.22/pkg/controller/deployment/util/deployment_util_test.go#L339                   | func TestEqualIgnoreHash                          | N       |
| deployment_test.go       | https://github.com/kubernetes/kubernetes/blob/release-1.22/pkg/controller/deployment/util/deployment_util_test.go#L432                   | func TestFindNewReplicaSet                        | N       |
| discovery.go             | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/controller/garbagecollector/garbagecollector.go#L696-L732                 | func GetDeletableResources                        | N       |
| discovery_test.go        | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/controller/garbagecollector/garbagecollector_test.go#L943-L990            | type fakeServerResources                          | N       |
| discovery_test.go        | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/controller/garbagecollector/garbagecollector_test.go#L707-L797            | func TestGetDeletableResources                    | N       |
| genutils.go              | https://github.com/kubernetes/kubernetes/blob/release-1.25/cmd/genutils/genutils.go                                                      | func OutDir                                       | N       |
| genutils_test.go         | https://github.com/kubernetes/kubernetes/blob/release-1.25/cmd/genutils/genutils_test.go#L23-L28                                         | func TestValidDir                                 | N       |
| genutils_test.go         | https://github.com/kubernetes/kubernetes/blob/release-1.25/cmd/genutils/genutils_test.go#L30-L35                                         | func TestInvalidDir                               | N       |
| genutils_test.go         | https://github.com/kubernetes/kubernetes/blob/release-1.25/cmd/genutils/genutils_test.go#L37-L42                                         | func TestNotDir                                   | N       |
| nodeselector.go          | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/apis/core/helper/helpers.go#L365-L397                                     | func NodeSelectorRequirementsAsSelector           | N       |
| nodeselector_test.go     | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/apis/core/helper/helpers_test.go#L138-L200                                | func TestNodeSelectorRequirementsAsSelector       | N       |
| objectwatcher.go         | https://github.com/kubernetes-sigs/kubefed/blob/master/pkg/controller/util/propagatedversion.go#L35-L43                                  | func ObjectVersion                                | N       |
| objectwatcher.go         | https://github.com/kubernetes-sigs/kubefed/blob/master/pkg/controller/util/propagatedversion.go#L45-L59                                  | func ObjectNeedsUpdate                            | N       |
| objectwatcher.go         | https://github.com/kubernetes-sigs/kubefed/blob/master/pkg/controller/util/meta.go#L63-L80                                               | func objectMetaObjEquivalent                      | Y       |
| podtemplate.go           | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/controller/controller_utils.go#L466-L472                                  | func getPodsLabelSet                              | N       |
| podtemplate.go           | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/controller/controller_utils.go#L474-L478                                  | func getPodsFinalizers                            | N       |
| podtemplate.go           | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/controller/controller_utils.go#L480-L486                                  | func getPodsAnnotationSet                         | N       |
| podtemplate.go           | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/controller/controller_utils.go#L488-L495                                  | func getPodsPrefix                                | N       |
| podtemplate.go           | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/controller/controller_utils.go#L539-L562                                  | func GetPodFromTemplate                           | Y       |
| requestinfo.go           | https://github.com/kubernetes/apiserver/blob/release-1.23/pkg/endpoints/request/requestinfo.go#L88-L247                                  | func NewRequestInfo                               | Y       |
| requestinfo.go           | https://github.com/kubernetes/apiserver/blob/release-1.23/pkg/endpoints/request/requestinfo.go#L267-L274                                 | func SplitPath                                    | Y       |
| requestinfo.go           | https://github.com/kubernetes/apiserver/blob/release-1.23/pkg/endpoints/request/requestinfo.go#L73-L74                                   | var specialVerbsNoSubresources                    | N       |
| requestinfo.go           | https://github.com/kubernetes/apiserver/blob/release-1.23/pkg/endpoints/request/requestinfo.go#L76-L78                                   | var namespaceSubresources                         | N       |
| requestinfo.go           | https://github.com/kubernetes/apiserver/blob/release-1.23/pkg/endpoints/request/requestinfo.go#L67-L71                                   | var specialVerbs                                  | N       |
| requestinfo_test.go      | https://github.com/kubernetes/apiserver/blob/release-1.23/pkg/endpoints/request/requestinfo_test.go#L30-L148                             | func TestGetAPIRequestInfo                        | Y       |
| requestinfo_test.go      | https://github.com/kubernetes/apiserver/blob/release-1.23/pkg/endpoints/request/requestinfo_test.go#L150-L184                            | func TestGetNonAPIRequestInfo                     | Y       |
| resourcename.go          | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/scheduler/util/utils.go#L144-L148                                         | func IsScalarResourceName                         | Y       |
| retain.go                | https://github.com/kubernetes-sigs/kubefed/blob/master/pkg/controller/sync/dispatch/retain.go                                            | func RetainServiceFields                          | Y       |
| retain.go                | https://github.com/kubernetes-sigs/kubefed/blob/master/pkg/controller/sync/dispatch/retain.go                                            | func retainServiceHealthCheckNodePort             | Y       |
| retain.go                | https://github.com/kubernetes-sigs/kubefed/blob/master/pkg/controller/sync/dispatch/retain.go                                            | func retainServiceClusterIP                       | Y       |
| retain.go                | https://github.com/kubernetes-sigs/kubefed/blob/master/pkg/controller/sync/dispatch/retain.go                                            | func RetainServiceAccountFields                   | Y       |
| retain_test.go           | https://github.com/kubernetes-sigs/kubefed/blob/master/pkg/controller/sync/dispatch/retain_test.go#L91-L174                              | func TestRetainHealthCheckNodePortInServiceFields | Y       |
| retain_test.go           | https://github.com/kubernetes-sigs/kubefed/blob/master/pkg/controller/sync/dispatch/retain_test.go#L176-L292                             | func TestRetainClusterIPInServiceFields           | Y       |
| taint.go                 | https://github.com/kubernetes/kubernetes/blob/release-1.23/staging/src/k8s.io/kubectl/pkg/cmd/taint/utils.go#L37-L73                     | func ParseTaints                                  | N       |
| taint.go                 | https://github.com/kubernetes/kubernetes/blob/release-1.23/staging/src/k8s.io/kubectl/pkg/cmd/taint/utils.go#L75-L118                    | func parseTaint                                   | N       |
| taint.go                 | https://github.com/kubernetes/kubernetes/blob/release-1.23/staging/src/k8s.io/kubectl/pkg/cmd/taint/utils.go#L120-L126                   | func validateTaintEffect                          | N       |
| validateclustertaints.go | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/apis/core/validation/validation.go#L5001-L5033                            | func ValidateClusterTaints                        | Y       |
| validateclustertaints.go | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/apis/core/validation/validation.go#L3305-L3326                            | func validateClusterTaintEffect                   | Y       |
| validatingmci.go         | https://github.com/kubernetes/kubernetes/blob/release-1.27/pkg/apis/networking/validation/validation.go#L326-L348                        | func ValidateIngressSpec                          | N       |
| validatingmci.go         | https://github.com/kubernetes/kubernetes/blob/release-1.27/pkg/apis/networking/validation/validation.go#L468C1-L509                      | func validateIngressBackend                       | N       |
| validatingmci.go         | https://github.com/kubernetes/kubernetes/blob/release-1.27/pkg/apis/networking/validation/validation.go#L547C1-L578                      | func validateIngressTypedLocalObjectReference     | N       |
| validatingmci.go         | https://github.com/kubernetes/kubernetes/blob/release-1.27/pkg/apis/networking/validation/validation.go#L379-L409                        | func validateIngressRules                         | N       |
| validatingmci.go         | https://github.com/kubernetes/kubernetes/blob/release-1.27/pkg/apis/networking/validation/validation.go#L411C1-L417                      | func validateIngressRuleValue                     | N       |
| validatingmci.go         | https://github.com/kubernetes/kubernetes/blob/release-1.27/pkg/apis/networking/validation/validation.go#L419-L428                        | func validateHTTPIngressRuleValue                 | N       |
| validatingmci.go         | https://github.com/kubernetes/kubernetes/blob/release-1.27/pkg/apis/networking/validation/validation.go#L430-L466                        | func validateHTTPIngressPath                      | N       |
| validatingmci.go         | https://github.com/kubernetes/kubernetes/blob/release-1.27/pkg/apis/networking/validation/validation.go#L299-L324                        | func validateIngressTLS                           | N       |
| validatingmci.go         | https://github.com/kubernetes/kubernetes/blob/release-1.27/pkg/apis/networking/validation/validation.go#L641-L646                        | func validateTLSSecretName                        | N       |
| validatingmci.go         | https://github.com/kubernetes/kubernetes/blob/release-1.27/pkg/apis/networking/validation/validation.go#L357-L377                        | func ValidateIngressLoadBalancerStatus            | N       |
| validatingmci_test.go    | https://github.com/kubernetes/kubernetes/blob/release-1.27/pkg/apis/networking/validation/validation_test.go                             | func TestValidateIngress                          | Y       |
| validatingmci_test.go    | https://github.com/kubernetes/kubernetes/blob/release-1.27/pkg/apis/networking/validation/validation_test.go                             | func TestValidateIngressTLS                       | N       |
| validatingmci_test.go    | https://github.com/kubernetes/kubernetes/blob/release-1.27/pkg/apis/networking/validation/validation_test.go                             | func TestValidateEmptyIngressTLS                  | N       |
| validatingmci_test.go    | https://github.com/kubernetes/kubernetes/blob/release-1.27/pkg/apis/networking/validation/validation_test.go                             | func TestValidateIngressStatusUpdate              | Y       |
| visitpod.go              | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/api/v1/pod/util.go#L53-L63                                                | type ContainerType                                | N       |
| visitpod.go              | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/api/v1/pod/util.go#L65-L66                                                | const AllContainers                               | N       |
| visitpod.go              | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/api/v1/pod/util.go#L78-L80                                                | type ContainerVisitor                             | N       |
| visitpod.go              | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/api/v1/pod/util.go#L82-L83                                                | type Visitor                                      | N       |
| visitpod.go              | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/api/v1/pod/util.go#L85-L94                                                | func skipEmptyNames                               | N       |
| visitpod.go              | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/api/v1/pod/util.go#L96-L123                                               | func VisitContainers                              | N       |
| visitpod.go              | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/api/v1/pod/util.go#L125-L195                                              | func VisitPodSecretNames                          | N       |
| visitpod.go              | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/api/v1/pod/util.go#L197-L213                                              | func visitContainerSecretNames                    | N       |
| visitpod.go              | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/api/v1/pod/util.go#L215-L243                                              | func VisitPodConfigmapNames                       | N       |
| visitpod.go              | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/api/v1/pod/util.go#L245-L261                                              | func visitContainerConfigmapNames                 | N       |
| visitpod_test.go         | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/api/v1/pod/util.go#L68-L76                                                | func AllFeatureEnabledContainers                  | N       |
| visitpod_test.go         | https://github.com/kubernetes/kubernetes/blob/release-1.23/staging/src/k8s.io/component-base/featuregate/testing/feature_gate.go#L26-L44 | func SetFeatureGateDuringTest                     | N       |
| visitpod_test.go         | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/api/v1/pod/util_test.go#L205-L392                                         | func TestVisitContainers                          | Y       |
| visitpod_test.go         | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/api/v1/pod/util_test.go#L394-L551                                         | func TestPodSecrets                               | Y       |
| visitpod_test.go         | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/api/v1/pod/util_test.go#L553-L591                                         | func collectResourcePaths                         | N       |
| visitpod_test.go         | https://github.com/kubernetes/kubernetes/blob/release-1.23/pkg/api/v1/pod/util_test.go#L593-L695                                         | func TestPodConfigmaps                            | Y       |
*/
