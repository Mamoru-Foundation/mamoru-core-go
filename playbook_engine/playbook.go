package playbook_engine

import (
	"encoding/json"
	"time"
)

type Severity string

const (
	SeverityInfo    Severity = "info"
	SeverityWarning Severity = "warning"
	SeverityError   Severity = "error"
	SeverityAlert   Severity = "alert"
)

type Trigger struct {
	Severity   Severity `json:"severity"`
	IncidentId string   `json:"incidentId"`
}

type RunConfirmations struct {
	Confirmations []RunConfirmation
}

type RunConfirmation struct {
	Logs    []string              `json:"logs"`
	Status  RunConfirmationStatus `json:"status"`
	StepSeq uint32                `json:"stepSeq"`
}

type RunConfirmationStatus struct {
	success bool
	outputs []StepOutput
}

func RunConfirmationStatusSuccess(outputs []StepOutput) RunConfirmationStatus {
	return RunConfirmationStatus{
		success: true,
		outputs: outputs,
	}
}

func RunConfirmationStatusFailed() RunConfirmationStatus {
	return RunConfirmationStatus{
		success: false,
		outputs: nil,
	}
}

type StepOutput struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Playbook struct {
	Task Step `json:"task"`
}

type Step struct {
	Single   *SingleStep    `json:"single,omitempty"`
	Steps    *StepsBlock    `json:"steps,omitempty"`
	Parallel *ParallelBlock `json:"parallel,omitempty"`
}

func StepSingle(step SingleStep) Step {
	return Step{
		Single:   &step,
		Steps:    nil,
		Parallel: nil,
	}
}

func StepSteps(steps StepsBlock) Step {
	return Step{
		Single:   nil,
		Steps:    &steps,
		Parallel: nil,
	}
}

func StepParallel(steps ParallelBlock) Step {
	return Step{
		Single:   nil,
		Steps:    nil,
		Parallel: &steps,
	}
}

type StepsBlock struct {
	Condition string `json:"condition,omitempty"`
	Steps     []Step `json:"steps"`
}

type ParallelBlock struct {
	Condition string `json:"condition,omitempty"`
	Steps     []Step `json:"steps"`
}

type SingleStep struct {
	Seq       uint32      `json:"seq"`
	Id        string      `json:"id,omitempty"`
	Condition string      `json:"condition,omitempty"`
	Run       string      `json:"run"`
	Params    []StepParam `json:"params,omitempty"`
}

func (p Playbook) ToJson() (string, error) {
	marshal, err := json.Marshal(p)
	if err != nil {
		return "", err
	}

	return string(marshal), nil
}

func (t Trigger) ToJson() (string, error) {
	marshal, err := json.Marshal(t)
	if err != nil {
		return "", err
	}

	return string(marshal), nil
}

func (rcs RunConfirmationStatus) MarshalJSON() ([]byte, error) {
	if rcs.success {
		return json.Marshal(map[string]interface{}{
			"type": "success",
			"content": struct {
				Outputs []StepOutput `json:"outputs"`
			}{
				Outputs: rcs.outputs,
			},
		})
	}

	return json.Marshal(map[string]string{"type": "failed"})
}

func (r RunConfirmations) ToJson() (string, error) {
	marshal, err := json.Marshal(r.Confirmations)
	if err != nil {
		return "", err
	}

	return string(marshal), nil
}

func formatTime(date time.Time) string {
	return date.UTC().Format(time.RFC3339)
}
