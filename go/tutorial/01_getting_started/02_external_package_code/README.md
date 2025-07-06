Steps:
- Enable dependency tracking for your code: `go mod init example/external`
- Locate and download the modules that contains the package you imported: `go mod tidy`
- Run script: `go run external.go`