// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     SysFlow.avsc
 */
package sfgo

import (
	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
	"io"
)

type NetworkFlow struct {
	ProcOID *OID `json:"procOID"`

	Ts int64 `json:"ts"`

	Tid int64 `json:"tid"`

	OpFlags int32 `json:"opFlags"`

	EndTs int64 `json:"endTs"`

	Sip int32 `json:"sip"`

	Sport int32 `json:"sport"`

	Dip int32 `json:"dip"`

	Dport int32 `json:"dport"`

	Proto int32 `json:"proto"`

	Fd int32 `json:"fd"`

	NumRRecvOps int64 `json:"numRRecvOps"`

	NumWSendOps int64 `json:"numWSendOps"`

	NumRRecvBytes int64 `json:"numRRecvBytes"`

	NumWSendBytes int64 `json:"numWSendBytes"`
}

const NetworkFlowAvroCRC64Fingerprint = "K\xd5\x14\xceg\xe2 \xd3"

func NewNetworkFlow() *NetworkFlow {
	return &NetworkFlow{}
}

func DeserializeNetworkFlow(r io.Reader) (*NetworkFlow, error) {
	t := NewNetworkFlow()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func DeserializeNetworkFlowFromSchema(r io.Reader, schema string) (*NetworkFlow, error) {
	t := NewNetworkFlow()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func writeNetworkFlow(r *NetworkFlow, w io.Writer) error {
	var err error
	err = writeOID(r.ProcOID, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Ts, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Tid, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.OpFlags, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.EndTs, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Sip, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Sport, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Dip, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Dport, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Proto, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Fd, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.NumRRecvOps, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.NumWSendOps, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.NumRRecvBytes, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.NumWSendBytes, w)
	if err != nil {
		return err
	}
	return err
}

func (r *NetworkFlow) Serialize(w io.Writer) error {
	return writeNetworkFlow(r, w)
}

func (r *NetworkFlow) Schema() string {
	return "{\"fields\":[{\"name\":\"procOID\",\"type\":{\"fields\":[{\"name\":\"createTS\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"hpid\",\"type\":\"long\"}],\"name\":\"OID\",\"namespace\":\"sysflow.type\",\"type\":\"record\"}},{\"name\":\"ts\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"tid\",\"type\":\"long\"},{\"name\":\"opFlags\",\"type\":\"int\"},{\"name\":\"endTs\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"sip\",\"type\":\"int\"},{\"name\":\"sport\",\"type\":\"int\"},{\"name\":\"dip\",\"type\":\"int\"},{\"name\":\"dport\",\"type\":\"int\"},{\"name\":\"proto\",\"type\":\"int\"},{\"name\":\"fd\",\"type\":\"int\"},{\"name\":\"numRRecvOps\",\"type\":\"long\"},{\"name\":\"numWSendOps\",\"type\":\"long\"},{\"name\":\"numRRecvBytes\",\"type\":\"long\"},{\"name\":\"numWSendBytes\",\"type\":\"long\"}],\"name\":\"sysflow.flow.NetworkFlow\",\"type\":\"record\"}"
}

func (r *NetworkFlow) SchemaName() string {
	return "sysflow.flow.NetworkFlow"
}

func (_ *NetworkFlow) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *NetworkFlow) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *NetworkFlow) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *NetworkFlow) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *NetworkFlow) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *NetworkFlow) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *NetworkFlow) SetString(v string)   { panic("Unsupported operation") }
func (_ *NetworkFlow) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NetworkFlow) Get(i int) types.Field {
	switch i {
	case 0:
		r.ProcOID = NewOID()

		return r.ProcOID
	case 1:
		return &types.Long{Target: &r.Ts}
	case 2:
		return &types.Long{Target: &r.Tid}
	case 3:
		return &types.Int{Target: &r.OpFlags}
	case 4:
		return &types.Long{Target: &r.EndTs}
	case 5:
		return &types.Int{Target: &r.Sip}
	case 6:
		return &types.Int{Target: &r.Sport}
	case 7:
		return &types.Int{Target: &r.Dip}
	case 8:
		return &types.Int{Target: &r.Dport}
	case 9:
		return &types.Int{Target: &r.Proto}
	case 10:
		return &types.Int{Target: &r.Fd}
	case 11:
		return &types.Long{Target: &r.NumRRecvOps}
	case 12:
		return &types.Long{Target: &r.NumWSendOps}
	case 13:
		return &types.Long{Target: &r.NumRRecvBytes}
	case 14:
		return &types.Long{Target: &r.NumWSendBytes}
	}
	panic("Unknown field index")
}

func (r *NetworkFlow) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *NetworkFlow) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ *NetworkFlow) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *NetworkFlow) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *NetworkFlow) Finalize()                        {}

func (_ *NetworkFlow) AvroCRC64Fingerprint() []byte {
	return []byte(NetworkFlowAvroCRC64Fingerprint)
}
