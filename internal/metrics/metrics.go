// SPDX-License-Identifier: Apache-2.0
package metrics

import (
    "net/http"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartMetricsServer() {
    http.Handle("/metrics", promhttp.Handler())
    go http.ListenAndServe(":9090", nil)
}
