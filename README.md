# examplesGo
 go mod init github.com/ETISDev/examplesGo/convert-images/png-to-webp

# Install the latest version of a program,
# ignoring go.mod in the current directory (if any). 
$ go install golang.org/x/tools/gopls@latest

# Install a specific version of a program. 
$ go install golang.org/x/tools/gopls@v0.6.4

# Install a program at the version selected by the module in the current directory. 
$ go install golang.org/x/tools/gopls

# Install all programs in a directory. 
$ go install ./cmd/...

To install any other package, simply:
go install <package_name>@<version>

go mod tidy