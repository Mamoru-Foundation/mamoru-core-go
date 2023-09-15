package playbook_engine

/*
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"github.com/Mamoru-Foundation/mamoru-core-go/generated_bindings"
	"time"
	"unsafe"
)

func Start(now time.Time, playbook Playbook, trigger Trigger) (*EngineResponse, error) {
	nowRfc3339 := formatTime(now)

	triggerJson, err := trigger.ToJson()
	if err != nil {
		return nil, err
	}

	playbookJson, err := playbook.ToJson()
	if err != nil {
		return nil, err
	}

	result := generated_bindings.FfiPlaybookStart(nowRfc3339, playbookJson, triggerJson)
	resultJson, err := unwrapError(result)
	if err != nil {
		return nil, err
	}

	response, err := EngineResponseFromJson(*resultJson)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func Resume(now time.Time, playbook Playbook, trigger Trigger, run PlaybookRun, confirmations RunConfirmations) (*EngineResponse, error) {
	nowRfc3339 := formatTime(now)

	playbookJson, err := playbook.ToJson()
	if err != nil {
		return nil, err
	}

	triggerJson, err := trigger.ToJson()
	if err != nil {
		return nil, err
	}

	runJson, err := run.ToJson()
	if err != nil {
		return nil, err
	}

	confirmationsJson, err := confirmations.ToJson()
	if err != nil {
		return nil, err
	}

	result := generated_bindings.FfiPlaybookResume(nowRfc3339, playbookJson, triggerJson, runJson, confirmationsJson)

	resultJson, err := unwrapError(result)
	if err != nil {
		return nil, err
	}

	response, err := EngineResponseFromJson(*resultJson)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func Validate(playbook Playbook) error {
	playbookJson, err := playbook.ToJson()
	if err != nil {
		return err
	}

	result := generated_bindings.FfiValidatePlaybook(playbookJson)

	_, err = unwrapError(result)

	return err
}

func unwrapError(result generated_bindings.FfiJsonResultT) (*string, error) {
	result.Deref()
	defer generated_bindings.FfiDropJsonResult(result)

	if result.IsError {
		message := C.GoString((*C.char)(unsafe.Pointer(result.ErrorMessage)))

		return nil, errors.New(message)
	} else {
		result := C.GoString((*C.char)(unsafe.Pointer(result.ResultJson)))

		return &result, nil
	}
}
