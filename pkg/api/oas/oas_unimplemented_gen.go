// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// GetTonConnectPayload implements getTonConnectPayload operation.
//
// Get a challenge for TON Connect.
//
// GET /tonconnect/payload
func (UnimplementedHandler) GetTonConnectPayload(ctx context.Context) (r *GetTonConnectPayloadOK, _ error) {
	return r, ht.ErrNotImplemented
}

// SubscribeToAccountEvents implements subscribeToAccountEvents operation.
//
// Subscribe to notifications about events in the TON blockchain for the specified address.
//
// POST /tonconnect/subscribe
func (UnimplementedHandler) SubscribeToAccountEvents(ctx context.Context, req *SubscribeToAccountEventsReq) error {
	return ht.ErrNotImplemented
}

// UnsubscribeFromAccountEvents implements unsubscribeFromAccountEvents operation.
//
// Unsubscribe from notifications about events in the TON blockchain for the specified address.
//
// POST /tonconnect/unsubscribe
func (UnimplementedHandler) UnsubscribeFromAccountEvents(ctx context.Context, req *UnsubscribeFromAccountEventsReq) error {
	return ht.ErrNotImplemented
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrorStatusCode) {
	r = new(ErrorStatusCode)
	return r
}
