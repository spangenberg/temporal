// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by thriftrw v1.17.0. DO NOT EDIT.
// @generated

package admin

import (
	bytes "bytes"
	base64 "encoding/base64"
	fmt "fmt"
	shared "github.com/uber/cadence/.gen/go/shared"
	multierr "go.uber.org/multierr"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

type DescribeWorkflowExecutionRequest struct {
	Domain    *string                   `json:"domain,omitempty"`
	Execution *shared.WorkflowExecution `json:"execution,omitempty"`
}

// ToWire translates a DescribeWorkflowExecutionRequest struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *DescribeWorkflowExecutionRequest) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Domain != nil {
		w, err = wire.NewValueString(*(v.Domain)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 10, Value: w}
		i++
	}
	if v.Execution != nil {
		w, err = v.Execution.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 20, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _WorkflowExecution_Read(w wire.Value) (*shared.WorkflowExecution, error) {
	var v shared.WorkflowExecution
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a DescribeWorkflowExecutionRequest struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a DescribeWorkflowExecutionRequest struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v DescribeWorkflowExecutionRequest
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *DescribeWorkflowExecutionRequest) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 10:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Domain = &x
				if err != nil {
					return err
				}

			}
		case 20:
			if field.Value.Type() == wire.TStruct {
				v.Execution, err = _WorkflowExecution_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a DescribeWorkflowExecutionRequest
// struct.
func (v *DescribeWorkflowExecutionRequest) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	if v.Domain != nil {
		fields[i] = fmt.Sprintf("Domain: %v", *(v.Domain))
		i++
	}
	if v.Execution != nil {
		fields[i] = fmt.Sprintf("Execution: %v", v.Execution)
		i++
	}

	return fmt.Sprintf("DescribeWorkflowExecutionRequest{%v}", strings.Join(fields[:i], ", "))
}

func _String_EqualsPtr(lhs, rhs *string) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this DescribeWorkflowExecutionRequest match the
// provided DescribeWorkflowExecutionRequest.
//
// This function performs a deep comparison.
func (v *DescribeWorkflowExecutionRequest) Equals(rhs *DescribeWorkflowExecutionRequest) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_String_EqualsPtr(v.Domain, rhs.Domain) {
		return false
	}
	if !((v.Execution == nil && rhs.Execution == nil) || (v.Execution != nil && rhs.Execution != nil && v.Execution.Equals(rhs.Execution))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of DescribeWorkflowExecutionRequest.
func (v *DescribeWorkflowExecutionRequest) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Domain != nil {
		enc.AddString("domain", *v.Domain)
	}
	if v.Execution != nil {
		err = multierr.Append(err, enc.AddObject("execution", v.Execution))
	}
	return err
}

// GetDomain returns the value of Domain if it is set or its
// zero value if it is unset.
func (v *DescribeWorkflowExecutionRequest) GetDomain() (o string) {
	if v != nil && v.Domain != nil {
		return *v.Domain
	}

	return
}

// IsSetDomain returns true if Domain is not nil.
func (v *DescribeWorkflowExecutionRequest) IsSetDomain() bool {
	return v != nil && v.Domain != nil
}

// GetExecution returns the value of Execution if it is set or its
// zero value if it is unset.
func (v *DescribeWorkflowExecutionRequest) GetExecution() (o *shared.WorkflowExecution) {
	if v != nil && v.Execution != nil {
		return v.Execution
	}

	return
}

// IsSetExecution returns true if Execution is not nil.
func (v *DescribeWorkflowExecutionRequest) IsSetExecution() bool {
	return v != nil && v.Execution != nil
}

type DescribeWorkflowExecutionResponse struct {
	ShardId                *string `json:"shardId,omitempty"`
	HistoryAddr            *string `json:"historyAddr,omitempty"`
	MutableStateInCache    *string `json:"mutableStateInCache,omitempty"`
	MutableStateInDatabase *string `json:"mutableStateInDatabase,omitempty"`
}

// ToWire translates a DescribeWorkflowExecutionResponse struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *DescribeWorkflowExecutionResponse) ToWire() (wire.Value, error) {
	var (
		fields [4]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.ShardId != nil {
		w, err = wire.NewValueString(*(v.ShardId)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 10, Value: w}
		i++
	}
	if v.HistoryAddr != nil {
		w, err = wire.NewValueString(*(v.HistoryAddr)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 20, Value: w}
		i++
	}
	if v.MutableStateInCache != nil {
		w, err = wire.NewValueString(*(v.MutableStateInCache)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 40, Value: w}
		i++
	}
	if v.MutableStateInDatabase != nil {
		w, err = wire.NewValueString(*(v.MutableStateInDatabase)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 50, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a DescribeWorkflowExecutionResponse struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a DescribeWorkflowExecutionResponse struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v DescribeWorkflowExecutionResponse
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *DescribeWorkflowExecutionResponse) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 10:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.ShardId = &x
				if err != nil {
					return err
				}

			}
		case 20:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.HistoryAddr = &x
				if err != nil {
					return err
				}

			}
		case 40:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.MutableStateInCache = &x
				if err != nil {
					return err
				}

			}
		case 50:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.MutableStateInDatabase = &x
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a DescribeWorkflowExecutionResponse
// struct.
func (v *DescribeWorkflowExecutionResponse) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [4]string
	i := 0
	if v.ShardId != nil {
		fields[i] = fmt.Sprintf("ShardId: %v", *(v.ShardId))
		i++
	}
	if v.HistoryAddr != nil {
		fields[i] = fmt.Sprintf("HistoryAddr: %v", *(v.HistoryAddr))
		i++
	}
	if v.MutableStateInCache != nil {
		fields[i] = fmt.Sprintf("MutableStateInCache: %v", *(v.MutableStateInCache))
		i++
	}
	if v.MutableStateInDatabase != nil {
		fields[i] = fmt.Sprintf("MutableStateInDatabase: %v", *(v.MutableStateInDatabase))
		i++
	}

	return fmt.Sprintf("DescribeWorkflowExecutionResponse{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this DescribeWorkflowExecutionResponse match the
// provided DescribeWorkflowExecutionResponse.
//
// This function performs a deep comparison.
func (v *DescribeWorkflowExecutionResponse) Equals(rhs *DescribeWorkflowExecutionResponse) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_String_EqualsPtr(v.ShardId, rhs.ShardId) {
		return false
	}
	if !_String_EqualsPtr(v.HistoryAddr, rhs.HistoryAddr) {
		return false
	}
	if !_String_EqualsPtr(v.MutableStateInCache, rhs.MutableStateInCache) {
		return false
	}
	if !_String_EqualsPtr(v.MutableStateInDatabase, rhs.MutableStateInDatabase) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of DescribeWorkflowExecutionResponse.
func (v *DescribeWorkflowExecutionResponse) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.ShardId != nil {
		enc.AddString("shardId", *v.ShardId)
	}
	if v.HistoryAddr != nil {
		enc.AddString("historyAddr", *v.HistoryAddr)
	}
	if v.MutableStateInCache != nil {
		enc.AddString("mutableStateInCache", *v.MutableStateInCache)
	}
	if v.MutableStateInDatabase != nil {
		enc.AddString("mutableStateInDatabase", *v.MutableStateInDatabase)
	}
	return err
}

// GetShardId returns the value of ShardId if it is set or its
// zero value if it is unset.
func (v *DescribeWorkflowExecutionResponse) GetShardId() (o string) {
	if v != nil && v.ShardId != nil {
		return *v.ShardId
	}

	return
}

// IsSetShardId returns true if ShardId is not nil.
func (v *DescribeWorkflowExecutionResponse) IsSetShardId() bool {
	return v != nil && v.ShardId != nil
}

// GetHistoryAddr returns the value of HistoryAddr if it is set or its
// zero value if it is unset.
func (v *DescribeWorkflowExecutionResponse) GetHistoryAddr() (o string) {
	if v != nil && v.HistoryAddr != nil {
		return *v.HistoryAddr
	}

	return
}

// IsSetHistoryAddr returns true if HistoryAddr is not nil.
func (v *DescribeWorkflowExecutionResponse) IsSetHistoryAddr() bool {
	return v != nil && v.HistoryAddr != nil
}

// GetMutableStateInCache returns the value of MutableStateInCache if it is set or its
// zero value if it is unset.
func (v *DescribeWorkflowExecutionResponse) GetMutableStateInCache() (o string) {
	if v != nil && v.MutableStateInCache != nil {
		return *v.MutableStateInCache
	}

	return
}

// IsSetMutableStateInCache returns true if MutableStateInCache is not nil.
func (v *DescribeWorkflowExecutionResponse) IsSetMutableStateInCache() bool {
	return v != nil && v.MutableStateInCache != nil
}

// GetMutableStateInDatabase returns the value of MutableStateInDatabase if it is set or its
// zero value if it is unset.
func (v *DescribeWorkflowExecutionResponse) GetMutableStateInDatabase() (o string) {
	if v != nil && v.MutableStateInDatabase != nil {
		return *v.MutableStateInDatabase
	}

	return
}

// IsSetMutableStateInDatabase returns true if MutableStateInDatabase is not nil.
func (v *DescribeWorkflowExecutionResponse) IsSetMutableStateInDatabase() bool {
	return v != nil && v.MutableStateInDatabase != nil
}

type GetWorkflowExecutionRawHistoryRequest struct {
	Domain          *string                   `json:"domain,omitempty"`
	Execution       *shared.WorkflowExecution `json:"execution,omitempty"`
	FirstEventId    *int64                    `json:"firstEventId,omitempty"`
	NextEventId     *int64                    `json:"nextEventId,omitempty"`
	MaximumPageSize *int32                    `json:"maximumPageSize,omitempty"`
	NextPageToken   []byte                    `json:"nextPageToken,omitempty"`
}

// ToWire translates a GetWorkflowExecutionRawHistoryRequest struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *GetWorkflowExecutionRawHistoryRequest) ToWire() (wire.Value, error) {
	var (
		fields [6]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Domain != nil {
		w, err = wire.NewValueString(*(v.Domain)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 10, Value: w}
		i++
	}
	if v.Execution != nil {
		w, err = v.Execution.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 20, Value: w}
		i++
	}
	if v.FirstEventId != nil {
		w, err = wire.NewValueI64(*(v.FirstEventId)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 30, Value: w}
		i++
	}
	if v.NextEventId != nil {
		w, err = wire.NewValueI64(*(v.NextEventId)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 40, Value: w}
		i++
	}
	if v.MaximumPageSize != nil {
		w, err = wire.NewValueI32(*(v.MaximumPageSize)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 50, Value: w}
		i++
	}
	if v.NextPageToken != nil {
		w, err = wire.NewValueBinary(v.NextPageToken), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 60, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a GetWorkflowExecutionRawHistoryRequest struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a GetWorkflowExecutionRawHistoryRequest struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v GetWorkflowExecutionRawHistoryRequest
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *GetWorkflowExecutionRawHistoryRequest) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 10:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Domain = &x
				if err != nil {
					return err
				}

			}
		case 20:
			if field.Value.Type() == wire.TStruct {
				v.Execution, err = _WorkflowExecution_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 30:
			if field.Value.Type() == wire.TI64 {
				var x int64
				x, err = field.Value.GetI64(), error(nil)
				v.FirstEventId = &x
				if err != nil {
					return err
				}

			}
		case 40:
			if field.Value.Type() == wire.TI64 {
				var x int64
				x, err = field.Value.GetI64(), error(nil)
				v.NextEventId = &x
				if err != nil {
					return err
				}

			}
		case 50:
			if field.Value.Type() == wire.TI32 {
				var x int32
				x, err = field.Value.GetI32(), error(nil)
				v.MaximumPageSize = &x
				if err != nil {
					return err
				}

			}
		case 60:
			if field.Value.Type() == wire.TBinary {
				v.NextPageToken, err = field.Value.GetBinary(), error(nil)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a GetWorkflowExecutionRawHistoryRequest
// struct.
func (v *GetWorkflowExecutionRawHistoryRequest) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [6]string
	i := 0
	if v.Domain != nil {
		fields[i] = fmt.Sprintf("Domain: %v", *(v.Domain))
		i++
	}
	if v.Execution != nil {
		fields[i] = fmt.Sprintf("Execution: %v", v.Execution)
		i++
	}
	if v.FirstEventId != nil {
		fields[i] = fmt.Sprintf("FirstEventId: %v", *(v.FirstEventId))
		i++
	}
	if v.NextEventId != nil {
		fields[i] = fmt.Sprintf("NextEventId: %v", *(v.NextEventId))
		i++
	}
	if v.MaximumPageSize != nil {
		fields[i] = fmt.Sprintf("MaximumPageSize: %v", *(v.MaximumPageSize))
		i++
	}
	if v.NextPageToken != nil {
		fields[i] = fmt.Sprintf("NextPageToken: %v", v.NextPageToken)
		i++
	}

	return fmt.Sprintf("GetWorkflowExecutionRawHistoryRequest{%v}", strings.Join(fields[:i], ", "))
}

func _I64_EqualsPtr(lhs, rhs *int64) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

func _I32_EqualsPtr(lhs, rhs *int32) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this GetWorkflowExecutionRawHistoryRequest match the
// provided GetWorkflowExecutionRawHistoryRequest.
//
// This function performs a deep comparison.
func (v *GetWorkflowExecutionRawHistoryRequest) Equals(rhs *GetWorkflowExecutionRawHistoryRequest) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_String_EqualsPtr(v.Domain, rhs.Domain) {
		return false
	}
	if !((v.Execution == nil && rhs.Execution == nil) || (v.Execution != nil && rhs.Execution != nil && v.Execution.Equals(rhs.Execution))) {
		return false
	}
	if !_I64_EqualsPtr(v.FirstEventId, rhs.FirstEventId) {
		return false
	}
	if !_I64_EqualsPtr(v.NextEventId, rhs.NextEventId) {
		return false
	}
	if !_I32_EqualsPtr(v.MaximumPageSize, rhs.MaximumPageSize) {
		return false
	}
	if !((v.NextPageToken == nil && rhs.NextPageToken == nil) || (v.NextPageToken != nil && rhs.NextPageToken != nil && bytes.Equal(v.NextPageToken, rhs.NextPageToken))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of GetWorkflowExecutionRawHistoryRequest.
func (v *GetWorkflowExecutionRawHistoryRequest) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Domain != nil {
		enc.AddString("domain", *v.Domain)
	}
	if v.Execution != nil {
		err = multierr.Append(err, enc.AddObject("execution", v.Execution))
	}
	if v.FirstEventId != nil {
		enc.AddInt64("firstEventId", *v.FirstEventId)
	}
	if v.NextEventId != nil {
		enc.AddInt64("nextEventId", *v.NextEventId)
	}
	if v.MaximumPageSize != nil {
		enc.AddInt32("maximumPageSize", *v.MaximumPageSize)
	}
	if v.NextPageToken != nil {
		enc.AddString("nextPageToken", base64.StdEncoding.EncodeToString(v.NextPageToken))
	}
	return err
}

// GetDomain returns the value of Domain if it is set or its
// zero value if it is unset.
func (v *GetWorkflowExecutionRawHistoryRequest) GetDomain() (o string) {
	if v != nil && v.Domain != nil {
		return *v.Domain
	}

	return
}

// IsSetDomain returns true if Domain is not nil.
func (v *GetWorkflowExecutionRawHistoryRequest) IsSetDomain() bool {
	return v != nil && v.Domain != nil
}

// GetExecution returns the value of Execution if it is set or its
// zero value if it is unset.
func (v *GetWorkflowExecutionRawHistoryRequest) GetExecution() (o *shared.WorkflowExecution) {
	if v != nil && v.Execution != nil {
		return v.Execution
	}

	return
}

// IsSetExecution returns true if Execution is not nil.
func (v *GetWorkflowExecutionRawHistoryRequest) IsSetExecution() bool {
	return v != nil && v.Execution != nil
}

// GetFirstEventId returns the value of FirstEventId if it is set or its
// zero value if it is unset.
func (v *GetWorkflowExecutionRawHistoryRequest) GetFirstEventId() (o int64) {
	if v != nil && v.FirstEventId != nil {
		return *v.FirstEventId
	}

	return
}

// IsSetFirstEventId returns true if FirstEventId is not nil.
func (v *GetWorkflowExecutionRawHistoryRequest) IsSetFirstEventId() bool {
	return v != nil && v.FirstEventId != nil
}

// GetNextEventId returns the value of NextEventId if it is set or its
// zero value if it is unset.
func (v *GetWorkflowExecutionRawHistoryRequest) GetNextEventId() (o int64) {
	if v != nil && v.NextEventId != nil {
		return *v.NextEventId
	}

	return
}

// IsSetNextEventId returns true if NextEventId is not nil.
func (v *GetWorkflowExecutionRawHistoryRequest) IsSetNextEventId() bool {
	return v != nil && v.NextEventId != nil
}

// GetMaximumPageSize returns the value of MaximumPageSize if it is set or its
// zero value if it is unset.
func (v *GetWorkflowExecutionRawHistoryRequest) GetMaximumPageSize() (o int32) {
	if v != nil && v.MaximumPageSize != nil {
		return *v.MaximumPageSize
	}

	return
}

// IsSetMaximumPageSize returns true if MaximumPageSize is not nil.
func (v *GetWorkflowExecutionRawHistoryRequest) IsSetMaximumPageSize() bool {
	return v != nil && v.MaximumPageSize != nil
}

// GetNextPageToken returns the value of NextPageToken if it is set or its
// zero value if it is unset.
func (v *GetWorkflowExecutionRawHistoryRequest) GetNextPageToken() (o []byte) {
	if v != nil && v.NextPageToken != nil {
		return v.NextPageToken
	}

	return
}

// IsSetNextPageToken returns true if NextPageToken is not nil.
func (v *GetWorkflowExecutionRawHistoryRequest) IsSetNextPageToken() bool {
	return v != nil && v.NextPageToken != nil
}

type GetWorkflowExecutionRawHistoryResponse struct {
	NextPageToken     []byte                             `json:"nextPageToken,omitempty"`
	HistoryBatches    []*shared.DataBlob                 `json:"historyBatches,omitempty"`
	ReplicationInfo   map[string]*shared.ReplicationInfo `json:"replicationInfo,omitempty"`
	EventStoreVersion *int32                             `json:"eventStoreVersion,omitempty"`
}

type _List_DataBlob_ValueList []*shared.DataBlob

func (v _List_DataBlob_ValueList) ForEach(f func(wire.Value) error) error {
	for i, x := range v {
		if x == nil {
			return fmt.Errorf("invalid [%v]: value is nil", i)
		}
		w, err := x.ToWire()
		if err != nil {
			return err
		}
		err = f(w)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v _List_DataBlob_ValueList) Size() int {
	return len(v)
}

func (_List_DataBlob_ValueList) ValueType() wire.Type {
	return wire.TStruct
}

func (_List_DataBlob_ValueList) Close() {}

type _Map_String_ReplicationInfo_MapItemList map[string]*shared.ReplicationInfo

func (m _Map_String_ReplicationInfo_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for k, v := range m {
		if v == nil {
			return fmt.Errorf("invalid [%v]: value is nil", k)
		}
		kw, err := wire.NewValueString(k), error(nil)
		if err != nil {
			return err
		}

		vw, err := v.ToWire()
		if err != nil {
			return err
		}
		err = f(wire.MapItem{Key: kw, Value: vw})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m _Map_String_ReplicationInfo_MapItemList) Size() int {
	return len(m)
}

func (_Map_String_ReplicationInfo_MapItemList) KeyType() wire.Type {
	return wire.TBinary
}

func (_Map_String_ReplicationInfo_MapItemList) ValueType() wire.Type {
	return wire.TStruct
}

func (_Map_String_ReplicationInfo_MapItemList) Close() {}

// ToWire translates a GetWorkflowExecutionRawHistoryResponse struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *GetWorkflowExecutionRawHistoryResponse) ToWire() (wire.Value, error) {
	var (
		fields [4]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.NextPageToken != nil {
		w, err = wire.NewValueBinary(v.NextPageToken), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 10, Value: w}
		i++
	}
	if v.HistoryBatches != nil {
		w, err = wire.NewValueList(_List_DataBlob_ValueList(v.HistoryBatches)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 20, Value: w}
		i++
	}
	if v.ReplicationInfo != nil {
		w, err = wire.NewValueMap(_Map_String_ReplicationInfo_MapItemList(v.ReplicationInfo)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 30, Value: w}
		i++
	}
	if v.EventStoreVersion != nil {
		w, err = wire.NewValueI32(*(v.EventStoreVersion)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 40, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _DataBlob_Read(w wire.Value) (*shared.DataBlob, error) {
	var v shared.DataBlob
	err := v.FromWire(w)
	return &v, err
}

func _List_DataBlob_Read(l wire.ValueList) ([]*shared.DataBlob, error) {
	if l.ValueType() != wire.TStruct {
		return nil, nil
	}

	o := make([]*shared.DataBlob, 0, l.Size())
	err := l.ForEach(func(x wire.Value) error {
		i, err := _DataBlob_Read(x)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Close()
	return o, err
}

func _ReplicationInfo_Read(w wire.Value) (*shared.ReplicationInfo, error) {
	var v shared.ReplicationInfo
	err := v.FromWire(w)
	return &v, err
}

func _Map_String_ReplicationInfo_Read(m wire.MapItemList) (map[string]*shared.ReplicationInfo, error) {
	if m.KeyType() != wire.TBinary {
		return nil, nil
	}

	if m.ValueType() != wire.TStruct {
		return nil, nil
	}

	o := make(map[string]*shared.ReplicationInfo, m.Size())
	err := m.ForEach(func(x wire.MapItem) error {
		k, err := x.Key.GetString(), error(nil)
		if err != nil {
			return err
		}

		v, err := _ReplicationInfo_Read(x.Value)
		if err != nil {
			return err
		}

		o[k] = v
		return nil
	})
	m.Close()
	return o, err
}

// FromWire deserializes a GetWorkflowExecutionRawHistoryResponse struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a GetWorkflowExecutionRawHistoryResponse struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v GetWorkflowExecutionRawHistoryResponse
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *GetWorkflowExecutionRawHistoryResponse) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 10:
			if field.Value.Type() == wire.TBinary {
				v.NextPageToken, err = field.Value.GetBinary(), error(nil)
				if err != nil {
					return err
				}

			}
		case 20:
			if field.Value.Type() == wire.TList {
				v.HistoryBatches, err = _List_DataBlob_Read(field.Value.GetList())
				if err != nil {
					return err
				}

			}
		case 30:
			if field.Value.Type() == wire.TMap {
				v.ReplicationInfo, err = _Map_String_ReplicationInfo_Read(field.Value.GetMap())
				if err != nil {
					return err
				}

			}
		case 40:
			if field.Value.Type() == wire.TI32 {
				var x int32
				x, err = field.Value.GetI32(), error(nil)
				v.EventStoreVersion = &x
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a GetWorkflowExecutionRawHistoryResponse
// struct.
func (v *GetWorkflowExecutionRawHistoryResponse) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [4]string
	i := 0
	if v.NextPageToken != nil {
		fields[i] = fmt.Sprintf("NextPageToken: %v", v.NextPageToken)
		i++
	}
	if v.HistoryBatches != nil {
		fields[i] = fmt.Sprintf("HistoryBatches: %v", v.HistoryBatches)
		i++
	}
	if v.ReplicationInfo != nil {
		fields[i] = fmt.Sprintf("ReplicationInfo: %v", v.ReplicationInfo)
		i++
	}
	if v.EventStoreVersion != nil {
		fields[i] = fmt.Sprintf("EventStoreVersion: %v", *(v.EventStoreVersion))
		i++
	}

	return fmt.Sprintf("GetWorkflowExecutionRawHistoryResponse{%v}", strings.Join(fields[:i], ", "))
}

func _List_DataBlob_Equals(lhs, rhs []*shared.DataBlob) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for i, lv := range lhs {
		rv := rhs[i]
		if !lv.Equals(rv) {
			return false
		}
	}

	return true
}

func _Map_String_ReplicationInfo_Equals(lhs, rhs map[string]*shared.ReplicationInfo) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for lk, lv := range lhs {
		rv, ok := rhs[lk]
		if !ok {
			return false
		}
		if !lv.Equals(rv) {
			return false
		}
	}
	return true
}

// Equals returns true if all the fields of this GetWorkflowExecutionRawHistoryResponse match the
// provided GetWorkflowExecutionRawHistoryResponse.
//
// This function performs a deep comparison.
func (v *GetWorkflowExecutionRawHistoryResponse) Equals(rhs *GetWorkflowExecutionRawHistoryResponse) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.NextPageToken == nil && rhs.NextPageToken == nil) || (v.NextPageToken != nil && rhs.NextPageToken != nil && bytes.Equal(v.NextPageToken, rhs.NextPageToken))) {
		return false
	}
	if !((v.HistoryBatches == nil && rhs.HistoryBatches == nil) || (v.HistoryBatches != nil && rhs.HistoryBatches != nil && _List_DataBlob_Equals(v.HistoryBatches, rhs.HistoryBatches))) {
		return false
	}
	if !((v.ReplicationInfo == nil && rhs.ReplicationInfo == nil) || (v.ReplicationInfo != nil && rhs.ReplicationInfo != nil && _Map_String_ReplicationInfo_Equals(v.ReplicationInfo, rhs.ReplicationInfo))) {
		return false
	}
	if !_I32_EqualsPtr(v.EventStoreVersion, rhs.EventStoreVersion) {
		return false
	}

	return true
}

type _List_DataBlob_Zapper []*shared.DataBlob

// MarshalLogArray implements zapcore.ArrayMarshaler, enabling
// fast logging of _List_DataBlob_Zapper.
func (l _List_DataBlob_Zapper) MarshalLogArray(enc zapcore.ArrayEncoder) (err error) {
	for _, v := range l {
		err = multierr.Append(err, enc.AppendObject(v))
	}
	return err
}

type _Map_String_ReplicationInfo_Zapper map[string]*shared.ReplicationInfo

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of _Map_String_ReplicationInfo_Zapper.
func (m _Map_String_ReplicationInfo_Zapper) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	for k, v := range m {
		err = multierr.Append(err, enc.AddObject((string)(k), v))
	}
	return err
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of GetWorkflowExecutionRawHistoryResponse.
func (v *GetWorkflowExecutionRawHistoryResponse) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.NextPageToken != nil {
		enc.AddString("nextPageToken", base64.StdEncoding.EncodeToString(v.NextPageToken))
	}
	if v.HistoryBatches != nil {
		err = multierr.Append(err, enc.AddArray("historyBatches", (_List_DataBlob_Zapper)(v.HistoryBatches)))
	}
	if v.ReplicationInfo != nil {
		err = multierr.Append(err, enc.AddObject("replicationInfo", (_Map_String_ReplicationInfo_Zapper)(v.ReplicationInfo)))
	}
	if v.EventStoreVersion != nil {
		enc.AddInt32("eventStoreVersion", *v.EventStoreVersion)
	}
	return err
}

// GetNextPageToken returns the value of NextPageToken if it is set or its
// zero value if it is unset.
func (v *GetWorkflowExecutionRawHistoryResponse) GetNextPageToken() (o []byte) {
	if v != nil && v.NextPageToken != nil {
		return v.NextPageToken
	}

	return
}

// IsSetNextPageToken returns true if NextPageToken is not nil.
func (v *GetWorkflowExecutionRawHistoryResponse) IsSetNextPageToken() bool {
	return v != nil && v.NextPageToken != nil
}

// GetHistoryBatches returns the value of HistoryBatches if it is set or its
// zero value if it is unset.
func (v *GetWorkflowExecutionRawHistoryResponse) GetHistoryBatches() (o []*shared.DataBlob) {
	if v != nil && v.HistoryBatches != nil {
		return v.HistoryBatches
	}

	return
}

// IsSetHistoryBatches returns true if HistoryBatches is not nil.
func (v *GetWorkflowExecutionRawHistoryResponse) IsSetHistoryBatches() bool {
	return v != nil && v.HistoryBatches != nil
}

// GetReplicationInfo returns the value of ReplicationInfo if it is set or its
// zero value if it is unset.
func (v *GetWorkflowExecutionRawHistoryResponse) GetReplicationInfo() (o map[string]*shared.ReplicationInfo) {
	if v != nil && v.ReplicationInfo != nil {
		return v.ReplicationInfo
	}

	return
}

// IsSetReplicationInfo returns true if ReplicationInfo is not nil.
func (v *GetWorkflowExecutionRawHistoryResponse) IsSetReplicationInfo() bool {
	return v != nil && v.ReplicationInfo != nil
}

// GetEventStoreVersion returns the value of EventStoreVersion if it is set or its
// zero value if it is unset.
func (v *GetWorkflowExecutionRawHistoryResponse) GetEventStoreVersion() (o int32) {
	if v != nil && v.EventStoreVersion != nil {
		return *v.EventStoreVersion
	}

	return
}

// IsSetEventStoreVersion returns true if EventStoreVersion is not nil.
func (v *GetWorkflowExecutionRawHistoryResponse) IsSetEventStoreVersion() bool {
	return v != nil && v.EventStoreVersion != nil
}
