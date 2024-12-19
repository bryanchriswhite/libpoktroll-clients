package main

/*
#cgo CFLAGS: -I${SRCDIR}/../include
#include <memory.h>
#include <protobuf.h>
*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/cosmos/cosmos-sdk/types"
	gogoproto "github.com/cosmos/gogoproto/proto"
)

const (
	// TODO_IN_THIS_COMMIT: godoc...
	NilGoRef = GoRef(-1)
	// TODO_IN_THIS_COMMIT: godoc...
	ZeroGoRef = GoRef(0)
)

var (
	goMemoryMap  = map[GoRef]any{}
	nextGoMemRef = GoRef(0)
)

type GoRef int64

// TODO_IN_THIS_COMMIT: godoc...
func SetGoMem(value any) C.go_ref {
	nextGoMemRef++
	goMemoryMap[nextGoMemRef] = value
	return C.go_ref(nextGoMemRef)
}

// TODO_IN_THIS_COMMIT: godoc...
func GetGoMem[T any](ref C.go_ref) (T, error) {
	value, ok := goMemoryMap[GoRef(ref)]
	if !ok {
		return *new(T), fmt.Errorf("go memory reference not found: %d", ref)
	}

	valueT, ok := value.(T)
	if !ok {
		return valueT, fmt.Errorf("expected %T, got: %T", *new(T), value)
	}

	return valueT, nil
}

//export GetGoProtoAsSerializedProto
func GetGoProtoAsSerializedProto(ref C.go_ref, cErr **C.char) unsafe.Pointer {
	value, ok := goMemoryMap[GoRef(ref)]
	if !ok {
		return C.NULL
	}

	proto_value, ok := value.(gogoproto.Message)
	if !ok {
		*cErr = C.CString(fmt.Sprintf("expected proto value, got: %T", value))
		return C.NULL
	}

	proto_bz, err := cdc.Marshal(proto_value)
	if err != nil {
		*cErr = C.CString(err.Error())
		return C.NULL
	}

	cSerializedProto := C.malloc(C.size_t(unsafe.Sizeof(C.serialized_proto{})))
	*(*C.serialized_proto)(cSerializedProto) = C.serialized_proto{
		type_url:        (*C.uint8_t)(C.CBytes([]byte(types.MsgTypeURL(proto_value)))),
		type_url_length: C.size_t(len(types.MsgTypeURL(proto_value))),
		data:            (*C.uint8_t)(C.CBytes(proto_bz)),
		data_length:     C.size_t(len(proto_bz)),
	}

	return unsafe.Pointer(cSerializedProto)
}

// TODO_IN_THIS_COMMIT: godoc...
//
//export FreeGoMem
func FreeGoMem(ref C.go_ref) {
	delete(goMemoryMap, GoRef(ref))
}
