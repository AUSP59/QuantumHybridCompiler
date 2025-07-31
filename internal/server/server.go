// SPDX-License-Identifier: Apache-2.0
package server

import (
    "net/http"
    "github.com/labstack/echo/v4"
)

// StartServer runs the REST API server
func StartServer() {
    e := echo.New()
    e.GET("/health", func(c echo.Context) error {
        return c.String(http.StatusOK, "QuantumHybridCompiler API running")
    })
    e.Logger.Fatal(e.Start(":8080"))
}
