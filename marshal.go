package kdl

import (
	"bytes"
	"io"

	"github.com/sblinch/kdl-go/document"
	"github.com/sblinch/kdl-go/internal/generator"
	"github.com/sblinch/kdl-go/internal/marshaler"
)

// Marshaler provides an interface for custom marshaling of a Go type into a Node
type Marshaler interface {
	MarshalKDL(node *document.Node) error
}

// ValueMarshaler provides an interface for custom marshaling of a Go type into a Value (such as a node argument or
// property)
type ValueMarshaler interface {
	MarshalKDLValue(value *document.Value) error
}

// Encoder implements an encoder for KDL
type Encoder struct {
	w       io.Writer
	Options marshaler.MarshalOptions
}

// Encode encodes v into KDL and writes it to the Encoder's writer, and returns a non-nil error on failure
func (e *Encoder) Encode(v interface{}) error {
	doc := document.New()
	if err := marshaler.MarshalWithOptions(v, doc, e.Options); err != nil {
		return err
	}

	g := generator.New(e.w)
	if err := g.Generate(doc); err != nil {
		return err
	}

	return nil
}

// NewEncoder creates a new Encoder that writes to w
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

// Marshal returns the KDL representation of v, or a non-nil error on failure
func Marshal(v interface{}) ([]byte, error) {
	doc := document.New()
	if err := marshaler.Marshal(v, doc); err != nil {
		return nil, err
	}

	b := bytes.Buffer{}
	g := generator.New(&b)
	if err := g.Generate(doc); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
