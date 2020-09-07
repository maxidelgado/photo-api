package ctxhelper

import "context"

const Key = "RequestMetadata"

type ContextHelper interface {
	GetRequestId() string
	SetRequestId(string)
}

type RequestHelper struct {
	requestId string
}

func New() ContextHelper {
	rh := RequestHelper{}

	return &rh
}

func WithContext(ctx context.Context) ContextHelper {
	if rh, ok := ctx.Value(Key).(ContextHelper); ok {
		return rh
	}
	return New()
}

func (r *RequestHelper) SetRequestId(rid string) {
	r.requestId = rid
}

func (r *RequestHelper) GetRequestId() string {
	return r.requestId
}
