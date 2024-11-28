package web

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// NoResponse tells the Respond function to not respond to the request. In these
// cases the app layer code has already done so.
type NoResponse struct{}

// NewNoResponse constructs a no reponse value.
func NewNoResponse() NoResponse {
	return NoResponse{}
}

// Encode implements the Encoder interface
func (NoResponse) Encode() ([]byte, string, error) {
	return nil, "", nil
}

// =============================================================================

type httpStatus interface {
	HTTPStatus() int
}

// Respond sends a response to the client.
func Respond(ctx context.Context, w http.ResponseWriter, data any, statusCode int) error {
	//if _, ok := dataModel.(NoResponse); ok {
	//	return nil
	//}

	// If the context has been canceled, it means the client is no longer
	// waiting for a response.
	if err := ctx.Err(); err != nil {
		if errors.Is(err, context.Canceled) {
			return errors.New("client disconnected, do not send response")
		}
	}

	setStatusCode(ctx, statusCode)

	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
		return nil
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return fmt.Errorf("respond: encode: %w", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if _, err := w.Write(jsonData); err != nil {
		return fmt.Errorf("respond: write: %w", err)
	}

	return nil
}
