# FeedbackTasks
(*FeedbackTasks*)

## Overview

FeedbackTasks

### Available Operations

* [List](#list) - List feedback tasks.
* [Create](#create) - Create a new task.
* [Get](#get) - Retrieves feedback task `taskId`.

## List

List feedback tasks.

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

    ctx := context.Background()
    res, err := s.FeedbackTasks.List(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Tasks != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListFeedbackTasksResponse](../../models/operations/listfeedbacktasksresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Create

Create a new task.

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
    var request *components.Task = &components.Task{
        JSONSchema: components.JSONSchema{},
        Name: "<value>",
        Instruction: "<value>",
        CompletionTagsSelector: components.CompletionTagsSelector{},
    }
    ctx := context.Background()
    res, err := s.FeedbackTasks.Create(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Task != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [components.Task](../../models/components/task.md)    | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateFeedbackTaskResponse](../../models/operations/createfeedbacktaskresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Get

Retrieves feedback task `taskId`.

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
    var taskID string = "<value>"
    ctx := context.Background()
    res, err := s.FeedbackTasks.Get(ctx, taskID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Task != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `taskID`                                              | *string*                                              | :heavy_check_mark:                                    | The task id to fetch.                                 |


### Response

**[*operations.GetFeedbackTaskResponse](../../models/operations/getfeedbacktaskresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
