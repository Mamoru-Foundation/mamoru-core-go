package playbook_engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// tests format time
func TestFormatTime(t *testing.T) {
	now, err := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	assert.NoError(t, err)

	nowRfc3339 := formatTime(now)
	assert.Equal(t, now.Format(time.RFC3339), nowRfc3339)
}

func TestTrigger_ToJson(t *testing.T) {
	trigger := Trigger{
		Severity:   SeverityInfo,
		IncidentId: "hello",
	}

	json, err := trigger.ToJson()
	assert.NoError(t, err)
	assert.Equal(t, `{"severity":"info","incidentId":"hello"}`, json)
}

func TestRunConfirmations_ToJson(t *testing.T) {
	runConfirmations := RunConfirmations{
		Confirmations: []RunConfirmation{
			{
				StepSeq: 1,
				Logs:    []string{"hello"},
				Status: RunConfirmationStatus{
					success: true,
					outputs: []StepOutput{
						{
							Name:  "hello",
							Value: "world",
						},
					},
				},
			},
			{
				StepSeq: 2,
				Logs:    []string{"hello"},
				Status: RunConfirmationStatus{
					success: false,
					outputs: nil,
				},
			},
		},
	}

	json, err := runConfirmations.ToJson()
	assert.NoError(t, err)

	assert.Equal(t, `[{"logs":["hello"],"status":{"content":{"outputs":[{"name":"hello","value":"world"}]},"type":"success"},"stepSeq":1},{"logs":["hello"],"status":{"type":"failed"},"stepSeq":2}]`, json)
}

func TestPlaybookRun_ToJson(t *testing.T) {
	now, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")

	run := PlaybookRun{
		Status: NewPlaybookRunStatusFailed(now, 1, "hello"),
		Steps: []StepRun{
			{
				StepSeq:   0,
				StartedAt: &now,
				Status:    NewStepRunStatusFailed(now),
				Logs:      nil,
			},
			{
				StepSeq:   1,
				StartedAt: nil,
				Status:    NewStepRunStatusPending(),
				Logs:      []string{"hello"},
			},
		},
	}

	json, err := run.ToJson()
	assert.NoError(t, err)

	assert.Equal(t, `{"status":{"type":"failed","content":{"finishedAt":"2021-01-01T00:00:00Z","stepSeq":1,"message":"hello"}},"steps":[{"stepSeq":0,"startedAt":"2021-01-01T00:00:00Z","status":{"type":"failed","content":{"failedAt":"2021-01-01T00:00:00Z"}},"logs":null},{"stepSeq":1,"status":{"type":"pending"},"logs":["hello"]}]}`, json)
}
