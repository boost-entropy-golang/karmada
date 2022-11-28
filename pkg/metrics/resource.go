package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	utilmetrics "github.com/karmada-io/karmada/pkg/util/metrics"
)

const (
	resourceMatchPolicyDurationMetricsName = "resource_match_policy_duration_seconds"
	resourceApplyPolicyDurationMetricsName = "resource_apply_policy_duration_seconds"
	policyApplyAttemptsMetricsName         = "policy_apply_attempts_total"
	syncWorkDurationMetricsName            = "binding_sync_work_duration_seconds"
	syncWorkloadDurationMetricsName        = "work_sync_workload_duration_seconds"
)

var (
	findMatchedPolicyDurationHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    resourceMatchPolicyDurationMetricsName,
		Help:    "Duration in seconds to find a matched propagation policy for the resource template.",
		Buckets: prometheus.ExponentialBuckets(0.001, 2, 12),
	}, []string{"apiVersion", "kind", "name", "namespace"})

	applyPolicyDurationHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    resourceApplyPolicyDurationMetricsName,
		Help:    "Duration in seconds to apply a propagation policy for the resource template. By the result, 'error' means a resource template failed to apply the policy. Otherwise 'success'.",
		Buckets: prometheus.ExponentialBuckets(0.001, 2, 12),
	}, []string{"apiVersion", "kind", "name", "namespace", "result"})

	policyApplyAttempts = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: policyApplyAttemptsMetricsName,
		Help: "Number of attempts to be applied for a propagation policy. By the result, 'error' means a resource template failed to apply the policy. Otherwise 'success'.",
	}, []string{"namespace", "name", "result"})

	syncWorkDurationHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    syncWorkDurationMetricsName,
		Help:    "Duration in seconds to sync works for a binding object. By the result, 'error' means a binding failed to sync works. Otherwise 'success'.",
		Buckets: prometheus.ExponentialBuckets(0.001, 2, 12),
	}, []string{"namespace", "name", "result"})

	syncWorkloadDurationHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    syncWorkloadDurationMetricsName,
		Help:    "Duration in seconds to sync the workload to a target cluster. By the result, 'error' means a work failed to sync workloads. Otherwise 'success'.",
		Buckets: prometheus.ExponentialBuckets(0.001, 2, 12),
	}, []string{"namespace", "name", "result"})
)

// ObserveFindMatchedPolicyLatency records the duration for the resource finding a matched policy.
func ObserveFindMatchedPolicyLatency(object *unstructured.Unstructured, start time.Time) {
	findMatchedPolicyDurationHistogram.WithLabelValues(object.GetAPIVersion(), object.GetKind(), object.GetName(), object.GetNamespace()).Observe(utilmetrics.DurationInSeconds(start))
}

// ObserveApplyPolicyAttemptAndLatency records the duration for the resource applying a policy and a applying attempt for the policy.
func ObserveApplyPolicyAttemptAndLatency(object *unstructured.Unstructured, policyMetaData metav1.ObjectMeta, err error, start time.Time) {
	applyPolicyDurationHistogram.WithLabelValues(object.GetAPIVersion(), object.GetKind(), object.GetName(), object.GetNamespace(), utilmetrics.GetResultByError(err)).Observe(utilmetrics.DurationInSeconds(start))
	policyApplyAttempts.WithLabelValues(policyMetaData.Namespace, policyMetaData.Name, utilmetrics.GetResultByError(err)).Inc()
}

// ObserveSyncWorkLatency records the duration to sync works for a binding object.
func ObserveSyncWorkLatency(bindingMetaData metav1.ObjectMeta, err error, start time.Time) {
	syncWorkDurationHistogram.WithLabelValues(bindingMetaData.Namespace, bindingMetaData.Name, utilmetrics.GetResultByError(err)).Observe(utilmetrics.DurationInSeconds(start))
}

// ObserveSyncWorkloadLatency records the duration to sync the workload to a target cluster.
func ObserveSyncWorkloadLatency(workMetadata metav1.ObjectMeta, err error, start time.Time) {
	syncWorkloadDurationHistogram.WithLabelValues(workMetadata.Namespace, workMetadata.Name, utilmetrics.GetResultByError(err)).Observe(utilmetrics.DurationInSeconds(start))
}

// ResourceCollectors returns the collectors about resources.
func ResourceCollectors() []prometheus.Collector {
	return []prometheus.Collector{
		applyPolicyDurationHistogram,
		findMatchedPolicyDurationHistogram,
		policyApplyAttempts,
		syncWorkDurationHistogram,
		syncWorkloadDurationHistogram,
	}
}

// ResourceCollectorsForAgent returns the collectors about resources for karmada-agent.
func ResourceCollectorsForAgent() []prometheus.Collector {
	return []prometheus.Collector{
		syncWorkloadDurationHistogram,
	}
}
