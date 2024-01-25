package query_validator

/*
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"fmt"
	"github.com/Mamoru-Foundation/mamoru-core-go/generated_bindings"
	"unsafe"
)

type Chain byte

const (
	ChainSui    Chain = generated_bindings.FfiChainTypeSui
	ChainEvm    Chain = generated_bindings.FfiChainTypeEvm
	ChainAptos  Chain = generated_bindings.FfiChainTypeAptos
	ChainCosmos Chain = generated_bindings.FfiChainTypeCosmos
)

func ValidateSql(chain Chain, query string, parameters map[string]string, versions map[string]string) error {
	ffiVersions, err := daemonVersions(versions)

	if err != nil {
		return fmt.Errorf("error validating versions: %w", err)
	}

	result := generated_bindings.FfiValidateSql(
		generated_bindings.FfiChainTypeT(chain),
		query,
		daemonParameters(parameters),
		ffiVersions,
	)

	return makeError(result)
}

func ValidateSqlRenders(query string, parameters map[string]string, versions map[string]string) error {
	ffiVersions, err := daemonVersions(versions)

	if err != nil {
		return fmt.Errorf("error validating versions: %w", err)
	}

	result := generated_bindings.FfiValidateSqlRenders(query, daemonParameters(parameters), ffiVersions)

	return makeError(result)
}

func ValidateAssemblyScript(chain Chain, bytes []byte, versions map[string]string) error {
	ffiVersions, err := daemonVersions(versions)

	if err != nil {
		return fmt.Errorf("error validating versions: %w", err)
	}

	cBytes := sliceToFfi(bytes)
	defer freeFfiSlice(cBytes)

	result := generated_bindings.FfiValidateAssemblyScript(generated_bindings.FfiChainTypeT(chain), cBytes, ffiVersions)

	return makeError(result)
}

func daemonParameters(parameters map[string]string) *generated_bindings.FfiDaemonParametersT {
	params := generated_bindings.FfiNewDaemonParameters()

	for key, value := range parameters {
		generated_bindings.FfiAppendDaemonParameter(params, key, value)
	}

	return params
}

func daemonVersions(versions map[string]string) (*generated_bindings.FfiDaemonVersionsT, error) {
	ffiVersions := generated_bindings.FfiNewDaemonVersions()

	for key, value := range versions {
		err := makeError(generated_bindings.FfiAppendDaemonVersion(ffiVersions, key, value))

		if err != nil {
			generated_bindings.FfiDropDaemonVersions(ffiVersions)

			return nil, err
		}
	}

	return ffiVersions, nil
}

func makeError(result generated_bindings.FfiValidationResultT) error {
	defer generated_bindings.FfiDropValidationResult(result)
	result.Deref()

	if result.IsError {
		message := C.GoString((*C.char)(unsafe.Pointer(result.Message)))

		return errors.New(message)
	} else {
		return nil
	}
}

func sliceToFfi(bytes []byte) generated_bindings.SliceRefUint8T {
	ptr := C.CBytes(bytes)

	return generated_bindings.SliceRefUint8T{
		Ptr: (*byte)(ptr),
		Len: uint64(len(bytes)),
	}
}

func freeFfiSlice(slice generated_bindings.SliceRefUint8T) {
	C.free(unsafe.Pointer(slice.Ptr))
	slice.Free()
}
