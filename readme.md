# fin-proto-go

A Go implementation of financial protocol codecs for various Chinese exchanges including SSE (Shanghai Stock Exchange), SZSE (Shenzhen Stock Exchange), and BSE (Beijing Stock Exchange), as well as risk control systems.

## Overview

This repository contains auto-generated Go code for encoding and decoding binary messages used in Chinese financial exchanges. The code is generated from protocol definition files using a custom protocol compiler (`fin-protoc`).

This repository has been applied to the [gt-auto](https://github.com/xinchentechnote/gt-auto) repository, a automated testing tool for financial systems(gateway, engine and so on).

## Structure

The repository is organized into several modules, each targeting a specific exchange or system:

- `bjse-trade-bin` - Beijing Stock Exchange trading protocols
- `codec` - Core encoding/decoding utilities and interfaces
- `risk-bin` - Risk control system protocols
- `sample-bin` - Sample protocol implementations
- `sse-bin` - Shanghai Stock Exchange protocols
- `szse-bin` - Shenzhen Stock Exchange protocols

## Key Features

- **Binary Message Encoding/Decoding**: Efficient serialization and deserialization of financial messages
- **Protocol Coverage**: Support for multiple Chinese exchanges (SSE, SZSE, BSE)
- **Auto-generated Code**: All message structures and codecs are generated from protocol definitions
- **Extensible**: Easy to add new protocols or extend existing ones
- **Checksum Support**: Built-in support for various checksum algorithms (CRC16, CRC32, exchange-specific)

## Installation

```bash
go get github.com/xinchentechnote/fin-proto-go
```

## Usage

Each exchange module contains message definitions and codecs. Here's a basic example using SSE messages:

```go
import (
    "bytes"
    "github.com/xinchentechnote/fin-proto-go/sse-bin/messages"
)

// Create and encode a message
logon := &messages.Logon{
    SenderCompId: "SENDER",
    TargetCompId: "TARGET",
    HeartBtInt:   30,
    PrtclVersion: "1.0",
    TradeDate:    20231201,
    Qsize:        100,
}

var buf bytes.Buffer
err := logon.Encode(&buf)
if err != nil {
    // handle error
}

// Decode a message
decoded := &messages.Logon{}
err = decoded.Decode(&buf)
if err != nil {
    // handle error
}
```

## Building

Each module contains a Makefile for building and testing:

```bash
# Build a specific module
cd sse-bin
make build

# Run tests
make test

# Format code
make fmt

# Run all targets
make all
```

To compile the protocol definitions from scratch:

```bash
make compile
```

This requires the `fin-protoc` compiler to be available in your PATH.
For more information about `fin-protoc`, see the [fin-protoc repository](https://github.com/xinchentechnote/fin-protoc).

## Testing

Run tests for all modules:

```bash
./build.sh
```

Or run tests for a specific module:

```bash
cd sse-bin
go test ./...
```

## Protocol Definitions

The actual protocol definitions are maintained in a separate repository and included as submodules in the `submodules/fin-proto` directory. These definitions are used by the `fin-protoc` compiler to generate the Go code.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
