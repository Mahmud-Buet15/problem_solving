## Tutorial: Get started with Go
- Link: https://go.dev/doc/tutorial/getting-started

## Notes:
- `fmt` package contains functions for formatting text, including printing to the console. This package is one of the standard library packages you got when you installed Go.
- Get list of go command: `go help`
    - Get more information on a command: `go help <command>`
- Enable dependency tracking for your code: `go mod init example/hello`
- Run your code: `go run .`  or `go run <script_name>`
- When you run `go mod tidy`, it locates and downloads the `rsc.io/<module_name>` module that contains the package you imported. By default, it downloads the latest version