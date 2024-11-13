package tracesmith

import "github.com/google/uuid"

type Chain struct {
	Name   string
	Runs   []*Run
	client *Client
}

func NewChain(client *Client, name string) *Chain {
	return &Chain{
		Name:   name,
		client: client,
	}
}

func (c *Chain) AddRun(name, runType string, inputs map[string]interface{}, parentID *uuid.UUID) (*Run, error) {
	run := NewRun(c.client, name, runType, inputs, parentID)
	if err := run.Start(); err != nil {
		return nil, err
	}
	c.Runs = append(c.Runs, run)
	return run, nil
}

func (c *Chain) EndAllRuns(outputs map[string]interface{}) error {
	for _, run := range c.Runs {
		if run.EndTime.IsZero() { // Ensure we only end runs that haven’t been ended yet
			if err := run.End(outputs); err != nil {
				return err
			}
		}
	}
	return nil
}
