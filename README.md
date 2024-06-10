# Calculator - Basics of go

## Prerequisites

1. Install [Go](https://go.dev/doc/install)

    For os / linux edit `~/.bashrc` or `~/.zshrc`:

    ```sh
    export GOROOT=/usr/local/go
    export GOPATH=$HOME/go
    export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
    ```

    To have above environment variables in current terminal session use command:

    ```sh
    source ~/.bashrc 
    ```

    Verify installation with:

    ```sh
    go version
    ```

2. Install [Visual Studio Code](https://code.visualstudio.com/download)

    In VSCode choose `extensions` from left menu and install `golang.go`

3. Install Go tools

    ```sh
    go install golang.org/x/lint/golint@latest
    go install golang.org/x/tools/cmd/goimports@latest
    go install golang.org/x/tools/gopls@latest
    go install honnef.co/go/tools/cmd/staticcheck@latest
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
    ```
