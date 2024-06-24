# github.com/log10-io/log10go

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


## 🏗 **Welcome to your new SDK!** 🏗

It has been generated successfully based on your OpenAPI spec. However, it is not yet ready for production use. Here are some next steps:
- [ ] 🛠 Make your SDK feel handcrafted by [customizing it](https://www.speakeasyapi.dev/docs/customize-sdks)
- [ ] ♻️ Refine your SDK quickly by iterating locally with the [Speakeasy CLI](https://github.com/speakeasy-api/speakeasy)
- [ ] 🎁 Publish your SDK to package managers by [configuring automatic publishing](https://www.speakeasyapi.dev/docs/advanced-setup/publish-sdks)
- [ ] ✨ When ready to productionize, delete this section from the README

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/log10-io/log10go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

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

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Completions](docs/sdks/completions/README.md)

* [Create](docs/sdks/completions/README.md#create) - Create a completion
* [Update](docs/sdks/completions/README.md#update) - Update completion by id.
* [ListUngraded](docs/sdks/completions/README.md#listungraded) - List ungraded completions i.e. completions that have not been associated with feedback but matches task selector.

### [Sessions](docs/sdks/sessions/README.md)

* [Create](docs/sdks/sessions/README.md#create) - Create a session

### [Feedback](docs/sdks/feedback/README.md)

* [Get](docs/sdks/feedback/README.md#get) - Fetch feedback by id.
* [List](docs/sdks/feedback/README.md#list) - List feedback
* [Upload](docs/sdks/feedback/README.md#upload) - Upload a piece of feedback

### [FeedbackTasks](docs/sdks/feedbacktasks/README.md)

* [List](docs/sdks/feedbacktasks/README.md#list) - List feedback tasks.
* [Create](docs/sdks/feedbacktasks/README.md#create) - Create a new task.
* [Get](docs/sdks/feedbacktasks/README.md#get) - Retrieves feedback task `taskId`.
<!-- End Available Resources and Operations [operations] -->

<!-- Start Global Parameters [global-parameters] -->
## Global Parameters

A parameter is configured globally. This parameter may be set on the SDK client instance itself during initialization. When configured as an option during SDK initialization, This global value will be used as the default on the operations that use it. When such operations are called, there is a place in each to override the global value, if needed.

For example, you can set `X-Log10-Organization` to `"<value>"` at SDK initialization and then you do not have to pass the same value on calls to operations like `Update`. But if you want to do so you may, which will locally override the global setting. See the example code below for a demonstration.


### Available Globals

The following global parameter is available.

| Name | Type | Required | Description |
| ---- | ---- |:--------:| ----------- |
| XLog10Organization | string |  | The XLog10Organization parameter. |


### Example

```go
package main

import (
	"context"
	"github.com/log10-io/log10go"
	"github.com/log10-io/log10go/models/components"
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
						Role:    components.ChatCompletionRequestFunctionMessageRoleFunction,
						Content: "<value>",
						Name:    "<value>",
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
			TopP:        log10go.Float64(1),
			User:        log10go.String("user-1234"),
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
<!-- End Global Parameters [global-parameters] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

### Example

```go
package main

import (
	"context"
	"errors"
	"github.com/log10-io/log10go"
	"github.com/log10-io/log10go/models/components"
	"github.com/log10-io/log10go/models/sdkerrors"
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
			TopP:        log10go.Float64(1),
			User:        log10go.String("user-1234"),
		},
	}

	var xLog10Organization *string = log10go.String("<value>")
	ctx := context.Background()
	res, err := s.Completions.Create(ctx, completion, xLog10Organization)
	if err != nil {

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://log10.io` | None |

#### Example

```go
package main

import (
	"context"
	"github.com/log10-io/log10go"
	"github.com/log10-io/log10go/models/components"
	"log"
)

func main() {
	s := log10go.New(
		log10go.WithServerIndex(0),
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
			TopP:        log10go.Float64(1),
			User:        log10go.String("user-1234"),
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


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"github.com/log10-io/log10go"
	"github.com/log10-io/log10go/models/components"
	"log"
)

func main() {
	s := log10go.New(
		log10go.WithServerURL("https://log10.io"),
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
			TopP:        log10go.Float64(1),
			User:        log10go.String("user-1234"),
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
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type         | Scheme       |
| ------------ | ------------ | ------------ |
| `Log10Token` | apiKey       | API key      |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"github.com/log10-io/log10go"
	"github.com/log10-io/log10go/models/components"
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
			TopP:        log10go.Float64(1),
			User:        log10go.String("user-1234"),
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
<!-- End Authentication [security] -->

<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
