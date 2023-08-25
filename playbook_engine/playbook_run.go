package playbook_engine

import (
	"encoding/json"
	"fmt"
	"time"
)

type EngineResponse struct {
	ExternalActions []ExternalAction `json:"externalActions"`
	Run             PlaybookRun      `json:"run"`
}

type ExternalAction struct {
	StepSeq uint32      `json:"stepSeq"`
	Action  string      `json:"action"`
	Params  []StepParam `json:"params,omitempty"`
}

type PlaybookRun struct {
	Status PlaybookRunStatus `json:"status"`
	Steps  []StepRun         `json:"steps"`
}

type StepRun struct {
	StepSeq   uint32        `json:"stepSeq"`
	StartedAt *time.Time    `json:"startedAt,omitempty"`
	Status    StepRunStatus `json:"status"`
	Logs      []string      `json:"logs"`
}

type StepRunStatusEnum string

const (
	StepRunStatusPending StepRunStatusEnum = "pending"
	StepRunStatusSkipped StepRunStatusEnum = "skipped"
	StepRunStatusRunning StepRunStatusEnum = "running"
	StepRunStatusSuccess StepRunStatusEnum = "success"
	StepRunStatusFailed  StepRunStatusEnum = "failed"
)

type StepRunStatus struct {
	status      StepRunStatusEnum
	runningData *StepRunRunningData
	successData *StepRunSuccessData
	failedData  *StepRunFailedData
}

func (s *StepRunStatus) IsPending() bool {
	return s.status == StepRunStatusPending
}

func (s *StepRunStatus) IsSkipped() bool {
	return s.status == StepRunStatusSkipped
}

func (s *StepRunStatus) IsRunning() bool {
	return s.status == StepRunStatusRunning
}

func (s *StepRunStatus) IsFailed() bool {
	return s.status == StepRunStatusFailed
}

func (s *StepRunStatus) IsSuccess() bool {
	return s.status == StepRunStatusSuccess
}

func (s *StepRunStatus) GetRunning() *StepRunRunningData {
	if s.IsRunning() {
		return s.runningData
	}

	return nil
}

func (s *StepRunStatus) GetSuccess() *StepRunSuccessData {
	if s.IsSuccess() {
		return s.successData
	}

	return nil
}

func (s *StepRunStatus) GetFailed() *StepRunFailedData {
	if s.IsFailed() {
		return s.failedData
	}

	return nil
}

func EngineResponseFromJson(jsonString string) (EngineResponse, error) {
	var engineResponse EngineResponse
	err := json.Unmarshal([]byte(jsonString), &engineResponse)
	if err != nil {
		return EngineResponse{}, err
	}

	return engineResponse, nil
}

func (s *StepRunStatus) UnmarshalJSON(data []byte) error {
	var temp struct {
		Type    StepRunStatusEnum `json:"type"`
		Content json.RawMessage   `json:"content,omitempty"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	switch temp.Type {
	case StepRunStatusPending:
		s.status = temp.Type
	case StepRunStatusSkipped:
		s.status = temp.Type
	case StepRunStatusRunning:
		var runningData StepRunRunningData
		if err := json.Unmarshal(temp.Content, &runningData); err != nil {
			return err
		}
		s.status = temp.Type
		s.runningData = &runningData
	case StepRunStatusSuccess:
		var successData StepRunSuccessData
		if err := json.Unmarshal(temp.Content, &successData); err != nil {
			return err
		}
		s.status = temp.Type
		s.successData = &successData
	case StepRunStatusFailed:
		var failedData StepRunFailedData
		if err := json.Unmarshal(temp.Content, &failedData); err != nil {
			return err
		}
		s.status = temp.Type
		s.failedData = &failedData
	default:
		return fmt.Errorf("unknown type: %s", temp.Type)
	}

	return nil
}

type StepRunRunningData struct {
	WaitingForConfirmation bool `json:"waitingForConfirmation"`
}

type StepRunSuccessData struct {
	FinishedAt time.Time    `json:"finishedAt"`
	Outputs    []StepOutput `json:"outputs"`
}

type StepRunFailedData struct {
	FailedAt time.Time `json:"finishedAt"`
}

type PlaybookRunStatusEnum string

const (
	PlaybookRunStatusRunning PlaybookRunStatusEnum = "running"
	PlaybookRunStatusSuccess PlaybookRunStatusEnum = "success"
	PlaybookRunStatusFailed  PlaybookRunStatusEnum = "failed"
)

type PlaybookRunSuccessData struct {
	FinishedAt time.Time `json:"finishedAt"`
}

type PlaybookRunFailedData struct {
	FinishedAt time.Time `json:"finishedAt"`
	StepSeq    uint32    `json:"stepSeq"`
	Message    string    `json:"message"`
}

type PlaybookRunStatus struct {
	status      PlaybookRunStatusEnum
	successData *PlaybookRunSuccessData
	failedData  *PlaybookRunFailedData
}

type StepParam struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (prs *PlaybookRunStatus) IsRunning() bool {
	return prs.status == PlaybookRunStatusRunning
}

func (prs *PlaybookRunStatus) IsFailed() bool {
	return prs.status == PlaybookRunStatusFailed
}

func (prs *PlaybookRunStatus) IsSuccess() bool {
	return prs.status == PlaybookRunStatusSuccess
}

func (prs *PlaybookRunStatus) GetSuccess() *PlaybookRunSuccessData {
	if prs.IsSuccess() {
		return prs.successData
	}

	return nil
}

func (prs *PlaybookRunStatus) GetFailed() *PlaybookRunFailedData {
	if prs.IsFailed() {
		return prs.failedData
	}

	return nil
}

func (prs *PlaybookRunStatus) UnmarshalJSON(data []byte) error {
	var temp struct {
		Type    string          `json:"type"`
		Content json.RawMessage `json:"content,omitempty"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	switch temp.Type {
	case "success":
		var successData PlaybookRunSuccessData
		if err := json.Unmarshal(temp.Content, &successData); err != nil {
			return err
		}
		prs.status = PlaybookRunStatusSuccess
		prs.successData = &successData
	case "failed":
		var failedData PlaybookRunFailedData
		if err := json.Unmarshal(temp.Content, &failedData); err != nil {
			return err
		}
		prs.status = PlaybookRunStatusFailed
		prs.failedData = &failedData
	case "running":
		prs.status = PlaybookRunStatusRunning
	default:
		return fmt.Errorf("unknown type: %s", temp.Type)
	}

	return nil
}
