package tracesmith

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Run struct {
	ID          uuid.UUID
	Name        string
	RunType     string
	StartTime   time.Time
	EndTime     time.Time
	Inputs      map[string]interface{}
	Outputs     map[string]interface{}
	ParentID    *uuid.UUID
	SessionName string
	Tags        []string
	Extra       map[string]interface{}
	client      *Client
}

func NewRun(client *Client, name, runType string, inputs map[string]interface{}, parentID *uuid.UUID) *Run {
	sessionName := os.Getenv("LANGSMITH_SESSION_NAME")
	if sessionName == "" {
		sessionName = "default" // Default if not set
	}

	tags := os.Getenv("LANGSMITH_TAGS")
	if tags == "" {
		tags = "langsmith,rest,my-trace" // Default tags if not set
	}
	tagsList := strings.Split(tags, ",")

	extraMetadataKey := os.Getenv("LANGSMITH_METADATA_KEY")
	if extraMetadataKey == "" {
		extraMetadataKey = "my_key"
	}

	extraMetadataValue := os.Getenv("LANGSMITH_METADATA_VALUE")
	if extraMetadataValue == "" {
		extraMetadataValue = "My value"
	}

	return &Run{
		ID:          uuid.New(),
		Name:        name,
		RunType:     runType,
		StartTime:   time.Now().UTC(),
		Inputs:      inputs,
		ParentID:    parentID,
		SessionName: sessionName,
		Tags:        tagsList,
		Extra: map[string]interface{}{
			"metadata": map[string]string{
				extraMetadataKey: extraMetadataValue,
			},
			"runtime": map[string]string{
				"platform": runtime.GOOS,
			},
		},
		client: client,
	}
}

func (r *Run) Start() error {
	data := r.toMap()
	return r.client.sendRequest("POST", "/runs", data)
}

func (r *Run) End(outputs map[string]interface{}) error {
	r.EndTime = time.Now().UTC()
	r.Outputs = outputs
	data := r.toMap()
	return r.client.sendRequest("PATCH", fmt.Sprintf("/runs/%s", r.ID.String()), data)
}

func (r *Run) toMap() map[string]interface{} {
	data := map[string]interface{}{
		"id":           r.ID.String(),
		"name":         r.Name,
		"run_type":     r.RunType,
		"start_time":   r.StartTime.Format(time.RFC3339),
		"inputs":       r.Inputs,
		"session_name": r.SessionName,
		"tags":         r.Tags,
		"extra":        r.Extra,
	}
	if !r.EndTime.IsZero() {
		data["end_time"] = r.EndTime.Format(time.RFC3339)
		data["outputs"] = r.Outputs
	}
	if r.ParentID != nil {
		data["parent_run_id"] = r.ParentID.String()
	}
	return data
}
