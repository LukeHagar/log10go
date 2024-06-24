# Completions
(*Completions*)

## Overview

Completions

### Available Operations

* [Create](#create) - Create a completion
* [Update](#update) - Update completion by id.
* [ListUngraded](#listungraded) - List ungraded completions i.e. completions that have not been associated with feedback but matches task selector.

## Create

Create a completion

### Example Usage

```go
package main

import(
	"github.com/log10-io/log10go"
	"github.com/log10-io/log10go/models/components"
	"context"
	"log"
)

func main() {
    s := log10go.New(
        log10go.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    completion := components.Completion{
        OrganizationID: "<value>",
        Request: &components.CreateChatCompletionRequest{
            Messages: []components.ChatCompletionRequestMessage{
                components.CreateChatCompletionRequestMessageChatCompletionRequestAssistantMessage(
                    components.ChatCompletionRequestAssistantMessage{
                        Role: components.ChatCompletionRequestAssistantMessageRoleAssistant,
                    },
                ),
            },
            Model: components.CreateModelTwo(
            components.TwoGpt4Turbo,
            ),
            N: log10go.Int64(1),
            ResponseFormat: &components.ResponseFormat{
                Type: components.CreateChatCompletionRequestTypeJSONObject.ToPointer(),
            },
            Temperature: log10go.Float64(1),
            TopP: log10go.Float64(1),
            User: log10go.String("user-1234"),
        },
    }

    var xLog10Organization *string = log10go.String("<value>")
    ctx := context.Background()
    res, err := s.Completions.Create(ctx, completion, xLog10Organization)
    if err != nil {
        log.Fatal(err)
    }
    if res.Any != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `completion`                                                   | [components.Completion](../../models/components/completion.md) | :heavy_check_mark:                                             | N/A                                                            |
| `xLog10Organization`                                           | **string*                                                      | :heavy_minus_sign:                                             | N/A                                                            |


### Response

**[*operations.CreateResponse](../../models/operations/createresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Update

Update completion by id.

### Example Usage

```go
package main

import(
	"github.com/log10-io/log10go"
	"github.com/log10-io/log10go/models/components"
	"context"
	"log"
)

func main() {
    s := log10go.New(
        log10go.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    var completionID string = "<value>"

    completion := components.Completion{
        OrganizationID: "<value>",
        Request: &components.CreateChatCompletionRequest{
            Messages: []components.ChatCompletionRequestMessage{
                components.CreateChatCompletionRequestMessageChatCompletionRequestFunctionMessage(
                    components.ChatCompletionRequestFunctionMessage{
                        Role: components.ChatCompletionRequestFunctionMessageRoleFunction,
                        Content: "<value>",
                        Name: "<value>",
                    },
                ),
            },
            Model: components.CreateModelTwo(
            components.TwoGpt4Turbo,
            ),
            N: log10go.Int64(1),
            ResponseFormat: &components.ResponseFormat{
                Type: components.CreateChatCompletionRequestTypeJSONObject.ToPointer(),
            },
            Temperature: log10go.Float64(1),
            TopP: log10go.Float64(1),
            User: log10go.String("user-1234"),
        },
    }

    var xLog10Organization *string = log10go.String("<value>")
    ctx := context.Background()
    res, err := s.Completions.Update(ctx, completionID, completion, xLog10Organization)
    if err != nil {
        log.Fatal(err)
    }
    if res.Completion != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `completionID`                                                 | *string*                                                       | :heavy_check_mark:                                             | The completion id to update.                                   |
| `completion`                                                   | [components.Completion](../../models/components/completion.md) | :heavy_check_mark:                                             | N/A                                                            |
| `xLog10Organization`                                           | **string*                                                      | :heavy_minus_sign:                                             | N/A                                                            |


### Response

**[*operations.UpdateResponse](../../models/operations/updateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListUngraded

List ungraded completions i.e. completions that have not been associated with feedback but matches task selector.

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
    res, err := s.Completions.ListUngraded(ctx, xLog10Organization)
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

**[*operations.ListUngradedResponse](../../models/operations/listungradedresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
