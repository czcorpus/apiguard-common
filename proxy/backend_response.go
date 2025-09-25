// Copyright 2022 Tomas Machalek <tomas.machalek@gmail.com>
// Copyright 2022 Martin Zimandl <martin.zimandl@gmail.com>
// Copyright 2022 Department of Linguistics,
//                Faculty of Arts, Charles University
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package proxy

import (
	"fmt"
	"io"
	"net/http"
)


type BackendResponse interface {
	GetBodyReader() io.ReadCloser
	CloseBodyReader() error
	GetHeaders() http.Header
	GetStatusCode() int
	IsDataStream() bool
	Error() error
}

// ----------------------

type EmptyReadCloser struct{}

func (rc EmptyReadCloser) Read(p []byte) (int, error) {
	return 0, io.EOF
}

func (rc EmptyReadCloser) Close() error {
	return nil
}

// ----------------------

type BackendZeroResponse struct {
}

func (sr *BackendZeroResponse) GetBodyReader() io.ReadCloser {
	return &EmptyReadCloser{}
}

func (sr *BackendZeroResponse) CloseBodyReader() error {
	return nil
}

func (sr *BackendZeroResponse) GetHeaders() http.Header {
	return map[string][]string{}
}

func (sr *BackendZeroResponse) GetStatusCode() int {
	return 0
}

func (sr *BackendZeroResponse) Error() error {
	return fmt.Errorf("the response is undefined")
}

func (sr *BackendZeroResponse) IsDataStream() bool {
	return false
}

// -----------------------------------------

// BackendSimpleResponse represents a backend response where we don't
// care about authentication and/or information returned via
// headers
type BackendSimpleResponse struct {
	BodyReader io.ReadCloser
	StatusCode int
	Err        error
}

func (sr *BackendSimpleResponse) GetBodyReader() io.ReadCloser {
	return sr.BodyReader
}

func (sr *BackendSimpleResponse) CloseBodyReader() error {
	return sr.BodyReader.Close()
}

func (sr *BackendSimpleResponse) GetHeaders() http.Header {
	return map[string][]string{}
}

func (sr *BackendSimpleResponse) GetStatusCode() int {
	return sr.StatusCode
}

func (sr *BackendSimpleResponse) Error() error {
	return sr.Err
}

func (sr *BackendSimpleResponse) IsDataStream() bool {
	return false
}
