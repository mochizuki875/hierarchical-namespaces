// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package monitoring

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newServiceMonitoringClientHook clientHook

// ServiceMonitoringCallOptions contains the retry settings for each method of ServiceMonitoringClient.
type ServiceMonitoringCallOptions struct {
	CreateService               []gax.CallOption
	GetService                  []gax.CallOption
	ListServices                []gax.CallOption
	UpdateService               []gax.CallOption
	DeleteService               []gax.CallOption
	CreateServiceLevelObjective []gax.CallOption
	GetServiceLevelObjective    []gax.CallOption
	ListServiceLevelObjectives  []gax.CallOption
	UpdateServiceLevelObjective []gax.CallOption
	DeleteServiceLevelObjective []gax.CallOption
}

func defaultServiceMonitoringClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("monitoring.googleapis.com:443"),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultServiceMonitoringCallOptions() *ServiceMonitoringCallOptions {
	return &ServiceMonitoringCallOptions{
		CreateService: []gax.CallOption{},
		GetService: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListServices: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateService: []gax.CallOption{},
		DeleteService: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateServiceLevelObjective: []gax.CallOption{},
		GetServiceLevelObjective: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListServiceLevelObjectives: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateServiceLevelObjective: []gax.CallOption{},
		DeleteServiceLevelObjective: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// ServiceMonitoringClient is a client for interacting with Cloud Monitoring API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type ServiceMonitoringClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// The gRPC API client.
	serviceMonitoringClient monitoringpb.ServiceMonitoringServiceClient

	// The call options for this service.
	CallOptions *ServiceMonitoringCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewServiceMonitoringClient creates a new service monitoring service client.
//
// The Cloud Monitoring Service-Oriented Monitoring API has endpoints for
// managing and querying aspects of a workspace’s services. These include the
// Service's monitored resources, its Service-Level Objectives, and a taxonomy
// of categorized Health Metrics.
func NewServiceMonitoringClient(ctx context.Context, opts ...option.ClientOption) (*ServiceMonitoringClient, error) {
	clientOpts := defaultServiceMonitoringClientOptions()

	if newServiceMonitoringClientHook != nil {
		hookOpts, err := newServiceMonitoringClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	c := &ServiceMonitoringClient{
		connPool:    connPool,
		CallOptions: defaultServiceMonitoringCallOptions(),

		serviceMonitoringClient: monitoringpb.NewServiceMonitoringServiceClient(connPool),
	}
	c.setGoogleClientInfo()

	return c, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ServiceMonitoringClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ServiceMonitoringClient) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ServiceMonitoringClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// CreateService create a Service.
func (c *ServiceMonitoringClient) CreateService(ctx context.Context, req *monitoringpb.CreateServiceRequest, opts ...gax.CallOption) (*monitoringpb.Service, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.CreateService[0:len(c.CallOptions.CreateService):len(c.CallOptions.CreateService)], opts...)
	var resp *monitoringpb.Service
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.serviceMonitoringClient.CreateService(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetService get the named Service.
func (c *ServiceMonitoringClient) GetService(ctx context.Context, req *monitoringpb.GetServiceRequest, opts ...gax.CallOption) (*monitoringpb.Service, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetService[0:len(c.CallOptions.GetService):len(c.CallOptions.GetService)], opts...)
	var resp *monitoringpb.Service
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.serviceMonitoringClient.GetService(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListServices list Services for this workspace.
func (c *ServiceMonitoringClient) ListServices(ctx context.Context, req *monitoringpb.ListServicesRequest, opts ...gax.CallOption) *ServiceIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListServices[0:len(c.CallOptions.ListServices):len(c.CallOptions.ListServices)], opts...)
	it := &ServiceIterator{}
	req = proto.Clone(req).(*monitoringpb.ListServicesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*monitoringpb.Service, string, error) {
		var resp *monitoringpb.ListServicesResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.serviceMonitoringClient.ListServices(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.Services, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	it.pageInfo.Token = req.PageToken
	return it
}

// UpdateService update this Service.
func (c *ServiceMonitoringClient) UpdateService(ctx context.Context, req *monitoringpb.UpdateServiceRequest, opts ...gax.CallOption) (*monitoringpb.Service, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "service.name", url.QueryEscape(req.GetService().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UpdateService[0:len(c.CallOptions.UpdateService):len(c.CallOptions.UpdateService)], opts...)
	var resp *monitoringpb.Service
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.serviceMonitoringClient.UpdateService(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteService soft delete this Service.
func (c *ServiceMonitoringClient) DeleteService(ctx context.Context, req *monitoringpb.DeleteServiceRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.DeleteService[0:len(c.CallOptions.DeleteService):len(c.CallOptions.DeleteService)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.serviceMonitoringClient.DeleteService(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// CreateServiceLevelObjective create a ServiceLevelObjective for the given Service.
func (c *ServiceMonitoringClient) CreateServiceLevelObjective(ctx context.Context, req *monitoringpb.CreateServiceLevelObjectiveRequest, opts ...gax.CallOption) (*monitoringpb.ServiceLevelObjective, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.CreateServiceLevelObjective[0:len(c.CallOptions.CreateServiceLevelObjective):len(c.CallOptions.CreateServiceLevelObjective)], opts...)
	var resp *monitoringpb.ServiceLevelObjective
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.serviceMonitoringClient.CreateServiceLevelObjective(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetServiceLevelObjective get a ServiceLevelObjective by name.
func (c *ServiceMonitoringClient) GetServiceLevelObjective(ctx context.Context, req *monitoringpb.GetServiceLevelObjectiveRequest, opts ...gax.CallOption) (*monitoringpb.ServiceLevelObjective, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetServiceLevelObjective[0:len(c.CallOptions.GetServiceLevelObjective):len(c.CallOptions.GetServiceLevelObjective)], opts...)
	var resp *monitoringpb.ServiceLevelObjective
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.serviceMonitoringClient.GetServiceLevelObjective(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListServiceLevelObjectives list the ServiceLevelObjectives for the given Service.
func (c *ServiceMonitoringClient) ListServiceLevelObjectives(ctx context.Context, req *monitoringpb.ListServiceLevelObjectivesRequest, opts ...gax.CallOption) *ServiceLevelObjectiveIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListServiceLevelObjectives[0:len(c.CallOptions.ListServiceLevelObjectives):len(c.CallOptions.ListServiceLevelObjectives)], opts...)
	it := &ServiceLevelObjectiveIterator{}
	req = proto.Clone(req).(*monitoringpb.ListServiceLevelObjectivesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*monitoringpb.ServiceLevelObjective, string, error) {
		var resp *monitoringpb.ListServiceLevelObjectivesResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.serviceMonitoringClient.ListServiceLevelObjectives(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.ServiceLevelObjectives, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	it.pageInfo.Token = req.PageToken
	return it
}

// UpdateServiceLevelObjective update the given ServiceLevelObjective.
func (c *ServiceMonitoringClient) UpdateServiceLevelObjective(ctx context.Context, req *monitoringpb.UpdateServiceLevelObjectiveRequest, opts ...gax.CallOption) (*monitoringpb.ServiceLevelObjective, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "service_level_objective.name", url.QueryEscape(req.GetServiceLevelObjective().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UpdateServiceLevelObjective[0:len(c.CallOptions.UpdateServiceLevelObjective):len(c.CallOptions.UpdateServiceLevelObjective)], opts...)
	var resp *monitoringpb.ServiceLevelObjective
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.serviceMonitoringClient.UpdateServiceLevelObjective(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteServiceLevelObjective delete the given ServiceLevelObjective.
func (c *ServiceMonitoringClient) DeleteServiceLevelObjective(ctx context.Context, req *monitoringpb.DeleteServiceLevelObjectiveRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.DeleteServiceLevelObjective[0:len(c.CallOptions.DeleteServiceLevelObjective):len(c.CallOptions.DeleteServiceLevelObjective)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.serviceMonitoringClient.DeleteServiceLevelObjective(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// ServiceIterator manages a stream of *monitoringpb.Service.
type ServiceIterator struct {
	items    []*monitoringpb.Service
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*monitoringpb.Service, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ServiceIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ServiceIterator) Next() (*monitoringpb.Service, error) {
	var item *monitoringpb.Service
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ServiceIterator) bufLen() int {
	return len(it.items)
}

func (it *ServiceIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// ServiceLevelObjectiveIterator manages a stream of *monitoringpb.ServiceLevelObjective.
type ServiceLevelObjectiveIterator struct {
	items    []*monitoringpb.ServiceLevelObjective
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*monitoringpb.ServiceLevelObjective, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ServiceLevelObjectiveIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ServiceLevelObjectiveIterator) Next() (*monitoringpb.ServiceLevelObjective, error) {
	var item *monitoringpb.ServiceLevelObjective
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ServiceLevelObjectiveIterator) bufLen() int {
	return len(it.items)
}

func (it *ServiceLevelObjectiveIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
