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
