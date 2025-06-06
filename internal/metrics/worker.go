package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

var (
	labels = []string{
		"namespace", "worker", "pool",
	}

	nodeLabels = []string{
		"nodeName", "pool",
	}

	GpuTflopsRequest = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "gpu_tflops_request",
		},
		labels,
	)

	GpuTflopsLimit = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "gpu_tflops_limit",
		},
		labels,
	)

	VramBytesRequest = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "vram_bytes_request",
		},
		labels,
	)

	VramBytesLimit = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "vram_bytes_limit",
		},
		labels,
	)

	AllocatedTflopsPercent = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "allocated_compute_percentage",
		},
		nodeLabels,
	)

	AllocatedVramBytes = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "allocated_memory_bytes",
		},
		nodeLabels,
	)
)

func init() {
	metrics.Registry.MustRegister(GpuTflopsRequest, GpuTflopsLimit, VramBytesRequest, VramBytesLimit)
}
