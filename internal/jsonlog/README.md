# JSON Log Package

This folder contains the JSON logging implementation for the Kelodi API. The JSON Log package provides structured logging in JSON format, which is useful for logging in production environments where logs are often consumed by log aggregation and analysis tools.

## Files

### [jsonlog.go](jsonlog.go)

- **Logger Implementation**: Defines the `Logger` struct and methods for structured logging in JSON format. Includes methods for logging at different levels (info, error, fatal) and for writing log entries to an output destination.
