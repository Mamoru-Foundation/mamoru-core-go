// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Mon, 22 Apr 2024 19:42:01 WEST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package generated_bindings

/*
#cgo CFLAGS: -I${SRCDIR}/../packaged/include
#cgo LDFLAGS: -lmamoru_core_go
#cgo darwin,arm64 LDFLAGS: -L${SRCDIR}/../packaged/lib/darwin-arm64
#cgo darwin,amd64 LDFLAGS: -L${SRCDIR}/../packaged/lib/darwin-amd64
#cgo linux,amd64 LDFLAGS: -Wl,--no-as-needed -ldl -lm -L${SRCDIR}/../packaged/lib/linux-amd64
#include <libmamoru_query_validator_go.h>
#include <libmamoru_playbook_engine_go.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

// FfiDropJsonResult function as declared in include/libmamoru_playbook_engine_go.h:35
func FfiDropJsonResult(Result FfiJsonResultT) {
	cResult, cResultAllocMap := Result.PassValue()
	C.ffi_drop_json_result(cResult)
	runtime.KeepAlive(cResultAllocMap)
}

// FfiPlaybookResume function as declared in include/libmamoru_playbook_engine_go.h:44
func FfiPlaybookResume(Datetime string, PlaybookJson string, TriggerJson string, RunJson string, ConfirmationsJson string) FfiJsonResultT {
	Datetime = safeString(Datetime)
	cDatetime, cDatetimeAllocMap := unpackPCharString(Datetime)
	PlaybookJson = safeString(PlaybookJson)
	cPlaybookJson, cPlaybookJsonAllocMap := unpackPCharString(PlaybookJson)
	TriggerJson = safeString(TriggerJson)
	cTriggerJson, cTriggerJsonAllocMap := unpackPCharString(TriggerJson)
	RunJson = safeString(RunJson)
	cRunJson, cRunJsonAllocMap := unpackPCharString(RunJson)
	ConfirmationsJson = safeString(ConfirmationsJson)
	cConfirmationsJson, cConfirmationsJsonAllocMap := unpackPCharString(ConfirmationsJson)
	__ret := C.ffi_playbook_resume(cDatetime, cPlaybookJson, cTriggerJson, cRunJson, cConfirmationsJson)
	runtime.KeepAlive(ConfirmationsJson)
	runtime.KeepAlive(cConfirmationsJsonAllocMap)
	runtime.KeepAlive(RunJson)
	runtime.KeepAlive(cRunJsonAllocMap)
	runtime.KeepAlive(TriggerJson)
	runtime.KeepAlive(cTriggerJsonAllocMap)
	runtime.KeepAlive(PlaybookJson)
	runtime.KeepAlive(cPlaybookJsonAllocMap)
	runtime.KeepAlive(Datetime)
	runtime.KeepAlive(cDatetimeAllocMap)
	__v := *NewFfiJsonResultTRef(unsafe.Pointer(&__ret))
	return __v
}

// FfiPlaybookStart function as declared in include/libmamoru_playbook_engine_go.h:57
func FfiPlaybookStart(Datetime string, PlaybookJson string, TriggerJson string) FfiJsonResultT {
	Datetime = safeString(Datetime)
	cDatetime, cDatetimeAllocMap := unpackPCharString(Datetime)
	PlaybookJson = safeString(PlaybookJson)
	cPlaybookJson, cPlaybookJsonAllocMap := unpackPCharString(PlaybookJson)
	TriggerJson = safeString(TriggerJson)
	cTriggerJson, cTriggerJsonAllocMap := unpackPCharString(TriggerJson)
	__ret := C.ffi_playbook_start(cDatetime, cPlaybookJson, cTriggerJson)
	runtime.KeepAlive(TriggerJson)
	runtime.KeepAlive(cTriggerJsonAllocMap)
	runtime.KeepAlive(PlaybookJson)
	runtime.KeepAlive(cPlaybookJsonAllocMap)
	runtime.KeepAlive(Datetime)
	runtime.KeepAlive(cDatetimeAllocMap)
	__v := *NewFfiJsonResultTRef(unsafe.Pointer(&__ret))
	return __v
}

// FfiValidatePlaybook function as declared in include/libmamoru_playbook_engine_go.h:66
func FfiValidatePlaybook(PlaybookJson string) FfiJsonResultT {
	PlaybookJson = safeString(PlaybookJson)
	cPlaybookJson, cPlaybookJsonAllocMap := unpackPCharString(PlaybookJson)
	__ret := C.ffi_validate_playbook(cPlaybookJson)
	runtime.KeepAlive(PlaybookJson)
	runtime.KeepAlive(cPlaybookJsonAllocMap)
	__v := *NewFfiJsonResultTRef(unsafe.Pointer(&__ret))
	return __v
}

// FfiValidateSql function as declared in include/libmamoru_query_validator_go.h:60
func FfiValidateSql(Chain FfiChainTypeT, Query string, Parameters *FfiDaemonParametersT, Versions *FfiDaemonVersionsT) FfiValidationResultT {
	cChain, cChainAllocMap := (C.FfiChainType_t)(Chain), cgoAllocsUnknown
	Query = safeString(Query)
	cQuery, cQueryAllocMap := unpackPCharString(Query)
	cParameters, cParametersAllocMap := (*C.FfiDaemonParameters_t)(unsafe.Pointer(Parameters)), cgoAllocsUnknown
	cVersions, cVersionsAllocMap := (*C.FfiDaemonVersions_t)(unsafe.Pointer(Versions)), cgoAllocsUnknown
	__ret := C.ffi_validate_sql(cChain, cQuery, cParameters, cVersions)
	runtime.KeepAlive(cVersionsAllocMap)
	runtime.KeepAlive(cParametersAllocMap)
	runtime.KeepAlive(Query)
	runtime.KeepAlive(cQueryAllocMap)
	runtime.KeepAlive(cChainAllocMap)
	__v := *NewFfiValidationResultTRef(unsafe.Pointer(&__ret))
	return __v
}

// FfiValidateSqlRenders function as declared in include/libmamoru_query_validator_go.h:69
func FfiValidateSqlRenders(Query string, Parameters *FfiDaemonParametersT, Versions *FfiDaemonVersionsT) FfiValidationResultT {
	Query = safeString(Query)
	cQuery, cQueryAllocMap := unpackPCharString(Query)
	cParameters, cParametersAllocMap := (*C.FfiDaemonParameters_t)(unsafe.Pointer(Parameters)), cgoAllocsUnknown
	cVersions, cVersionsAllocMap := (*C.FfiDaemonVersions_t)(unsafe.Pointer(Versions)), cgoAllocsUnknown
	__ret := C.ffi_validate_sql_renders(cQuery, cParameters, cVersions)
	runtime.KeepAlive(cVersionsAllocMap)
	runtime.KeepAlive(cParametersAllocMap)
	runtime.KeepAlive(Query)
	runtime.KeepAlive(cQueryAllocMap)
	__v := *NewFfiValidationResultTRef(unsafe.Pointer(&__ret))
	return __v
}

// FfiDropValidationResult function as declared in include/libmamoru_query_validator_go.h:74
func FfiDropValidationResult(Result FfiValidationResultT) {
	cResult, cResultAllocMap := Result.PassValue()
	C.ffi_drop_validation_result(cResult)
	runtime.KeepAlive(cResultAllocMap)
}

// FfiNewDaemonParameters function as declared in include/libmamoru_query_validator_go.h:77
func FfiNewDaemonParameters() *FfiDaemonParametersT {
	__ret := C.ffi_new_daemon_parameters()
	__v := *(**FfiDaemonParametersT)(unsafe.Pointer(&__ret))
	return __v
}

// FfiAppendDaemonParameter function as declared in include/libmamoru_query_validator_go.h:79
func FfiAppendDaemonParameter(Parameters *FfiDaemonParametersT, Key string, Value string) {
	cParameters, cParametersAllocMap := (*C.FfiDaemonParameters_t)(unsafe.Pointer(Parameters)), cgoAllocsUnknown
	Key = safeString(Key)
	cKey, cKeyAllocMap := unpackPCharString(Key)
	Value = safeString(Value)
	cValue, cValueAllocMap := unpackPCharString(Value)
	C.ffi_append_daemon_parameter(cParameters, cKey, cValue)
	runtime.KeepAlive(Value)
	runtime.KeepAlive(cValueAllocMap)
	runtime.KeepAlive(Key)
	runtime.KeepAlive(cKeyAllocMap)
	runtime.KeepAlive(cParametersAllocMap)
}

// FfiNewDaemonVersions function as declared in include/libmamoru_query_validator_go.h:84
func FfiNewDaemonVersions() *FfiDaemonVersionsT {
	__ret := C.ffi_new_daemon_versions()
	__v := *(**FfiDaemonVersionsT)(unsafe.Pointer(&__ret))
	return __v
}

// FfiAppendDaemonVersion function as declared in include/libmamoru_query_validator_go.h:86
func FfiAppendDaemonVersion(Versions *FfiDaemonVersionsT, Key string, Value string) FfiValidationResultT {
	cVersions, cVersionsAllocMap := (*C.FfiDaemonVersions_t)(unsafe.Pointer(Versions)), cgoAllocsUnknown
	Key = safeString(Key)
	cKey, cKeyAllocMap := unpackPCharString(Key)
	Value = safeString(Value)
	cValue, cValueAllocMap := unpackPCharString(Value)
	__ret := C.ffi_append_daemon_version(cVersions, cKey, cValue)
	runtime.KeepAlive(Value)
	runtime.KeepAlive(cValueAllocMap)
	runtime.KeepAlive(Key)
	runtime.KeepAlive(cKeyAllocMap)
	runtime.KeepAlive(cVersionsAllocMap)
	__v := *NewFfiValidationResultTRef(unsafe.Pointer(&__ret))
	return __v
}

// FfiDropDaemonVersions function as declared in include/libmamoru_query_validator_go.h:91
func FfiDropDaemonVersions(Versions *FfiDaemonVersionsT) {
	cVersions, cVersionsAllocMap := (*C.FfiDaemonVersions_t)(unsafe.Pointer(Versions)), cgoAllocsUnknown
	C.ffi_drop_daemon_versions(cVersions)
	runtime.KeepAlive(cVersionsAllocMap)
}
