package analytics

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dgraph-io/dgo/v210/protos/api"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
)

func NewSpan(name string) (opentracing.Span, error) {
	tracer := opentracing.GlobalTracer()
	return tracer.StartSpan(name)
}

func NewSpanWithParent(parent opentracing.Span, name string) (opentracing.Span, error) {
	tracer := opentracing.GlobalTracer()
	return tracer.StartSpan(name, opentracing.ChildOf(parent.Context()))
}

func NewSpanWithTag(parent opentracing.Span, name string, tags map[string]interface{}) (opentracing.Span, error) {
	tracer := opentracing.GlobalTracer()
	return tracer.StartSpan(name, opentracing.ChildOf(parent.Context()), opentracing.Tags(tags))
}

func CloseSpan(span opentracing.Span) {
	if err := span.Finish(); err!= nil {
		log.Printf("Error closing span: %s", err)
	}
}

func GetRandomUUID() (string, error) {
	u, err := uuid.NewRandom()
	if err!= nil {
		return "", err
	}
	return u.String(), nil
}

func GetUnixTimestamp() (int64, error) {
	return time.Now().UnixNano()/int64(time.Millisecond), nil
}

func NewDgraphClient() (*api.DgraphClient, error) {
	return api.NewDgraphClient("localhost:9080")
}

func NewDgraphClientWithTimeout(ctx context.Context, timeout time.Duration) (*api.DgraphClient, error) {
	return api.NewDgraphClientWithTimeout("localhost:9080", timeout)
}

func DgraphClientRequest(client *api.DgraphClient, req string, variables map[string]interface{}) (*api.DgraphResponse, error) {
	var resp *api.DgraphResponse
	err := client.NewBatch().Do(context.Background(), req, variables, &resp)
	if err!= nil {
		return nil, fmt.Errorf("error making Dgraph request: %s", err)
	}
	return resp, nil
}

func DgraphClientBatchRequest(client *api.DgraphClient, req string, variables map[string]interface{}) (*api.DgraphResponse, error) {
	var resp *api.DgraphResponse
	err := client.NewBatch().Do(context.Background(), req, variables, &resp)
	if err!= nil {
		return nil, fmt.Errorf("error making Dgraph request: %s", err)
	}
	return resp, nil
}