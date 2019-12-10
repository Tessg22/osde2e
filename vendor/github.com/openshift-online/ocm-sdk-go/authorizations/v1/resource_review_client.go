/*
Copyright (c) 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/authorizations/v1

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/openshift-online/ocm-sdk-go/errors"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// ResourceReviewClient is the client of the 'resource_review' resource.
//
// Manages resource review.
type ResourceReviewClient struct {
	transport http.RoundTripper
	path      string
	metric    string
}

// NewResourceReviewClient creates a new client for the 'resource_review'
// resource using the given transport to sned the requests and receive the
// responses.
func NewResourceReviewClient(transport http.RoundTripper, path string, metric string) *ResourceReviewClient {
	client := new(ResourceReviewClient)
	client.transport = transport
	client.path = path
	client.metric = metric
	return client
}

// Post creates a request for the 'post' method.
//
// Returns the list of identifiers of the resources that an account can
// perform the specified action upon.
func (c *ResourceReviewClient) Post() *ResourceReviewPostRequest {
	request := new(ResourceReviewPostRequest)
	request.transport = c.transport
	request.path = c.path
	request.metric = c.metric
	return request
}

// ResourceReviewPostRequest is the request for the 'post' method.
type ResourceReviewPostRequest struct {
	transport http.RoundTripper
	path      string
	metric    string
	query     url.Values
	header    http.Header
	request   *ResourceReviewRequest
}

// Parameter adds a query parameter.
func (r *ResourceReviewPostRequest) Parameter(name string, value interface{}) *ResourceReviewPostRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *ResourceReviewPostRequest) Header(name string, value interface{}) *ResourceReviewPostRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Request sets the value of the 'request' parameter.
//
//
func (r *ResourceReviewPostRequest) Request(value *ResourceReviewRequest) *ResourceReviewPostRequest {
	r.request = value
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *ResourceReviewPostRequest) Send() (result *ResourceReviewPostResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *ResourceReviewPostRequest) SendContext(ctx context.Context) (result *ResourceReviewPostResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.SetHeader(r.header, r.metric)
	buffer := new(bytes.Buffer)
	err = r.marshal(buffer)
	if err != nil {
		return
	}
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: "POST",
		URL:    uri,
		Header: header,
		Body:   ioutil.NopCloser(buffer),
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = new(ResourceReviewPostResponse)
	result.status = response.StatusCode
	result.header = response.Header
	if result.status >= 400 {
		result.err, err = errors.UnmarshalError(response.Body)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	err = result.unmarshal(response.Body)
	if err != nil {
		return
	}
	return
}

// marshall is the method used internally to marshal requests for the
// 'post' method.
func (r *ResourceReviewPostRequest) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.request.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// ResourceReviewPostResponse is the response for the 'post' method.
type ResourceReviewPostResponse struct {
	status int
	header http.Header
	err    *errors.Error
	review *ResourceReview
}

// Status returns the response status code.
func (r *ResourceReviewPostResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *ResourceReviewPostResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *ResourceReviewPostResponse) Error() *errors.Error {
	return r.err
}

// Review returns the value of the 'review' parameter.
//
//
func (r *ResourceReviewPostResponse) Review() *ResourceReview {
	if r == nil {
		return nil
	}
	return r.review
}

// GetReview returns the value of the 'review' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *ResourceReviewPostResponse) GetReview() (value *ResourceReview, ok bool) {
	ok = r != nil && r.review != nil
	if ok {
		value = r.review
	}
	return
}

// unmarshal is the method used internally to unmarshal responses to the
// 'post' method.
func (r *ResourceReviewPostResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(resourceReviewData)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}
	r.review, err = data.unwrap()
	if err != nil {
		return err
	}
	return err
}
