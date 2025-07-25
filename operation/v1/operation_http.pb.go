// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.4
// - protoc             v6.30.2
// source: operation/v1/operation.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationOperationAppealReview = "/api.operation.v1.Operation/AppealReview"

type OperationHTTPServer interface {
	AppealReview(context.Context, *AppealReviewRequest) (*AppealReviewReply, error)
}

func RegisterOperationHTTPServer(s *http.Server, srv OperationHTTPServer) {
	r := s.Route("/")
	r.POST("operation/v1/review/appeal", _Operation_AppealReview1_HTTP_Handler(srv))
}

func _Operation_AppealReview1_HTTP_Handler(srv OperationHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AppealReviewRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOperationAppealReview)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AppealReview(ctx, req.(*AppealReviewRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AppealReviewReply)
		return ctx.Result(200, reply)
	}
}

type OperationHTTPClient interface {
	AppealReview(ctx context.Context, req *AppealReviewRequest, opts ...http.CallOption) (rsp *AppealReviewReply, err error)
}

type OperationHTTPClientImpl struct {
	cc *http.Client
}

func NewOperationHTTPClient(client *http.Client) OperationHTTPClient {
	return &OperationHTTPClientImpl{client}
}

func (c *OperationHTTPClientImpl) AppealReview(ctx context.Context, in *AppealReviewRequest, opts ...http.CallOption) (*AppealReviewReply, error) {
	var out AppealReviewReply
	pattern := "operation/v1/review/appeal"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOperationAppealReview))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
