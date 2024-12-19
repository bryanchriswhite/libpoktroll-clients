package main

/*
#include <memory.h>
#include <context.h>
#include <protobuf.h>
*/
import "C"
import (
	"context"

	"cosmossdk.io/depinject"
)

// TODO_IN_THIS_COMMIT: godoc...
//
//export NewQueryClient
func NewQueryClient(depsRef C.go_ref, queryNodeRPCURL *C.char, cErr **C.char) C.go_ref {
	deps, err := GetGoMem[depinject.Config](depsRef)
	if err != nil {
		*cErr = C.CString(err.Error())
		return C.go_ref(NilGoRef)
	}

	multiClient, err := NewMultiQueryClient(deps, C.GoString(queryNodeRPCURL))
	if err != nil {
		*cErr = C.CString(err.Error())
		return C.go_ref(NilGoRef)
	}

	return SetGoMem(multiClient)
}

// TODO_IN_THIS_COMMIT: godoc...
//
//export QueryClient_GetSharedParams
func QueryClient_GetSharedParams(depsRef C.go_ref, cErr **C.char) C.go_ref {
	multiClient, err := GetGoMem[MultiQueryClient](depsRef)
	if err != nil {
		*cErr = C.CString(err.Error())
		return C.go_ref(ZeroGoRef)
	}

	sharedParams, err := multiClient.GetSharedParams(context.TODO())
	if err != nil {
		*cErr = C.CString(err.Error())
		return C.go_ref(ZeroGoRef)
	}

	return SetGoMem(sharedParams)
}

// TODO_IN_THIS_COMMIT: godoc...
//
//export QueryClient_GetApplicationParams
func QueryClient_GetApplicationParams(depsRef C.go_ref, op *C.AsyncOperation) {

}

// TODO_IN_THIS_COMMIT: godoc...
//
//export QueryClient_GetGetewayParams
func QueryClient_GetGetewayParams(depsRef C.go_ref, op *C.AsyncOperation) {

}

// TODO_IN_THIS_COMMIT: godoc...
//
//export QueryClient_GetSupplierParams
func QueryClient_GetSupplierParams(depsRef C.go_ref, op *C.AsyncOperation) {

}

// TODO_IN_THIS_COMMIT: godoc...
//
//export QueryClient_GetSessionParams
func QueryClient_GetSessionParams(depsRef C.go_ref, op *C.AsyncOperation) {

}

// TODO_IN_THIS_COMMIT: godoc...
//
//export QueryClient_GetServiceParams
func QueryClient_GetServiceParams(depsRef C.go_ref, op *C.AsyncOperation) {

}

// TODO_IN_THIS_COMMIT: godoc...
//
//export QueryClient_GetProofParams
func QueryClient_GetProofParams(depsRef C.go_ref, op *C.AsyncOperation) {

}

// TODO_IN_THIS_COMMIT: godoc...
//
//export QueryClient_GetTokenomicsParams
func QueryClient_GetTokenomicsParams(depsRef C.go_ref, op *C.AsyncOperation) {

}
