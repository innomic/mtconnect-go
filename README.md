# mtconnect-go

**mtconnect-go** is a Go language library that provides an efficient way to work with the MTConnect standard. This library facilitates the generation of Go code from the XML Schema provided by various versions of MTConnect, enabling developers to easily integrate MTConnect data streams into their Go applications.

## Features

- **Version-Specific Packages**: **mtconnect-go** includes prepared packages for each version of the MTConnect standard, ensuring compatibility and ease of use across different versions.
- **Go Structs for XML Unmarshalling**: All types are defined as Go structs, allowing for straightforward XML unmarshalling. This feature makes it simple to work with MTConnect data in Go applications.

## Generate Codes

```
make codegen VERSION=2.1
```

## Usage

To use **mtconnect-go**, simply import the package corresponding to the version of the MTConnect standard you are working with. Below is an example demonstrating how to use the library with the MTConnect v2.2 schema:

```go
import (
  "encoding/xml"
  "net/http"
  "io"
  "github.com/innomic/mtconnect-go/schema/v2.2/mtstreams"
)

func main() {
    resp, err := http.Get("https://mtconnect.trakhound.com/current")
    if err != nil {
        // handle error
    }
    defer resp.Body.Close()

    respBodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }

    var unmarshalled mtstreams.MTConnectStreamsType
    err = xml.Unmarshal(respBodyBytes, &unmarshalled)
    if err != nil {
        // handle error
    }

    // Use unmarshalled data
}
```

## Example Struct

Here is an example of a struct representing the MTConnectStreamsType, which can be unmarshalled from XML:

```go
type MTConnectStreamsType struct {
    Header  *HeaderType  `xml:"Header"`
    Streams *StreamsType `xml:"Streams"`
}
```

## Roadmap

- **Client SDK Development**: Future plans include the development of a client SDK to streamline the process of connecting to MTConnect agents and retrieving data.
- **Adapter SDK Development**: An adapter SDK is also on the roadmap, which will assist in the creation of MTConnect adapters for various types of equipment and data sources.

## Contributing

Contributions to mtconnect-go are highly welcome! Whether it's improving the code, adding new features, or enhancing documentation, your input is valuable. Please feel free to fork the repository, make changes, and submit pull requests.

## License

The XML schema is owned by [MTConnect Institute](https://github.com/mtconnect).

**mtconnect-go** is licensed under the Apache License, Version 2.0, consistent with [MTConnect schema's licensing](https://github.com/mtconnect/schema/blob/master/LICENSE.TXT).

## Links

- https://schemas.mtconnect.org/
- https://github.com/mtconnect
- https://github.com/mtconnect/schema
- https://github.com/TrakHound/MTConnect.NET
