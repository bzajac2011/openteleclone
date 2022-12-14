// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ptrace // import "go.opentelemetry.io/collector/pdata/ptrace"

import (
	"bytes"

	"github.com/gogo/protobuf/jsonpb"

	"go.opentelemetry.io/collector/pdata/internal"
	otlptrace "go.opentelemetry.io/collector/pdata/internal/data/protogen/trace/v1"
	"go.opentelemetry.io/collector/pdata/ptrace/internal/ptracejson"
)

// NewJSONMarshaler returns a model.Marshaler. Marshals to OTLP json bytes.
func NewJSONMarshaler() Marshaler {
	return &jsonMarshaler{delegate: jsonpb.Marshaler{}}
}

type jsonMarshaler struct {
	delegate jsonpb.Marshaler
}

func (e *jsonMarshaler) MarshalTraces(td Traces) ([]byte, error) {
	buf := bytes.Buffer{}
	pb := internal.TracesToProto(internal.Traces(td))
	err := e.delegate.Marshal(&buf, &pb)
	return buf.Bytes(), err
}

// NewJSONUnmarshaler returns a model.Unmarshaler. Unmarshalls from OTLP json bytes.
func NewJSONUnmarshaler() Unmarshaler {
	return &jsonUnmarshaler{}
}

type jsonUnmarshaler struct{}

func (jsonUnmarshaler) UnmarshalTraces(buf []byte) (Traces, error) {
	var td otlptrace.TracesData
	if err := ptracejson.UnmarshalTraceData(buf, &td); err != nil {
		return Traces{}, err
	}
	return Traces(internal.TracesFromProto(td)), nil
}
