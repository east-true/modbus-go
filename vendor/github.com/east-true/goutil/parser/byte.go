package parser

const (
	INT16 uint8 = iota
	UINT16
	INT32
	UINT32
	INT64
	UINT64
	FLOAT32
	FLOAT64
	INT16ARR
	UINT16ARR
	INT32ARR
	UINT32ARR
	INT64ARR
	UINT64ARR
	FLOAT32ARR
	FLOAT64ARR

	LITTLE_LOWER string = "little"
	BIG_LOWER    string = "big"
)

type Parser interface {
	ToBit(datum byte) []uint8
	ToBitArr(datum []byte) []uint8
	ToInt16(b []byte) int16
	ToInt16Arr(b []byte) ([]int16, error)
	ToUint16(b []byte) uint16
	ToUint16Arr(b []byte) ([]uint16, error)
	ToInt32(b []byte) int32
	ToInt32Arr(b []byte) ([]int32, error)
	ToUint32(b []byte) uint32
	ToUint32Arr(b []byte) ([]uint32, error)
	ToInt64(b []byte) int64
	ToInt64Arr(b []byte) ([]int64, error)
	ToUint64(b []byte) uint64
	ToUint64Arr(b []byte) ([]uint64, error)
	ToFloat32(b []byte) float32
	ToFloat32Arr(b []byte) ([]float32, error)
	ToFloat64(b []byte) float64
	ToFloat64Arr(b []byte) ([]float64, error)
}

func New(order string) Parser {
	if order == LITTLE_LOWER {
		return &LittleByteParser{}
	} else {
		return &BigByteParser{}
	}
}
