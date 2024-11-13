# TraceSmith

**TraceSmith** is a Go implementation of the LangSmith tracing REST API. This project provides a straightforward wrapper for interacting with the LangSmith API, enabling developers to integrate tracing into Go applications with ease.

For more details on the LangSmith API, refer to the official documentation: [LangSmith Tracing REST API](https://docs.smith.langchain.com/old/cookbook/tracing-examples/rest).

## Overview

TraceSmith simplifies the process of:
- Starting and ending trace runs.
- Organizing multiple runs within chains.
- Adding custom metadata, tags, and session names for improved trace management.

## Installation

To add TraceSmith to your project, run:

```bash
go get github.com/gelhteag/tracesmith
```

## Usage

### 1. Set up Environment Variables

TraceSmith relies on the following environment variables to configure the connection and metadata:

- **Required**:
  - `LANGCHAIN_ENDPOINT`: The endpoint for the LangSmith tracing API.
  - `LANGCHAIN_API_KEY`: Your LangSmith API key for authentication.
  - `LANGCHAIN_PROJECT`: The project or session name for grouping traces.

- **Optional**:
  - `LANGSMITH_SESSION_NAME`: Custom session name to organize and group traces.
  - `LANGSMITH_TAGS`: Comma-separated tags (e.g., `"tag1,tag2,tag3"`) to label trace runs for easy identification.
  - `LANGSMITH_METADATA_KEY`: Key for additional metadata in traces.
  - `LANGSMITH_METADATA_VALUE`: Value for the metadata key, allowing you to add custom details to trace runs.

Set these environment variables in your shell or `.env` file, for example:

```bash
export LANGCHAIN_ENDPOINT="https://api.smith.langchain.com"
export LANGCHAIN_API_KEY="your-api-key"
export LANGCHAIN_PROJECT="your-project-name"
export LANGSMITH_SESSION_NAME="session_name"
export LANGSMITH_TAGS="tag1,tag2,tag3"
export LANGSMITH_METADATA_KEY="custom_key"
export LANGSMITH_METADATA_VALUE="custom_value"
```

### 2. Using TraceSmith in Your Code

```go
package main

import (
    "fmt"
    "github.com/yourusername/tracesmith/pkg/tracesmith"
    "github.com/sirupsen/logrus"
)

func main() {
    client := tracesmith.NewClient()
    chain := tracesmith.NewChain(client, "ExampleChain")

    inputs := map[string]interface{}{"query": "example input"}
    run, err := chain.AddRun("ExampleRun", "chain", inputs, nil)
    if err != nil {
        logrus.WithError(err).Fatal("Error starting run")
    }

    fmt.Println("Processing...")

    outputs := map[string]interface{}{"result": "example output"}
    if err := run.End(outputs); err != nil {
        logrus.WithError(err).Fatal("Error ending run")
    }

    if err := chain.EndAllRuns(outputs); err != nil {
        logrus.WithError(err).Fatal("Error ending all runs in chain")
    }

    logrus.Infoln("Run completed successfully")
}
```

## License

This project is licensed under the MIT License.
