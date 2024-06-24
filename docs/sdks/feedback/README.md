# Feedback
(*Feedback*)

## Overview

Feedback

### Available Operations

* [Get](#get) - Fetch feedback by id.
* [List](#list) - List feedback
* [Upload](#upload) - Upload a piece of feedback

## Get

Fetch feedback by id.

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
    var feedbackID string = "<value>"

    var xLog10Organization *string = log10go.String("<value>")
    ctx := context.Background()
    res, err := s.Feedback.Get(ctx, feedbackID, xLog10Organization)
    if err != nil {
        log.Fatal(err)
    }
    if res.Feedback != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `feedbackID`                                          | *string*                                              | :heavy_check_mark:                                    | The feedback id to fetch.                             |
| `xLog10Organization`                                  | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   |


### Response

**[*operations.GetResponse](../../models/operations/getresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## List

List feedback

### Example Usage

```go
package main

import(
	"github.com/log10-io/log10go"
	"github.com/log10-io/log10go/models/operations"
	"context"
	"log"
)

func main() {
    s := log10go.New(
        log10go.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    var xLog10Organization *string = log10go.String("<value>")

    var requestBody *operations.ListRequestBody = &operations.ListRequestBody{}
    ctx := context.Background()
    res, err := s.Feedback.List(ctx, xLog10Organization, requestBody)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                 | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `ctx`                                                                     | [context.Context](https://pkg.go.dev/context#Context)                     | :heavy_check_mark:                                                        | The context to use for the request.                                       |
| `xLog10Organization`                                                      | **string*                                                                 | :heavy_minus_sign:                                                        | N/A                                                                       |
| `requestBody`                                                             | [*operations.ListRequestBody](../../models/operations/listrequestbody.md) | :heavy_minus_sign:                                                        | N/A                                                                       |


### Response

**[*operations.ListResponse](../../models/operations/listresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Upload

Upload a piece of feedback

### Example Usage

```go
package main

import(
	"github.com/log10-io/log10go"
	"github.com/log10-io/log10go/models/operations"
	"context"
	"log"
)

func main() {
    s := log10go.New(
        log10go.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    var requestBody operations.UploadRequestBody = operations.CreateUploadRequestBodyOne(
            operations.One{
                TaskID: "<value>",
                JSONValues: operations.JSONValues{},
                MatchedCompletionIds: []string{
                    "<value>",
                },
                Comment: "The slim & simple Maple Gaming Keyboard from Dev Byte comes with a sleek body and 7- Color RGB LED Back-lighting for smart functionality",
                CompletionTagsSelector: []string{
                    "<value>",
                },
            },
    )

    var xLog10Organization *string = log10go.String("<value>")
    ctx := context.Background()
    res, err := s.Feedback.Upload(ctx, requestBody, xLog10Organization)
    if err != nil {
        log.Fatal(err)
    }
    if res.Feedback != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `requestBody`                                                                | [operations.UploadRequestBody](../../models/operations/uploadrequestbody.md) | :heavy_check_mark:                                                           | N/A                                                                          |
| `xLog10Organization`                                                         | **string*                                                                    | :heavy_minus_sign:                                                           | N/A                                                                          |


### Response

**[*operations.UploadResponse](../../models/operations/uploadresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
