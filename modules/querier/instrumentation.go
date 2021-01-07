package querier

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/weaveworks/common/instrument"
	"github.com/weaveworks/common/middleware"
)

func (q *Querier) NewInstrumentedHandler(reg prometheus.Registerer) http.Handler {
	// Prometheus histograms for requests to the querier.
	querierRequestDuration := promauto.With(reg).NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "tempo",
		Name:      "querier_request_duration_seconds",
		Help:      "Time (in seconds) spent serving HTTP requests to the querier.",
		Buckets:   instrument.DefBuckets,
	}, []string{"method", "route", "status_code", "ws"})

	receivedMessageSize := promauto.With(reg).NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "tempo",
		Name:      "querier_request_message_bytes",
		Help:      "Size (in bytes) of messages received in the request to the querier.",
		Buckets:   middleware.BodySizeBuckets,
	}, []string{"method", "route"})

	sentMessageSize := promauto.With(reg).NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "tempo",
		Name:      "querier_response_message_bytes",
		Help:      "Size (in bytes) of messages sent in response by the querier.",
		Buckets:   middleware.BodySizeBuckets,
	}, []string{"method", "route"})

	inflightRequests := promauto.With(reg).NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "tempo",
		Name:      "querier_inflight_requests",
		Help:      "Current number of inflight requests to the querier.",
	}, []string{"method", "route"})

	router := mux.NewRouter()

	// Use a separate metric for the querier in order to differentiate requests from the query-frontend
	inst := middleware.Instrument{
		RouteMatcher:     router,
		Duration:         querierRequestDuration,
		RequestBodySize:  receivedMessageSize,
		ResponseBodySize: sentMessageSize,
		InflightRequests: inflightRequests,
	}
	middlewares := middleware.Merge(inst)
	router.Use(middlewares.Wrap)

	router.Path("/querier/api/traces/{traceID}").Handler(http.HandlerFunc(q.TraceByIDHandler))

	// Add a middleware to extract the trace context and add a header.
	handler := nethttp.MiddlewareFunc(opentracing.GlobalTracer(), router.ServeHTTP, nethttp.OperationNameFunc(func(r *http.Request) string {
		return "internalQuerier"
	}))
	return handler
}