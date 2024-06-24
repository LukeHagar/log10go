# Sessions
(*Sessions*)

## Overview

Sessions

### Available Operations

* [Create](#create) - Create a session

## Create

Create a session

### Example Usage

```go
package main

import(
	"github.com/log10-io/log10go"
	"context"
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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `xLog10Organization`                                  | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   |


### Response

**[*operations.CreateSessionResponse](../../models/operations/createsessionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
