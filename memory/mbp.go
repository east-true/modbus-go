package memory

import (
	"strings"

	"github.com/east-true/goutil/parser"
)

var TypeMap map[string]uint8 = map[string]uint8{
	"bit16":   parser.BIT16,
	"bit32":   parser.BIT32,
	"bit64":   parser.BIT64,
	"int16":   parser.INT16,
	"int32":   parser.INT32,
	"int64":   parser.INT64,
	"uint16":  parser.UINT16,
	"uint32":  parser.UINT32,
	"uint64":  parser.UINT64,
	"float32": parser.FLOAT32,
	"float64": parser.FLOAT64,
}

// MBP : modbus parser
type MBP struct {
	t uint8
	p *parser.Parser
}

func NewParser(order, dataType string) *MBP {
	lower := strings.ToLower(order)
	return &MBP{
		t: TypeMap[dataType],
		p: parser.New(lower),
	}
}

func (mbp *MBP) ToBit(b []byte) []interface{} {
	bits := mbp.p.ToBitArr(b)
	wrapper := make([]interface{}, len(bits))
	for i := range wrapper {
		wrapper[i] = bits[i]
	}
	return wrapper
}

func (mbp *MBP) ToData(b []byte) ([]interface{}, error) {
	return mpb.p.ToAnyOf(b) // FIXME : parser.ToAnyOf reutrn any....
}
