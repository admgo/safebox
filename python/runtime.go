package python

import (
	"context"
	"os/exec"
)

type PythonRuntime struct {
	config   *PythonRuntimeConfig
	capturer *Capturer
}

func NewPythonRuntime() *PythonRuntime {
	return &PythonRuntime{
		config: &PythonRuntimeConfig{
			MaxWorkers:         10,
			MaxRequests:        100,
			WorkerTimeout:      10,
			pythonPath:         "python3",
			PythonPipMirrorURL: "https://pypi.tuna.tsinghua.edu.cn/simple",
			PythonLibPaths:     []string{},
			PipMirrorURL:       "https://pypi.tuna.tsinghua.edu.cn/simple",
		},
	}
}

func (r *PythonRuntime) Name() string {
	return "python"
}

func (r *PythonRuntime) Version() string {
	return ""
}

func (r *PythonRuntime) Run(code string) (chan []byte, chan []byte, chan bool, error) {
	cmd := exec.CommandContext(context.Background(), "go", "test", "-run", "TestLongTimeOutput")
	err := r.capturer.CaptureOutput(cmd)
	if err != nil {
		return nil, nil, nil, err
	}
	return r.capturer.GetStdout(), r.capturer.GetStderr(), r.capturer.GetDone(), nil
}

type PythonRuntimeConfig struct {
	MaxWorkers         uint16
	MaxRequests        uint16
	WorkerTimeout      uint16
	pythonPath         string
	PythonPipMirrorURL string
	PythonLibPaths     []string
	PipMirrorURL       string
}
