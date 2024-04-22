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

const (
	// FfiChainTypeSui as declared in include/libmamoru_query_validator_go.h:29
	FfiChainTypeSui = iota
	// FfiChainTypeEvm as declared in include/libmamoru_query_validator_go.h:31
	FfiChainTypeEvm = 1
	// FfiChainTypeAptos as declared in include/libmamoru_query_validator_go.h:33
	FfiChainTypeAptos = 2
	// FfiChainTypeCosmos as declared in include/libmamoru_query_validator_go.h:35
	FfiChainTypeCosmos = 3
)
