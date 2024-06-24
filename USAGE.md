<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"github.com/log10-io/log10go"
	"log"
)

func main() {
	s := log10go.New(
		log10go.WithSecurity("<YOUR_API_KEY_HERE>"),
	)
	var xLog10Organization *string = log10go.String("<value>")
	ctx := context.Background()
	res, err := s.Sessions.Create(ctx, xLog10Organization)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->