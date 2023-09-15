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

func (r PlaybookRun) ToJson() (string, error) {
	marshal, err := json.Marshal(r)
	if err != nil {
		return "", err
	}

	return string(marshal), nil
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

func NewStepRunStatusPending() StepRunStatus {
	return StepRunStatus{
		status: StepRunStatusPending,
	}
}

func NewStepRunStatusSkipped() StepRunStatus {
	return StepRunStatus{
		status: StepRunStatusSkipped,
	}
}

func NewStepRunStatusRunning(waitingForConfirmation bool) StepRunStatus {
	return StepRunStatus{
		status: StepRunStatusRunning,
		runningData: &StepRunRunningData{
			WaitingForConfirmation: waitingForConfirmation,
		},
	}
}

func NewStepRunStatusSuccess(finishedAt time.Time, outputs []StepOutput) StepRunStatus {
	return StepRunStatus{
		status: StepRunStatusSuccess,
		successData: &StepRunSuccessData{
			FinishedAt: finishedAt,
			Outputs:    outputs,
		},
	}
}

func NewStepRunStatusFailed(failedAt time.Time) StepRunStatus {
	return StepRunStatus{
		status: StepRunStatusFailed,
		failedData: &StepRunFailedData{
			FailedAt: failedAt,
		},
	}
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

func (s *StepRunStatus) MarshalJSON() ([]byte, error) {
	var temp struct {
		Type    StepRunStatusEnum `json:"type"`
		Content interface{}       `json:"content,omitempty"`
	}

	temp.Type = s.status

	switch s.status {
	case StepRunStatusPending, StepRunStatusSkipped:
		// No extra content to marshal for these cases
	case StepRunStatusRunning:
		temp.Content = s.runningData
	case StepRunStatusSuccess:
		temp.Content = s.successData
	case StepRunStatusFailed:
		temp.Content = s.failedData
	default:
		return nil, fmt.Errorf("unknown status: %s", s.status)
	}

	return json.Marshal(temp)
}

type StepRunRunningData struct {
	WaitingForConfirmation bool `json:"waitingForConfirmation"`
}

type StepRunSuccessData struct {
	FinishedAt time.Time    `json:"finishedAt"`
	Outputs    []StepOutput `json:"outputs"`
}

type StepRunFailedData struct {
	FailedAt time.Time `json:"failedAt"`
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

func NewPlaybookRunStatusRunning() PlaybookRunStatus {
	return PlaybookRunStatus{
		status: PlaybookRunStatusRunning,
	}
}

func NewPlaybookRunStatusSuccess(finishedAt time.Time) PlaybookRunStatus {
	return PlaybookRunStatus{
		status: PlaybookRunStatusSuccess,
		successData: &PlaybookRunSuccessData{
			FinishedAt: finishedAt,
		},
	}
}

func NewPlaybookRunStatusFailed(finishedAt time.Time, stepSeq uint32, message string) PlaybookRunStatus {
	return PlaybookRunStatus{
		status: PlaybookRunStatusFailed,
		failedData: &PlaybookRunFailedData{
			FinishedAt: finishedAt,
			StepSeq:    stepSeq,
			Message:    message,
		},
	}
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

func (prs PlaybookRunStatus) MarshalJSON() ([]byte, error) {
	var temp struct {
		Type    string      `json:"type"`
		Content interface{} `json:"content,omitempty"`
	}

	switch prs.status {
	case PlaybookRunStatusSuccess:
		temp.Type = string(PlaybookRunStatusSuccess)
		temp.Content = prs.successData
	case PlaybookRunStatusFailed:
		temp.Type = string(PlaybookRunStatusFailed)
		temp.Content = prs.failedData
	case PlaybookRunStatusRunning:
		temp.Type = string(PlaybookRunStatusRunning)
	default:
		return nil, fmt.Errorf("unknown status: %s", prs.status)
	}

	return json.Marshal(temp)
}
