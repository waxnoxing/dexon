// Copyright 2018 The dexon-consensus Authors
// This file is part of the dexon-consensus library.
//
// The dexon-consensus library is free software: you can redistribute it
// and/or modify it under the terms of the GNU Lesser General Public License as
// published by the Free Software Foundation, either version 3 of the License,
// or (at your option) any later version.
//
// The dexon-consensus library is distributed in the hope that it will be
// useful, but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser
// General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the dexon-consensus library. If not, see
// <http://www.gnu.org/licenses/>.

package common

import "log"

// Logger define the way to receive logs from Consensus instance.
// NOTE: parameter in 'ctx' should be paired as key-value mapping. For example,
//       to log an error with message:
//           logger.Error("some message", "error", err)
//       which is similar to loggers with context:
//           logger.Error("some message", map[string]interface{}{
//              "error": err,
//           })
type Logger interface {
	// Info logs info level logs.
	Trace(msg string, ctx ...interface{})
	Debug(msg string, ctx ...interface{})
	Info(msg string, ctx ...interface{})
	Warn(msg string, ctx ...interface{})
	Error(msg string, ctx ...interface{})
}

// NullLogger logs nothing.
type NullLogger struct{}

// Trace implements Logger interface.
func (logger *NullLogger) Trace(msg string, ctx ...interface{}) {
}

// Debug implements Logger interface.
func (logger *NullLogger) Debug(msg string, ctx ...interface{}) {
}

// Info implements Logger interface.
func (logger *NullLogger) Info(msg string, ctx ...interface{}) {
}

// Warn implements Logger interface.
func (logger *NullLogger) Warn(msg string, ctx ...interface{}) {
}

// Error implements Logger interface.
func (logger *NullLogger) Error(msg string, ctx ...interface{}) {
}

// SimpleLogger logs everything.
type SimpleLogger struct{}

// composeVargs makes (msg, ctx...) could be pass to log.Println
func composeVargs(msg string, ctxs []interface{}) []interface{} {
	args := []interface{}{msg}
	for _, c := range ctxs {
		args = append(args, c)
	}
	return args
}

// Trace implements Logger interface.
func (logger *SimpleLogger) Trace(msg string, ctx ...interface{}) {
	log.Println(composeVargs(msg, ctx)...)
}

// Debug implements Logger interface.
func (logger *SimpleLogger) Debug(msg string, ctx ...interface{}) {
	log.Println(composeVargs(msg, ctx)...)
}

// Info implements Logger interface.
func (logger *SimpleLogger) Info(msg string, ctx ...interface{}) {
	log.Println(composeVargs(msg, ctx)...)
}

// Warn implements Logger interface.
func (logger *SimpleLogger) Warn(msg string, ctx ...interface{}) {
	log.Println(composeVargs(msg, ctx)...)
}

// Error implements Logger interface.
func (logger *SimpleLogger) Error(msg string, ctx ...interface{}) {
	log.Println(composeVargs(msg, ctx)...)
}

// CustomLogger logs everything.
type CustomLogger struct {
	logger *log.Logger
}

// NewCustomLogger creates a new custom logger.
func NewCustomLogger(logger *log.Logger) *CustomLogger {
	return &CustomLogger{
		logger: logger,
	}
}

// Trace implements Logger interface.
func (logger *CustomLogger) Trace(msg string, ctx ...interface{}) {
	logger.logger.Println(composeVargs(msg, ctx)...)
}

// Debug implements Logger interface.
func (logger *CustomLogger) Debug(msg string, ctx ...interface{}) {
	logger.logger.Println(composeVargs(msg, ctx)...)
}

// Info implements Logger interface.
func (logger *CustomLogger) Info(msg string, ctx ...interface{}) {
	logger.logger.Println(composeVargs(msg, ctx)...)
}

// Warn implements Logger interface.
func (logger *CustomLogger) Warn(msg string, ctx ...interface{}) {
	logger.logger.Println(composeVargs(msg, ctx)...)
}

// Error implements Logger interface.
func (logger *CustomLogger) Error(msg string, ctx ...interface{}) {
	logger.logger.Println(composeVargs(msg, ctx)...)
}
