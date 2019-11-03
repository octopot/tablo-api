// Code generated by protoc-gen-twirp v5.8.0, DO NOT EDIT.
// source: v1/vendor.proto

package v1

import bytes "bytes"
import strings "strings"
import context "context"
import fmt "fmt"
import ioutil "io/ioutil"
import http "net/http"
import strconv "strconv"

import jsonpb "github.com/golang/protobuf/jsonpb"
import proto "github.com/golang/protobuf/proto"
import twirp "github.com/twitchtv/twirp"
import ctxsetters "github.com/twitchtv/twirp/ctxsetters"

import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

// =======================
// VendorService Interface
// =======================

type VendorService interface {
	Connect(context.Context, *Secret) (*Vendor, error)

	Connected(context.Context, *google_protobuf1.Empty) (*VendorList, error)

	Disconnect(context.Context, *URI) (*Void, error)

	Sources(context.Context, *URI) (*SourceNode, error)
}

// =============================
// VendorService Protobuf Client
// =============================

type vendorServiceProtobufClient struct {
	client HTTPClient
	urls   [4]string
}

// NewVendorServiceProtobufClient creates a Protobuf client that implements the VendorService interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewVendorServiceProtobufClient(addr string, client HTTPClient) VendorService {
	prefix := urlBase(addr) + VendorServicePathPrefix
	urls := [4]string{
		prefix + "Connect",
		prefix + "Connected",
		prefix + "Disconnect",
		prefix + "Sources",
	}
	if httpClient, ok := client.(*http.Client); ok {
		return &vendorServiceProtobufClient{
			client: withoutRedirects(httpClient),
			urls:   urls,
		}
	}
	return &vendorServiceProtobufClient{
		client: client,
		urls:   urls,
	}
}

func (c *vendorServiceProtobufClient) Connect(ctx context.Context, in *Secret) (*Vendor, error) {
	ctx = ctxsetters.WithPackageName(ctx, "octolab.api.tablo.v1")
	ctx = ctxsetters.WithServiceName(ctx, "VendorService")
	ctx = ctxsetters.WithMethodName(ctx, "Connect")
	out := new(Vendor)
	err := doProtobufRequest(ctx, c.client, c.urls[0], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorServiceProtobufClient) Connected(ctx context.Context, in *google_protobuf1.Empty) (*VendorList, error) {
	ctx = ctxsetters.WithPackageName(ctx, "octolab.api.tablo.v1")
	ctx = ctxsetters.WithServiceName(ctx, "VendorService")
	ctx = ctxsetters.WithMethodName(ctx, "Connected")
	out := new(VendorList)
	err := doProtobufRequest(ctx, c.client, c.urls[1], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorServiceProtobufClient) Disconnect(ctx context.Context, in *URI) (*Void, error) {
	ctx = ctxsetters.WithPackageName(ctx, "octolab.api.tablo.v1")
	ctx = ctxsetters.WithServiceName(ctx, "VendorService")
	ctx = ctxsetters.WithMethodName(ctx, "Disconnect")
	out := new(Void)
	err := doProtobufRequest(ctx, c.client, c.urls[2], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorServiceProtobufClient) Sources(ctx context.Context, in *URI) (*SourceNode, error) {
	ctx = ctxsetters.WithPackageName(ctx, "octolab.api.tablo.v1")
	ctx = ctxsetters.WithServiceName(ctx, "VendorService")
	ctx = ctxsetters.WithMethodName(ctx, "Sources")
	out := new(SourceNode)
	err := doProtobufRequest(ctx, c.client, c.urls[3], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// =========================
// VendorService JSON Client
// =========================

type vendorServiceJSONClient struct {
	client HTTPClient
	urls   [4]string
}

// NewVendorServiceJSONClient creates a JSON client that implements the VendorService interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewVendorServiceJSONClient(addr string, client HTTPClient) VendorService {
	prefix := urlBase(addr) + VendorServicePathPrefix
	urls := [4]string{
		prefix + "Connect",
		prefix + "Connected",
		prefix + "Disconnect",
		prefix + "Sources",
	}
	if httpClient, ok := client.(*http.Client); ok {
		return &vendorServiceJSONClient{
			client: withoutRedirects(httpClient),
			urls:   urls,
		}
	}
	return &vendorServiceJSONClient{
		client: client,
		urls:   urls,
	}
}

func (c *vendorServiceJSONClient) Connect(ctx context.Context, in *Secret) (*Vendor, error) {
	ctx = ctxsetters.WithPackageName(ctx, "octolab.api.tablo.v1")
	ctx = ctxsetters.WithServiceName(ctx, "VendorService")
	ctx = ctxsetters.WithMethodName(ctx, "Connect")
	out := new(Vendor)
	err := doJSONRequest(ctx, c.client, c.urls[0], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorServiceJSONClient) Connected(ctx context.Context, in *google_protobuf1.Empty) (*VendorList, error) {
	ctx = ctxsetters.WithPackageName(ctx, "octolab.api.tablo.v1")
	ctx = ctxsetters.WithServiceName(ctx, "VendorService")
	ctx = ctxsetters.WithMethodName(ctx, "Connected")
	out := new(VendorList)
	err := doJSONRequest(ctx, c.client, c.urls[1], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorServiceJSONClient) Disconnect(ctx context.Context, in *URI) (*Void, error) {
	ctx = ctxsetters.WithPackageName(ctx, "octolab.api.tablo.v1")
	ctx = ctxsetters.WithServiceName(ctx, "VendorService")
	ctx = ctxsetters.WithMethodName(ctx, "Disconnect")
	out := new(Void)
	err := doJSONRequest(ctx, c.client, c.urls[2], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorServiceJSONClient) Sources(ctx context.Context, in *URI) (*SourceNode, error) {
	ctx = ctxsetters.WithPackageName(ctx, "octolab.api.tablo.v1")
	ctx = ctxsetters.WithServiceName(ctx, "VendorService")
	ctx = ctxsetters.WithMethodName(ctx, "Sources")
	out := new(SourceNode)
	err := doJSONRequest(ctx, c.client, c.urls[3], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ============================
// VendorService Server Handler
// ============================

type vendorServiceServer struct {
	VendorService
	hooks *twirp.ServerHooks
}

func NewVendorServiceServer(svc VendorService, hooks *twirp.ServerHooks) TwirpServer {
	return &vendorServiceServer{
		VendorService: svc,
		hooks:         hooks,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *vendorServiceServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// VendorServicePathPrefix is used for all URL paths on a twirp VendorService server.
// Requests are always: POST VendorServicePathPrefix/method
// It can be used in an HTTP mux to route twirp requests along with non-twirp requests on other routes.
const VendorServicePathPrefix = "/twirp/octolab.api.tablo.v1.VendorService/"

func (s *vendorServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "octolab.api.tablo.v1")
	ctx = ctxsetters.WithServiceName(ctx, "VendorService")
	ctx = ctxsetters.WithResponseWriter(ctx, resp)

	var err error
	ctx, err = callRequestReceived(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	if req.Method != "POST" {
		msg := fmt.Sprintf("unsupported method %q (only POST is allowed)", req.Method)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}

	switch req.URL.Path {
	case "/twirp/octolab.api.tablo.v1.VendorService/Connect":
		s.serveConnect(ctx, resp, req)
		return
	case "/twirp/octolab.api.tablo.v1.VendorService/Connected":
		s.serveConnected(ctx, resp, req)
		return
	case "/twirp/octolab.api.tablo.v1.VendorService/Disconnect":
		s.serveDisconnect(ctx, resp, req)
		return
	case "/twirp/octolab.api.tablo.v1.VendorService/Sources":
		s.serveSources(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}
}

func (s *vendorServiceServer) serveConnect(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveConnectJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveConnectProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *vendorServiceServer) serveConnectJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Connect")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(Secret)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *Vendor
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.VendorService.Connect(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Vendor and nil error while calling Connect. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes := buf.Bytes()
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *vendorServiceServer) serveConnectProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Connect")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to read request body"))
		return
	}
	reqContent := new(Secret)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *Vendor
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.VendorService.Connect(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Vendor and nil error while calling Connect. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *vendorServiceServer) serveConnected(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveConnectedJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveConnectedProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *vendorServiceServer) serveConnectedJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Connected")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(google_protobuf1.Empty)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *VendorList
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.VendorService.Connected(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *VendorList and nil error while calling Connected. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes := buf.Bytes()
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *vendorServiceServer) serveConnectedProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Connected")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to read request body"))
		return
	}
	reqContent := new(google_protobuf1.Empty)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *VendorList
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.VendorService.Connected(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *VendorList and nil error while calling Connected. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *vendorServiceServer) serveDisconnect(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveDisconnectJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveDisconnectProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *vendorServiceServer) serveDisconnectJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Disconnect")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(URI)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *Void
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.VendorService.Disconnect(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Void and nil error while calling Disconnect. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes := buf.Bytes()
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *vendorServiceServer) serveDisconnectProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Disconnect")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to read request body"))
		return
	}
	reqContent := new(URI)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *Void
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.VendorService.Disconnect(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Void and nil error while calling Disconnect. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *vendorServiceServer) serveSources(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveSourcesJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveSourcesProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *vendorServiceServer) serveSourcesJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Sources")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(URI)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *SourceNode
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.VendorService.Sources(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *SourceNode and nil error while calling Sources. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes := buf.Bytes()
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *vendorServiceServer) serveSourcesProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Sources")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to read request body"))
		return
	}
	reqContent := new(URI)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *SourceNode
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.VendorService.Sources(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *SourceNode and nil error while calling Sources. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *vendorServiceServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor1, 0
}

func (s *vendorServiceServer) ProtocGenTwirpVersion() string {
	return "v5.8.0"
}

func (s *vendorServiceServer) PathPrefix() string {
	return VendorServicePathPrefix
}

var twirpFileDescriptor1 = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x8d, 0x5d, 0xd7, 0x6d, 0x6f, 0xf4, 0x7d, 0x54, 0xa3, 0x0a, 0x05, 0x83, 0x90, 0x15, 0x09,
	0x29, 0xb0, 0xb0, 0x89, 0xd9, 0x40, 0x0b, 0x2c, 0x52, 0x02, 0xa9, 0x14, 0x7e, 0xe4, 0xb4, 0x2c,
	0xd8, 0x4d, 0x3c, 0x17, 0x33, 0xc2, 0xf6, 0x58, 0xe3, 0xa9, 0xa5, 0xac, 0x79, 0x01, 0x5e, 0x80,
	0x17, 0xe0, 0x29, 0x91, 0x67, 0x9c, 0x08, 0x89, 0xfc, 0x54, 0x48, 0xec, 0x26, 0x37, 0xe7, 0x1c,
	0x9f, 0x73, 0x7f, 0xe0, 0x56, 0x3d, 0x0c, 0x6b, 0x2c, 0x98, 0x90, 0x41, 0x29, 0x85, 0x12, 0xe4,
	0x44, 0x24, 0x4a, 0x64, 0x74, 0x1e, 0xd0, 0x92, 0x07, 0x8a, 0xce, 0x33, 0x11, 0xd4, 0x43, 0xef,
	0x6e, 0x2a, 0x44, 0x9a, 0x61, 0xa8, 0x31, 0xf3, 0xeb, 0xcf, 0x21, 0xe6, 0xa5, 0x5a, 0x18, 0x8a,
	0xd7, 0x68, 0x24, 0x22, 0xcf, 0x45, 0x61, 0x0a, 0xfd, 0xef, 0x16, 0xb8, 0x33, 0x4c, 0x24, 0x2a,
	0x72, 0x02, 0xfb, 0x4a, 0x7c, 0xc5, 0xa2, 0x67, 0xf9, 0xd6, 0xe0, 0x28, 0x36, 0x3f, 0x08, 0x01,
	0xa7, 0xa0, 0x39, 0xf6, 0x6c, 0x5d, 0xd4, 0x6f, 0xe2, 0x43, 0x97, 0x61, 0x95, 0x48, 0x5e, 0x2a,
	0x2e, 0x8a, 0xde, 0x9e, 0xfe, 0xeb, 0xf7, 0x12, 0x39, 0x85, 0xc3, 0x52, 0x8a, 0x9a, 0x33, 0x94,
	0x3d, 0xc7, 0xb7, 0x06, 0xff, 0x47, 0xf7, 0x83, 0x75, 0x6e, 0x83, 0x0f, 0x2d, 0x2a, 0x5e, 0xe1,
	0xfb, 0x3f, 0x2c, 0x70, 0x3f, 0xea, 0x9c, 0xe4, 0x21, 0xd8, 0x9c, 0x69, 0x3f, 0xdd, 0xe8, 0xce,
	0x7a, 0x81, 0xab, 0xf8, 0x22, 0xb6, 0x39, 0xfb, 0x4b, 0x9f, 0x11, 0x38, 0x6a, 0x51, 0xe2, 0x0d,
	0x3d, 0x6a, 0x6c, 0xff, 0x25, 0x80, 0xb1, 0x37, 0xe5, 0x95, 0x22, 0x8f, 0xc1, 0xc9, 0x78, 0xa5,
	0x7a, 0x96, 0xbf, 0x37, 0xe8, 0x46, 0xf7, 0xd6, 0x2b, 0x18, 0x7c, 0xac, 0x91, 0x7d, 0x0e, 0xee,
	0x4c, 0x5c, 0xcb, 0x04, 0xff, 0x79, 0xbc, 0xc6, 0xaa, 0xf9, 0xd4, 0xcd, 0xad, 0x1a, 0x7c, 0x6b,
	0xf5, 0x9b, 0xbd, 0x14, 0x78, 0x27, 0x18, 0x92, 0x67, 0xe0, 0xe4, 0xa8, 0x68, 0xeb, 0xf8, 0xc1,
	0x36, 0x81, 0x06, 0x1f, 0xbc, 0x45, 0x45, 0x63, 0x4d, 0x21, 0x4f, 0x61, 0x3f, 0xf9, 0xc2, 0x33,
	0xa6, 0x03, 0x74, 0x23, 0x7f, 0x17, 0x77, 0xd2, 0x89, 0x0d, 0x81, 0x9c, 0x82, 0x9b, 0x21, 0xad,
	0xb1, 0xd2, 0x01, 0x77, 0x50, 0x9b, 0x9c, 0x93, 0x4e, 0xdc, 0x32, 0xbc, 0xe7, 0xe0, 0x34, 0x1e,
	0x56, 0xdd, 0xb3, 0x36, 0x77, 0xcf, 0xfe, 0xa3, 0x7b, 0x23, 0x17, 0x1c, 0x46, 0x15, 0x7d, 0xd4,
	0x87, 0xc3, 0xe5, 0x0a, 0x10, 0x00, 0xf7, 0xcd, 0xc5, 0xe5, 0xe4, 0x6a, 0x74, 0xdc, 0x69, 0xde,
	0x97, 0xf1, 0x78, 0x3a, 0x7d, 0x7f, 0x6c, 0x45, 0x3f, 0x6d, 0xf8, 0xcf, 0x4c, 0x79, 0x86, 0xb2,
	0xe6, 0x09, 0x92, 0x31, 0x1c, 0x9c, 0x8b, 0xa2, 0xc0, 0x44, 0x91, 0x4d, 0xad, 0xd6, 0x77, 0xe7,
	0x6d, 0xdd, 0x19, 0x32, 0x86, 0xa3, 0x56, 0x06, 0x19, 0xb9, 0x1d, 0x98, 0xe3, 0x0e, 0x96, 0xc7,
	0x1d, 0x8c, 0x9b, 0xe3, 0xf6, 0xfc, 0x6d, 0x12, 0x7a, 0xf6, 0xe7, 0x00, 0xaf, 0x78, 0x95, 0xb4,
	0x86, 0x36, 0x2f, 0x9b, 0xe7, 0x6d, 0x90, 0x12, 0x9c, 0x91, 0xd7, 0x70, 0x60, 0xda, 0x5c, 0x6d,
	0x53, 0xd8, 0x39, 0xdb, 0xd1, 0x8b, 0x4f, 0x67, 0xa9, 0x58, 0xa1, 0x84, 0x4c, 0x43, 0x4c, 0x44,
	0xb5, 0xa8, 0x14, 0xe6, 0xa1, 0xc6, 0x87, 0xbc, 0x50, 0x28, 0x0b, 0x9a, 0x85, 0x29, 0x16, 0x28,
	0xa9, 0x42, 0x16, 0xd2, 0x92, 0x87, 0xf5, 0xf0, 0xac, 0x1e, 0xce, 0x5d, 0x9d, 0xfe, 0xc9, 0xaf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x50, 0xb7, 0x19, 0xa7, 0x11, 0x05, 0x00, 0x00,
}
