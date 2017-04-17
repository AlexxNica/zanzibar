// Code generated by zanzibar
// @generated

package bar

import (
	"context"

	"github.com/uber/zanzibar/examples/example-gateway/build/clients"
	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"

	"github.com/uber/zanzibar/examples/example-gateway/build/clients/bar"
	clientsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/bar/bar"
	endpointsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/bar/bar"
)

// HandleNormalRequest handles "/bar/bar-path".
func HandleNormalRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
	clients *clients.Clients,
) {
	var requestBody NormalHTTPRequest
	if ok := req.ReadAndUnmarshalBody(&requestBody); !ok {
		return
	}

	// TODO(sindelar): Switch to zanzibar.Headers when tchannel
	// generation is implemented.
	headers := map[string]string{}
	for k, v := range map[string]string{} {
		headers[v] = req.Header.Get(k)
	}

	workflow := NormalEndpoint{
		Clients: clients,
		Logger:  req.Logger,
		Request: req,
	}

	response, respHeaders, err := workflow.Handle(ctx, headers, &requestBody)
	if err != nil {
		req.Logger.Warn("Workflow for endpoint returned error",
			zap.String("error", err.Error()),
		)
		res.SendErrorString(500, "Unexpected server error")
		return
	}

	res.WriteJSON(200, respHeaders, response)
}

// NormalEndpoint calls thrift client Bar.Normal
type NormalEndpoint struct {
	Clients *clients.Clients
	Logger  *zap.Logger
	Request *zanzibar.ServerHTTPRequest
}

// Handle calls thrift client.
func (w NormalEndpoint) Handle(
	ctx context.Context,
	headers map[string]string,
	r *NormalHTTPRequest,
) (*endpointsBarBar.BarResponse, map[string]string, error) {
	clientRequest := convertToNormalClientRequest(r)

	clientHeaders := map[string]string{}
	for k, v := range map[string]string{} {
		headers[v] = headers[k]
	}

	clientRespBody, respHeaders, err := w.Clients.Bar.Normal(
		ctx, clientHeaders, clientRequest,
	)
	if err != nil {
		w.Logger.Warn("Could not make client request",
			zap.String("error", err.Error()),
		)
		return nil, nil, err
	}

	endRespHead := map[string]string{}
	for k, v := range map[string]string{} {
		endRespHead[v] = respHeaders[k]
	}

	response := convertNormalClientResponse(clientRespBody)
	return response, endRespHead, nil
}

func convertToNormalClientRequest(body *NormalHTTPRequest) *barClient.NormalHTTPRequest {
	clientRequest := &barClient.NormalHTTPRequest{}

	clientRequest.Request = (*clientsBarBar.BarRequest)(body.Request)

	return clientRequest
}
func convertNormalClientResponse(body *clientsBarBar.BarResponse) *endpointsBarBar.BarResponse {
	// TODO: Add response fields mapping here.
	downstreamResponse := &endpointsBarBar.BarResponse{}
	return downstreamResponse
}
