package tests

import (
	"github.com/Mamoru-Foundation/mamoru-core-go/playbook_engine"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestValidateUnknownTask(t *testing.T) {
	playbook := playbook_engine.Playbook{
		Task: playbook_engine.StepSteps(playbook_engine.StepsBlock{
			Steps: []playbook_engine.Step{
				playbook_engine.StepSingle(playbook_engine.SingleStep{
					Seq:    1,
					Run:    "dont_exist@1",
					Params: nil,
				}),
			},
		}),
	}

	err := playbook_engine.Validate(playbook)

	assert.EqualError(t, err, "Unknown task: dont_exist@1")
}

func TestValidateMissingParameter(t *testing.T) {
	playbook := playbook_engine.Playbook{
		Task: playbook_engine.StepSteps(playbook_engine.StepsBlock{
			Steps: []playbook_engine.Step{
				playbook_engine.StepSingle(playbook_engine.SingleStep{
					Seq: 1,
					Run: "dummy@1",
					Params: []playbook_engine.StepParam{
						{
							Name:  "hey",
							Value: "{{ incident.id }}",
						},
					},
				}),
			},
		}),
	}

	err := playbook_engine.Validate(playbook)

	assert.EqualError(t, err, `Missing required parameters for task "dummy@1": "["foo"]"`)
}

func TestValidateOk(t *testing.T) {
	playbook := playbook_engine.Playbook{
		Task: playbook_engine.StepSteps(playbook_engine.StepsBlock{
			Steps: []playbook_engine.Step{
				playbook_engine.StepSingle(playbook_engine.SingleStep{
					Seq: 1,
					Run: "dummy@1",
					Params: []playbook_engine.StepParam{
						{
							Name:  "foo",
							Value: "{{ incident.id }}",
						},
					},
				}),
			},
		}),
	}

	err := playbook_engine.Validate(playbook)

	assert.NoError(t, err)
}

func TestPlaybook(t *testing.T) {
	now, err := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	assert.NoError(t, err)

	trigger := playbook_engine.Trigger{
		Severity:   playbook_engine.SeverityAlert,
		IncidentId: "hello",
	}

	playbook := playbook_engine.Playbook{
		Task: playbook_engine.StepSteps(playbook_engine.StepsBlock{
			Steps: []playbook_engine.Step{
				playbook_engine.StepSingle(playbook_engine.SingleStep{
					Seq: 1,
					Run: "dummy@1",
					Params: []playbook_engine.StepParam{
						{
							Name:  "foo",
							Value: "{{ incident.id }}",
						},
					},
				}),
			},
		}),
	}

	result, err := playbook_engine.Start(now, playbook, trigger)

	assert.NoError(t, err)

	stepStatus := result.Run.Steps[0].Status

	assert.Equal(t, stepStatus.IsRunning(), true)
	assert.Equal(t, result.ExternalActions, []playbook_engine.ExternalAction{
		{
			StepSeq: 1,
			Action:  "dummy@1",
			Params: []playbook_engine.StepParam{
				{Name: "foo", Value: "hello"},
			},
		},
	})

	now = now.Add(time.Second * 10)

	confirmations := playbook_engine.RunConfirmations{
		Confirmations: []playbook_engine.RunConfirmation{
			{
				StepSeq: 1,
				Logs:    []string{"hello"},
				Status: playbook_engine.RunConfirmationStatusSuccess([]playbook_engine.StepOutput{
					{
						Name:  "hello",
						Value: "world",
					},
				}),
			},
		},
	}

	result, err = playbook_engine.Resume(now, playbook, trigger, result.Run, confirmations)

	assert.NoError(t, err)
	assert.Equal(t, result.ExternalActions, []playbook_engine.ExternalAction{})
	assert.Equal(t, result.Run.Status.IsSuccess(), true)

	stepStatus = result.Run.Steps[0].Status

	assert.Equal(t, stepStatus.IsSuccess(), true)
	assert.Equal(t, stepStatus.GetSuccess(), &playbook_engine.StepRunSuccessData{
		FinishedAt: now,
		Outputs: []playbook_engine.StepOutput{
			{
				Name:  "hello",
				Value: "world",
			},
		},
	})
}
