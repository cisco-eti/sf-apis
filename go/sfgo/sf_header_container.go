// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCES:
 *     GraphletRecord.avsc
 *     SysFlow.avsc
 */
package sfgo

import (
	"io"

	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/container"
	"github.com/actgardner/gogen-avro/v7/vm"
)

func NewSFHeaderWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewSFHeader()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type SFHeaderReader struct {
	r io.Reader
	p *vm.Program
}

func NewSFHeaderReader(r io.Reader) (*SFHeaderReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewSFHeader()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &SFHeaderReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r SFHeaderReader) Read() (*SFHeader, error) {
	t := NewSFHeader()
	err := vm.Eval(r.r, r.p, t)
	return t, err
}