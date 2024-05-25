// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/log10-io/log10go/internal/utils"
)

type Two string

const (
	TwoGpt4Turbo         Two = "gpt-4-turbo"
	TwoGpt4Turbo20240409 Two = "gpt-4-turbo-2024-04-09"
	TwoGpt40125Preview   Two = "gpt-4-0125-preview"
	TwoGpt4TurboPreview  Two = "gpt-4-turbo-preview"
	TwoGpt41106Preview   Two = "gpt-4-1106-preview"
	TwoGpt4VisionPreview Two = "gpt-4-vision-preview"
	TwoGpt4              Two = "gpt-4"
	TwoGpt40314          Two = "gpt-4-0314"
	TwoGpt40613          Two = "gpt-4-0613"
	TwoGpt432k           Two = "gpt-4-32k"
	TwoGpt432k0314       Two = "gpt-4-32k-0314"
	TwoGpt432k0613       Two = "gpt-4-32k-0613"
	TwoGpt35Turbo        Two = "gpt-3.5-turbo"
	TwoGpt35Turbo16k     Two = "gpt-3.5-turbo-16k"
	TwoGpt35Turbo0301    Two = "gpt-3.5-turbo-0301"
	TwoGpt35Turbo0613    Two = "gpt-3.5-turbo-0613"
	TwoGpt35Turbo1106    Two = "gpt-3.5-turbo-1106"
	TwoGpt35Turbo0125    Two = "gpt-3.5-turbo-0125"
	TwoGpt35Turbo16k0613 Two = "gpt-3.5-turbo-16k-0613"
)

func (e Two) ToPointer() *Two {
	return &e
}
func (e *Two) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "gpt-4-turbo":
		fallthrough
	case "gpt-4-turbo-2024-04-09":
		fallthrough
	case "gpt-4-0125-preview":
		fallthrough
	case "gpt-4-turbo-preview":
		fallthrough
	case "gpt-4-1106-preview":
		fallthrough
	case "gpt-4-vision-preview":
		fallthrough
	case "gpt-4":
		fallthrough
	case "gpt-4-0314":
		fallthrough
	case "gpt-4-0613":
		fallthrough
	case "gpt-4-32k":
		fallthrough
	case "gpt-4-32k-0314":
		fallthrough
	case "gpt-4-32k-0613":
		fallthrough
	case "gpt-3.5-turbo":
		fallthrough
	case "gpt-3.5-turbo-16k":
		fallthrough
	case "gpt-3.5-turbo-0301":
		fallthrough
	case "gpt-3.5-turbo-0613":
		fallthrough
	case "gpt-3.5-turbo-1106":
		fallthrough
	case "gpt-3.5-turbo-0125":
		fallthrough
	case "gpt-3.5-turbo-16k-0613":
		*e = Two(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Two: %v", v)
	}
}

type ModelType string

const (
	ModelTypeStr ModelType = "str"
	ModelTypeTwo ModelType = "2"
)

// Model - ID of the model to use. See the [model endpoint compatibility](/docs/models/model-endpoint-compatibility) table for details on which models work with the Chat API.
type Model struct {
	Str *string
	Two *Two

	Type ModelType
}

func CreateModelStr(str string) Model {
	typ := ModelTypeStr

	return Model{
		Str:  &str,
		Type: typ,
	}
}

func CreateModelTwo(two Two) Model {
	typ := ModelTypeTwo

	return Model{
		Two:  &two,
		Type: typ,
	}
}

func (u *Model) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = ModelTypeStr
		return nil
	}

	var two Two = Two("")
	if err := utils.UnmarshalJSON(data, &two, "", true, true); err == nil {
		u.Two = &two
		u.Type = ModelTypeTwo
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Model", string(data))
}

func (u Model) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Two != nil {
		return utils.MarshalJSON(u.Two, "", true)
	}

	return nil, errors.New("could not marshal union type Model: all fields are null")
}

// CreateChatCompletionRequestType - Must be one of `text` or `json_object`.
type CreateChatCompletionRequestType string

const (
	CreateChatCompletionRequestTypeText       CreateChatCompletionRequestType = "text"
	CreateChatCompletionRequestTypeJSONObject CreateChatCompletionRequestType = "json_object"
)

func (e CreateChatCompletionRequestType) ToPointer() *CreateChatCompletionRequestType {
	return &e
}
func (e *CreateChatCompletionRequestType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text":
		fallthrough
	case "json_object":
		*e = CreateChatCompletionRequestType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateChatCompletionRequestType: %v", v)
	}
}

// ResponseFormat - An object specifying the format that the model must output. Compatible with [GPT-4 Turbo](/docs/models/gpt-4-and-gpt-4-turbo) and all GPT-3.5 Turbo models newer than `gpt-3.5-turbo-1106`.
//
// Setting to `{ "type": "json_object" }` enables JSON mode, which guarantees the message the model generates is valid JSON.
//
// **Important:** when using JSON mode, you **must** also instruct the model to produce JSON yourself via a system or user message. Without this, the model may generate an unending stream of whitespace until the generation reaches the token limit, resulting in a long-running and seemingly "stuck" request. Also note that the message content may be partially cut off if `finish_reason="length"`, which indicates the generation exceeded `max_tokens` or the conversation exceeded the max context length.
type ResponseFormat struct {
	// Must be one of `text` or `json_object`.
	Type *CreateChatCompletionRequestType `default:"text" json:"type"`
}

func (r ResponseFormat) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *ResponseFormat) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ResponseFormat) GetType() *CreateChatCompletionRequestType {
	if o == nil {
		return nil
	}
	return o.Type
}

type StopType string

const (
	StopTypeStr        StopType = "str"
	StopTypeArrayOfStr StopType = "arrayOfStr"
)

// Stop - Up to 4 sequences where the API will stop generating further tokens.
type Stop struct {
	Str        *string
	ArrayOfStr []string

	Type StopType
}

func CreateStopStr(str string) Stop {
	typ := StopTypeStr

	return Stop{
		Str:  &str,
		Type: typ,
	}
}

func CreateStopArrayOfStr(arrayOfStr []string) Stop {
	typ := StopTypeArrayOfStr

	return Stop{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *Stop) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = StopTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = StopTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Stop", string(data))
}

func (u Stop) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type Stop: all fields are null")
}

// One - `none` means the model will not call a function and instead generates a message. `auto` means the model can pick between generating a message or calling a function.
type One string

const (
	OneNone One = "none"
	OneAuto One = "auto"
)

func (e One) ToPointer() *One {
	return &e
}
func (e *One) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "auto":
		*e = One(v)
		return nil
	default:
		return fmt.Errorf("invalid value for One: %v", v)
	}
}

type CreateChatCompletionRequestFunctionCallType string

const (
	CreateChatCompletionRequestFunctionCallTypeOne                              CreateChatCompletionRequestFunctionCallType = "1"
	CreateChatCompletionRequestFunctionCallTypeChatCompletionFunctionCallOption CreateChatCompletionRequestFunctionCallType = "ChatCompletionFunctionCallOption"
)

// CreateChatCompletionRequestFunctionCall - Deprecated in favor of `tool_choice`.
//
// Controls which (if any) function is called by the model.
// `none` means the model will not call a function and instead generates a message.
// `auto` means the model can pick between generating a message or calling a function.
// Specifying a particular function via `{"name": "my_function"}` forces the model to call that function.
//
// `none` is the default when no functions are present. `auto` is the default if functions are present.
//
// Deprecated type: This will be removed in a future release, please migrate away from it as soon as possible.
type CreateChatCompletionRequestFunctionCall struct {
	One                              *One
	ChatCompletionFunctionCallOption *ChatCompletionFunctionCallOption

	Type CreateChatCompletionRequestFunctionCallType
}

func CreateCreateChatCompletionRequestFunctionCallOne(one One) CreateChatCompletionRequestFunctionCall {
	typ := CreateChatCompletionRequestFunctionCallTypeOne

	return CreateChatCompletionRequestFunctionCall{
		One:  &one,
		Type: typ,
	}
}

func CreateCreateChatCompletionRequestFunctionCallChatCompletionFunctionCallOption(chatCompletionFunctionCallOption ChatCompletionFunctionCallOption) CreateChatCompletionRequestFunctionCall {
	typ := CreateChatCompletionRequestFunctionCallTypeChatCompletionFunctionCallOption

	return CreateChatCompletionRequestFunctionCall{
		ChatCompletionFunctionCallOption: &chatCompletionFunctionCallOption,
		Type:                             typ,
	}
}

func (u *CreateChatCompletionRequestFunctionCall) UnmarshalJSON(data []byte) error {

	var chatCompletionFunctionCallOption ChatCompletionFunctionCallOption = ChatCompletionFunctionCallOption{}
	if err := utils.UnmarshalJSON(data, &chatCompletionFunctionCallOption, "", true, true); err == nil {
		u.ChatCompletionFunctionCallOption = &chatCompletionFunctionCallOption
		u.Type = CreateChatCompletionRequestFunctionCallTypeChatCompletionFunctionCallOption
		return nil
	}

	var one One = One("")
	if err := utils.UnmarshalJSON(data, &one, "", true, true); err == nil {
		u.One = &one
		u.Type = CreateChatCompletionRequestFunctionCallTypeOne
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CreateChatCompletionRequestFunctionCall", string(data))
}

func (u CreateChatCompletionRequestFunctionCall) MarshalJSON() ([]byte, error) {
	if u.One != nil {
		return utils.MarshalJSON(u.One, "", true)
	}

	if u.ChatCompletionFunctionCallOption != nil {
		return utils.MarshalJSON(u.ChatCompletionFunctionCallOption, "", true)
	}

	return nil, errors.New("could not marshal union type CreateChatCompletionRequestFunctionCall: all fields are null")
}

type CreateChatCompletionRequest struct {
	// A list of messages comprising the conversation so far. [Example Python code](https://cookbook.openai.com/examples/how_to_format_inputs_to_chatgpt_models).
	Messages []ChatCompletionRequestMessage `json:"messages"`
	// ID of the model to use. See the [model endpoint compatibility](/docs/models/model-endpoint-compatibility) table for details on which models work with the Chat API.
	Model Model `json:"model"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.
	//
	// [See more information about frequency and presence penalties.](/docs/guides/text-generation/parameter-details)
	//
	FrequencyPenalty *float64 `default:"0" json:"frequency_penalty"`
	// Modify the likelihood of specified tokens appearing in the completion.
	//
	// Accepts a JSON object that maps tokens (specified by their token ID in the tokenizer) to an associated bias value from -100 to 100. Mathematically, the bias is added to the logits generated by the model prior to sampling. The exact effect will vary per model, but values between -1 and 1 should decrease or increase likelihood of selection; values like -100 or 100 should result in a ban or exclusive selection of the relevant token.
	//
	LogitBias map[string]int64 `json:"logit_bias,omitempty"`
	// Whether to return log probabilities of the output tokens or not. If true, returns the log probabilities of each output token returned in the `content` of `message`.
	Logprobs *bool `default:"false" json:"logprobs"`
	// An integer between 0 and 20 specifying the number of most likely tokens to return at each token position, each with an associated log probability. `logprobs` must be set to `true` if this parameter is used.
	TopLogprobs *int64 `json:"top_logprobs,omitempty"`
	// The maximum number of [tokens](/tokenizer) that can be generated in the chat completion.
	//
	// The total length of input tokens and generated tokens is limited by the model's context length. [Example Python code](https://cookbook.openai.com/examples/how_to_count_tokens_with_tiktoken) for counting tokens.
	//
	MaxTokens *int64 `json:"max_tokens,omitempty"`
	// How many chat completion choices to generate for each input message. Note that you will be charged based on the number of generated tokens across all of the choices. Keep `n` as `1` to minimize costs.
	N *int64 `default:"1" json:"n"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing the model's likelihood to talk about new topics.
	//
	// [See more information about frequency and presence penalties.](/docs/guides/text-generation/parameter-details)
	//
	PresencePenalty *float64 `default:"0" json:"presence_penalty"`
	// An object specifying the format that the model must output. Compatible with [GPT-4 Turbo](/docs/models/gpt-4-and-gpt-4-turbo) and all GPT-3.5 Turbo models newer than `gpt-3.5-turbo-1106`.
	//
	// Setting to `{ "type": "json_object" }` enables JSON mode, which guarantees the message the model generates is valid JSON.
	//
	// **Important:** when using JSON mode, you **must** also instruct the model to produce JSON yourself via a system or user message. Without this, the model may generate an unending stream of whitespace until the generation reaches the token limit, resulting in a long-running and seemingly "stuck" request. Also note that the message content may be partially cut off if `finish_reason="length"`, which indicates the generation exceeded `max_tokens` or the conversation exceeded the max context length.
	//
	ResponseFormat *ResponseFormat `json:"response_format,omitempty"`
	// This feature is in Beta.
	// If specified, our system will make a best effort to sample deterministically, such that repeated requests with the same `seed` and parameters should return the same result.
	// Determinism is not guaranteed, and you should refer to the `system_fingerprint` response parameter to monitor changes in the backend.
	//
	Seed *int64 `json:"seed,omitempty"`
	// Up to 4 sequences where the API will stop generating further tokens.
	//
	Stop *Stop `json:"stop,omitempty"`
	// If set, partial message deltas will be sent, like in ChatGPT. Tokens will be sent as data-only [server-sent events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events#Event_stream_format) as they become available, with the stream terminated by a `data: [DONE]` message. [Example Python code](https://cookbook.openai.com/examples/how_to_stream_completions).
	//
	Stream *bool `default:"false" json:"stream"`
	// Options for streaming response. Only set this when you set `stream: true`.
	//
	StreamOptions *ChatCompletionStreamOptions `json:"stream_options,omitempty"`
	// What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.
	//
	// We generally recommend altering this or `top_p` but not both.
	//
	Temperature *float64 `default:"1" json:"temperature"`
	// An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability mass are considered.
	//
	// We generally recommend altering this or `temperature` but not both.
	//
	TopP *float64 `default:"1" json:"top_p"`
	// A list of tools the model may call. Currently, only functions are supported as a tool. Use this to provide a list of functions the model may generate JSON inputs for. A max of 128 functions are supported.
	//
	Tools []ChatCompletionTool `json:"tools,omitempty"`
	// Controls which (if any) tool is called by the model.
	// `none` means the model will not call any tool and instead generates a message.
	// `auto` means the model can pick between generating a message or calling one or more tools.
	// `required` means the model must call one or more tools.
	// Specifying a particular tool via `{"type": "function", "function": {"name": "my_function"}}` forces the model to call that tool.
	//
	// `none` is the default when no tools are present. `auto` is the default if tools are present.
	//
	ToolChoice *ChatCompletionToolChoiceOption `json:"tool_choice,omitempty"`
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. [Learn more](/docs/guides/safety-best-practices/end-user-ids).
	//
	User *string `json:"user,omitempty"`
	// Deprecated in favor of `tool_choice`.
	//
	// Controls which (if any) function is called by the model.
	// `none` means the model will not call a function and instead generates a message.
	// `auto` means the model can pick between generating a message or calling a function.
	// Specifying a particular function via `{"name": "my_function"}` forces the model to call that function.
	//
	// `none` is the default when no functions are present. `auto` is the default if functions are present.
	//
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	FunctionCall *CreateChatCompletionRequestFunctionCall `json:"function_call,omitempty"`
	// Deprecated in favor of `tools`.
	//
	// A list of functions the model may generate JSON inputs for.
	//
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Functions []ChatCompletionFunctions `json:"functions,omitempty"`
}

func (c CreateChatCompletionRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateChatCompletionRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateChatCompletionRequest) GetMessages() []ChatCompletionRequestMessage {
	if o == nil {
		return []ChatCompletionRequestMessage{}
	}
	return o.Messages
}

func (o *CreateChatCompletionRequest) GetModel() Model {
	if o == nil {
		return Model{}
	}
	return o.Model
}

func (o *CreateChatCompletionRequest) GetFrequencyPenalty() *float64 {
	if o == nil {
		return nil
	}
	return o.FrequencyPenalty
}

func (o *CreateChatCompletionRequest) GetLogitBias() map[string]int64 {
	if o == nil {
		return nil
	}
	return o.LogitBias
}

func (o *CreateChatCompletionRequest) GetLogprobs() *bool {
	if o == nil {
		return nil
	}
	return o.Logprobs
}

func (o *CreateChatCompletionRequest) GetTopLogprobs() *int64 {
	if o == nil {
		return nil
	}
	return o.TopLogprobs
}

func (o *CreateChatCompletionRequest) GetMaxTokens() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxTokens
}

func (o *CreateChatCompletionRequest) GetN() *int64 {
	if o == nil {
		return nil
	}
	return o.N
}

func (o *CreateChatCompletionRequest) GetPresencePenalty() *float64 {
	if o == nil {
		return nil
	}
	return o.PresencePenalty
}

func (o *CreateChatCompletionRequest) GetResponseFormat() *ResponseFormat {
	if o == nil {
		return nil
	}
	return o.ResponseFormat
}

func (o *CreateChatCompletionRequest) GetSeed() *int64 {
	if o == nil {
		return nil
	}
	return o.Seed
}

func (o *CreateChatCompletionRequest) GetStop() *Stop {
	if o == nil {
		return nil
	}
	return o.Stop
}

func (o *CreateChatCompletionRequest) GetStream() *bool {
	if o == nil {
		return nil
	}
	return o.Stream
}

func (o *CreateChatCompletionRequest) GetStreamOptions() *ChatCompletionStreamOptions {
	if o == nil {
		return nil
	}
	return o.StreamOptions
}

func (o *CreateChatCompletionRequest) GetTemperature() *float64 {
	if o == nil {
		return nil
	}
	return o.Temperature
}

func (o *CreateChatCompletionRequest) GetTopP() *float64 {
	if o == nil {
		return nil
	}
	return o.TopP
}

func (o *CreateChatCompletionRequest) GetTools() []ChatCompletionTool {
	if o == nil {
		return nil
	}
	return o.Tools
}

func (o *CreateChatCompletionRequest) GetToolChoice() *ChatCompletionToolChoiceOption {
	if o == nil {
		return nil
	}
	return o.ToolChoice
}

func (o *CreateChatCompletionRequest) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *CreateChatCompletionRequest) GetFunctionCall() *CreateChatCompletionRequestFunctionCall {
	if o == nil {
		return nil
	}
	return o.FunctionCall
}

func (o *CreateChatCompletionRequest) GetFunctions() []ChatCompletionFunctions {
	if o == nil {
		return nil
	}
	return o.Functions
}
