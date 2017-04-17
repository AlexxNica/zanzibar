// Code generated by zanzibar
// @generated

package contacts

import (
	"context"

	"github.com/uber/zanzibar/examples/example-gateway/build/clients"
	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"

	endpointsContactsContacts "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/contacts/contacts"
	customContacts "github.com/uber/zanzibar/examples/example-gateway/endpoints/contacts"
)

// HandleSaveContactsRequest handles "/contacts/:userUUID/contacts".
func HandleSaveContactsRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
	clients *clients.Clients,
) {
	var requestBody endpointsContactsContacts.SaveContactsRequest
	if ok := req.ReadAndUnmarshalBody(&requestBody); !ok {
		return
	}

	// TODO(sindelar): Switch to zanzibar.Headers when tchannel
	// generation is implemented.
	headers := map[string]string{}
	for k, v := range map[string]string{} {
		headers[v] = req.Header.Get(k)
	}

	workflow := customContacts.SaveContactsEndpoint{
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

	res.WriteJSON(202, respHeaders, response)
}
