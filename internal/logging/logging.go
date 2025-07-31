// SPDX-License-Identifier: Apache-2.0
package logging

import "go.uber.org/zap"

var Logger *zap.Logger

func InitLogger() {
    Logger, _ = zap.NewProduction()
}
