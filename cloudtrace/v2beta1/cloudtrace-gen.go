// Copyright 2020 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package cloudtrace provides access to the Cloud Trace API.
//
// For product documentation, see: https://cloud.google.com/trace
//
// Creating a client
//
// Usage example:
//
//   import "google.golang.org/api/cloudtrace/v2beta1"
//   ...
//   ctx := context.Background()
//   cloudtraceService, err := cloudtrace.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for authentication.
//
// For information on how to create and obtain Application Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// Other authentication options
//
// By default, all available scopes (see "Constants") are used to authenticate. To restrict scopes, use option.WithScopes:
//
//   cloudtraceService, err := cloudtrace.NewService(ctx, option.WithScopes(cloudtrace.TraceReadonlyScope))
//
// To use an API key for authentication (note: some APIs do not support API keys), use option.WithAPIKey:
//
//   cloudtraceService, err := cloudtrace.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth flow), use option.WithTokenSource:
//
//   config := &oauth2.Config{...}
//   // ...
//   token, err := config.Exchange(ctx, ...)
//   cloudtraceService, err := cloudtrace.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See https://godoc.org/google.golang.org/api/option/ for details on options.
package cloudtrace // import "google.golang.org/api/cloudtrace/v2beta1"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	googleapi "google.golang.org/api/googleapi"
	gensupport "google.golang.org/api/internal/gensupport"
	option "google.golang.org/api/option"
	internaloption "google.golang.org/api/option/internaloption"
	htransport "google.golang.org/api/transport/http"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = internaloption.WithDefaultEndpoint

const apiId = "cloudtrace:v2beta1"
const apiName = "cloudtrace"
const apiVersion = "v2beta1"
const basePath = "https://cloudtrace.googleapis.com/"
const mtlsBasePath = "https://cloudtrace.mtls.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// Write Trace data for a project or application
	TraceAppendScope = "https://www.googleapis.com/auth/trace.append"

	// Read Trace data for a project or application
	TraceReadonlyScope = "https://www.googleapis.com/auth/trace.readonly"
)

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	scopesOption := option.WithScopes(
		"https://www.googleapis.com/auth/cloud-platform",
		"https://www.googleapis.com/auth/trace.append",
		"https://www.googleapis.com/auth/trace.readonly",
	)
	// NOTE: prepend, so we don't override user-specified scopes.
	opts = append([]option.ClientOption{scopesOption}, opts...)
	opts = append(opts, internaloption.WithDefaultEndpoint(basePath))
	opts = append(opts, internaloption.WithDefaultMTLSEndpoint(mtlsBasePath))
	client, endpoint, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	s, err := New(client)
	if err != nil {
		return nil, err
	}
	if endpoint != "" {
		s.BasePath = endpoint
	}
	return s, nil
}

// New creates a new Service. It uses the provided http.Client for requests.
//
// Deprecated: please use NewService instead.
// To provide a custom HTTP client, use option.WithHTTPClient.
// If you are using google.golang.org/api/googleapis/transport.APIKey, use option.WithAPIKey with NewService instead.
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Projects *ProjectsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.TraceSinks = NewProjectsTraceSinksService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	TraceSinks *ProjectsTraceSinksService
}

func NewProjectsTraceSinksService(s *Service) *ProjectsTraceSinksService {
	rs := &ProjectsTraceSinksService{s: s}
	return rs
}

type ProjectsTraceSinksService struct {
	s *Service
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated
// empty messages in your APIs. A typical example is to use it as the
// request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// ListTraceSinksResponse: Result returned from `ListTraceSinks`.
type ListTraceSinksResponse struct {
	// NextPageToken: If there might be more results than appear in this
	// response, then
	// `nextPageToken` is included.  To get the next set of results, call
	// the same
	// method again using the value of `nextPageToken` as `pageToken`.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Sinks: A list of sinks.
	Sinks []*TraceSink `json:"sinks,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListTraceSinksResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListTraceSinksResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// OutputConfig: OutputConfig contains a destination for writing trace
// data.
type OutputConfig struct {
	// Destination: The destination for writing trace data. Currently only
	// BigQuery is
	// supported. E.g.:
	//
	//
	// "bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]"
	Destination string `json:"destination,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Destination") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Destination") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *OutputConfig) MarshalJSON() ([]byte, error) {
	type NoMethod OutputConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TraceSink: Describes a sink used to export traces to a BigQuery
// dataset. The sink must
// be created within a project.
type TraceSink struct {
	// Name: Required. The canonical sink resource name, unique within the
	// project. Must be of the
	// form:
	//
	//   project/[PROJECT_NUMBER]/traceSinks/[SINK_ID].
	//
	// E.g.:
	//
	//   "projects/12345/traceSinks/my-project-trace-sink".
	//
	// Sink identifiers are limited to 256 characters and can include only
	// the
	// following characters: upper and lower-case alphanumeric
	// characters,
	// underscores, hyphens, and periods.
	Name string `json:"name,omitempty"`

	// OutputConfig: Required. The export destination.
	OutputConfig *OutputConfig `json:"outputConfig,omitempty"`

	// WriterIdentity: Output only. A service account name for exporting the
	// data. This field is
	// set by sinks.create and sinks.update. The service account will need
	// to be
	// granted write access to the destination specified in the
	// output
	// configuration, see [Granting access for
	// a
	// resource](/iam/docs/granting-roles-to-service-accounts#granting_acce
	// ss_to_a_service_account_for_a_resource).
	// To create tables and write data this account will need the dataEditor
	// role.
	// Read more about roles in the
	// [BigQuery
	// documentation](https://cloud.google.com/bigquery/docs/access
	// -control).
	// E.g.:
	//
	//    "service-00000001@00000002.iam.gserviceaccount.com"
	WriterIdentity string `json:"writerIdentity,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Name") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Name") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TraceSink) MarshalJSON() ([]byte, error) {
	type NoMethod TraceSink
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "cloudtrace.projects.traceSinks.create":

type ProjectsTraceSinksCreateCall struct {
	s          *Service
	parent     string
	tracesink  *TraceSink
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Creates a sink that exports trace spans to a destination.
// The
// export of newly-ingested traces begins immediately, unless the
// sink's
// `writer_identity` is not permitted to write to the destination.  A
// sink can
// export traces only from the resource owning the sink (the 'parent').
func (r *ProjectsTraceSinksService) Create(parent string, tracesink *TraceSink) *ProjectsTraceSinksCreateCall {
	c := &ProjectsTraceSinksCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.tracesink = tracesink
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTraceSinksCreateCall) Fields(s ...googleapi.Field) *ProjectsTraceSinksCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTraceSinksCreateCall) Context(ctx context.Context) *ProjectsTraceSinksCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsTraceSinksCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsTraceSinksCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20200725")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.tracesink)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/traceSinks")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtrace.projects.traceSinks.create" call.
// Exactly one of *TraceSink or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *TraceSink.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsTraceSinksCreateCall) Do(opts ...googleapi.CallOption) (*TraceSink, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &TraceSink{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a sink that exports trace spans to a destination.  The\nexport of newly-ingested traces begins immediately, unless the sink's\n`writer_identity` is not permitted to write to the destination.  A sink can\nexport traces only from the resource owning the sink (the 'parent').",
	//   "flatPath": "v2beta1/projects/{projectsId}/traceSinks",
	//   "httpMethod": "POST",
	//   "id": "cloudtrace.projects.traceSinks.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The resource in which to create the sink (currently only project\nsinks are supported):\n\n    \"projects/[PROJECT_ID]\"\n\nExamples: `\"projects/my-trace-project\"`, `\"projects/123456789\"`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/traceSinks",
	//   "request": {
	//     "$ref": "TraceSink"
	//   },
	//   "response": {
	//     "$ref": "TraceSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/trace.append"
	//   ]
	// }

}

// method id "cloudtrace.projects.traceSinks.delete":

type ProjectsTraceSinksDeleteCall struct {
	s          *Service
	nameid     string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes a sink.
func (r *ProjectsTraceSinksService) Delete(nameid string) *ProjectsTraceSinksDeleteCall {
	c := &ProjectsTraceSinksDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.nameid = nameid
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTraceSinksDeleteCall) Fields(s ...googleapi.Field) *ProjectsTraceSinksDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTraceSinksDeleteCall) Context(ctx context.Context) *ProjectsTraceSinksDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsTraceSinksDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsTraceSinksDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20200725")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("DELETE", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.nameid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtrace.projects.traceSinks.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsTraceSinksDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a sink.",
	//   "flatPath": "v2beta1/projects/{projectsId}/traceSinks/{traceSinksId}",
	//   "httpMethod": "DELETE",
	//   "id": "cloudtrace.projects.traceSinks.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The full resource name of the sink to delete, including the\nparent resource and the sink identifier:\n\n    \"projects/[PROJECT_NUMBER]/traceSinks/[SINK_ID]\"\n\nExample: `\"projects/12345/traceSinks/my-sink-id\"`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/traceSinks/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/trace.append"
	//   ]
	// }

}

// method id "cloudtrace.projects.traceSinks.get":

type ProjectsTraceSinksGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Get a trace sink by name under the parent resource (GCP
// project).
func (r *ProjectsTraceSinksService) Get(name string) *ProjectsTraceSinksGetCall {
	c := &ProjectsTraceSinksGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTraceSinksGetCall) Fields(s ...googleapi.Field) *ProjectsTraceSinksGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsTraceSinksGetCall) IfNoneMatch(entityTag string) *ProjectsTraceSinksGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTraceSinksGetCall) Context(ctx context.Context) *ProjectsTraceSinksGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsTraceSinksGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsTraceSinksGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20200725")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtrace.projects.traceSinks.get" call.
// Exactly one of *TraceSink or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *TraceSink.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsTraceSinksGetCall) Do(opts ...googleapi.CallOption) (*TraceSink, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &TraceSink{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get a trace sink by name under the parent resource (GCP project).",
	//   "flatPath": "v2beta1/projects/{projectsId}/traceSinks/{traceSinksId}",
	//   "httpMethod": "GET",
	//   "id": "cloudtrace.projects.traceSinks.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The resource name of the sink:\n\n    \"projects/[PROJECT_NUMBER]/traceSinks/[SINK_ID]\"\n\nExample: `\"projects/12345/traceSinks/my-sink-id\"`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/traceSinks/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "TraceSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/trace.readonly"
	//   ]
	// }

}

// method id "cloudtrace.projects.traceSinks.list":

type ProjectsTraceSinksListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: List all sinks for the parent resource (GCP project).
func (r *ProjectsTraceSinksService) List(parent string) *ProjectsTraceSinksListCall {
	c := &ProjectsTraceSinksListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return from this request.
// Non-positive values are ignored.  The presence of `nextPageToken` in
// the
// response indicates that more results might be available.
func (c *ProjectsTraceSinksListCall) PageSize(pageSize int64) *ProjectsTraceSinksListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": If present, then
// retrieve the next batch of results from the
// preceding call to this method.  `pageToken` must be the value
// of
// `nextPageToken` from the previous response.  The values of other
// method
// parameters should be identical to those in the previous call.
func (c *ProjectsTraceSinksListCall) PageToken(pageToken string) *ProjectsTraceSinksListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTraceSinksListCall) Fields(s ...googleapi.Field) *ProjectsTraceSinksListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsTraceSinksListCall) IfNoneMatch(entityTag string) *ProjectsTraceSinksListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTraceSinksListCall) Context(ctx context.Context) *ProjectsTraceSinksListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsTraceSinksListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsTraceSinksListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20200725")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/traceSinks")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtrace.projects.traceSinks.list" call.
// Exactly one of *ListTraceSinksResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListTraceSinksResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsTraceSinksListCall) Do(opts ...googleapi.CallOption) (*ListTraceSinksResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListTraceSinksResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all sinks for the parent resource (GCP project).",
	//   "flatPath": "v2beta1/projects/{projectsId}/traceSinks",
	//   "httpMethod": "GET",
	//   "id": "cloudtrace.projects.traceSinks.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The maximum number of results to return from this request.\nNon-positive values are ignored.  The presence of `nextPageToken` in the\nresponse indicates that more results might be available.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. If present, then retrieve the next batch of results from the\npreceding call to this method.  `pageToken` must be the value of\n`nextPageToken` from the previous response.  The values of other method\nparameters should be identical to those in the previous call.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The parent resource whose sinks are to be listed (currently only\nproject parent resources are supported):\n\n    \"projects/[PROJECT_ID]\"",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/traceSinks",
	//   "response": {
	//     "$ref": "ListTraceSinksResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/trace.readonly"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsTraceSinksListCall) Pages(ctx context.Context, f func(*ListTraceSinksResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "cloudtrace.projects.traceSinks.patch":

type ProjectsTraceSinksPatchCall struct {
	s          *Service
	nameid     string
	tracesink  *TraceSink
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Patch: Updates a sink.  This method updates fields in the existing
// sink according
// to the provided update mask. The sink's name cannot be changed nor
// any
// output-only fields (e.g. the writer_identity).
func (r *ProjectsTraceSinksService) Patch(nameid string, tracesink *TraceSink) *ProjectsTraceSinksPatchCall {
	c := &ProjectsTraceSinksPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.nameid = nameid
	c.tracesink = tracesink
	return c
}

// UpdateMask sets the optional parameter "updateMask": Required. Field
// mask that specifies the fields in `trace_sink` that are to
// be updated. A sink field is overwritten if, and only if, it is
// in the update mask.  `name` and `writer_identity` fields cannot be
// updated.
//
// An empty updateMask is considered an error.
//
// For a detailed `FieldMask` definition,
// see
// https://developers.google.com/protocol-buffers/docs/reference/goog
// le.protobuf#fieldmask
//
// Example: `updateMask=output_config`.
func (c *ProjectsTraceSinksPatchCall) UpdateMask(updateMask string) *ProjectsTraceSinksPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTraceSinksPatchCall) Fields(s ...googleapi.Field) *ProjectsTraceSinksPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTraceSinksPatchCall) Context(ctx context.Context) *ProjectsTraceSinksPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsTraceSinksPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsTraceSinksPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20200725")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.tracesink)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("PATCH", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.nameid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtrace.projects.traceSinks.patch" call.
// Exactly one of *TraceSink or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *TraceSink.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsTraceSinksPatchCall) Do(opts ...googleapi.CallOption) (*TraceSink, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &TraceSink{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a sink.  This method updates fields in the existing sink according\nto the provided update mask. The sink's name cannot be changed nor any\noutput-only fields (e.g. the writer_identity).",
	//   "flatPath": "v2beta1/projects/{projectsId}/traceSinks/{traceSinksId}",
	//   "httpMethod": "PATCH",
	//   "id": "cloudtrace.projects.traceSinks.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The full resource name of the sink to update, including the\nparent resource and the sink identifier:\n\n    \"projects/[PROJECT_NUMBER]/traceSinks/[SINK_ID]\"\n\nExample: `\"projects/12345/traceSinks/my-sink-id\"`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/traceSinks/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Required. Field mask that specifies the fields in `trace_sink` that are to\nbe updated. A sink field is overwritten if, and only if, it is\nin the update mask.  `name` and `writer_identity` fields cannot be updated.\n\nAn empty updateMask is considered an error.\n\nFor a detailed `FieldMask` definition, see\nhttps://developers.google.com/protocol-buffers/docs/reference/google.protobuf#fieldmask\n\nExample: `updateMask=output_config`.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "request": {
	//     "$ref": "TraceSink"
	//   },
	//   "response": {
	//     "$ref": "TraceSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/trace.append"
	//   ]
	// }

}
