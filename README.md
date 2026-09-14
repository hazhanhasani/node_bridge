# BluePanel Node Bridge (Go)

Go bridge used by **BluePanel** integrations to communicate with **BluePanel Node** over gRPC or REST.

## Installation

```bash
go get github.com/hazhanhasani/node_bridge@main
```

## Usage example

```go
package main

import (
    "fmt"
    "log"

    "github.com/google/uuid"
    bridge "github.com/hazhanhasani/node_bridge"
    "github.com/hazhanhasani/node_bridge/common"
)

func main() {
    apiKey := uuid.New()

    node, err := bridge.New("127.0.0.1", bridge.GRPC,
        bridge.WithPort(62050),
        bridge.WithAPIKey(apiKey),
    )
    if err != nil {
        log.Fatal(err)
    }

    config := `{"inbounds": [], "outbounds": []}`
    if err := node.Start(config, common.BackendType_XRAY, nil, 60); err != nil {
        log.Fatal(err)
    }
    defer node.Stop()

    info, err := node.Info()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("BluePanel Node version: %s\n", info.NodeVersion)
}
```

This fork is maintained as part of the BluePanel stack and should be used instead of upstream bridge sources by BluePanel components.
