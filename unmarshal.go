package kdl

import (
	"io"

	"github.com/sblinch/kdl-go/document"
	"github.com/sblinch/kdl-go/internal/marshaler"
	"github.com/sblinch/kdl-go/internal/tokenizer"
)

// Unmarshaler provides an interface for custom unmarshaling of a node into a Go type
type Unmarshaler interface {
	UnmarshalKDL(node *document.Node) error
}

// ValueUnmarshaler provides an interface for custom unmarshaling of a Value (such as a node argument or property) into
// a Go type
type ValueUnmarshaler interface {
	UnmarshalKDLValue(value *document.Value) error
}

// Decoder implements a decoder for KDL
type Decoder struct {
	r       io.Reader
	Options marshaler.UnmarshalOptions
}

// Decode decodes KDL from the Decoder's reader into v; v must contain a pointer type. Returns a non-nil error on
// failure.
func (d *Decoder) Decode(v interface{}) error {
	s := tokenizer.New(d.r)
	s.RelaxedNonCompliant = d.Options.RelaxedNonCompliant
	if doc, err := parse(s); err != nil {
		return err
	} else {
		return marshaler.UnmarshalWithOptions(doc, v, d.Options)
	}
}

// NewDecoder returns a Decoder that reads from r
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

// Unmarshal unmarshals KDL from data into v; v must contain a pointer type. Returns a non-nil error on failure.
func Unmarshal(data []byte, v interface{}) error {
	s := tokenizer.NewSlice(data)
	if doc, err := parse(s); err != nil {
		return err
	} else {
		return marshaler.Unmarshal(doc, v)
	}
}
