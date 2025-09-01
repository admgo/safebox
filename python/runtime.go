package python

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/google/uuid"
)

type PythonRuntime struct {
	config    *PythonRuntimeConfig
	capturer  *Capturer
	workspace *Workspace
}

type PytemplateData struct {
	Code string
}

func (r *PytemplateData) Indent(spaces int, v string) string {
	pad := strings.Repeat(" ", spaces)
	lines := strings.Split(v, "\n")
	for i, line := range lines {
		// do not indent in first line
		if line != "" && i > 0 {
			lines[i] = pad + line
		}
	}
	return strings.Join(lines, "\n")
}

func NewPythonRuntime(w *Workspace) *PythonRuntime {
	r := &PythonRuntime{
		config: &PythonRuntimeConfig{
			MaxWorkers:         10,
			MaxRequests:        100,
			WorkerTimeout:      10,
			pythonPath:         "python3",
			PythonPipMirrorURL: "https://pypi.tuna.tsinghua.edu.cn/simple",
			PythonLibPaths:     []string{},
			PipMirrorURL:       "https://pypi.tuna.tsinghua.edu.cn/simple",
		},
		capturer:  NewOutputCapturer(),
		workspace: w,
	}
	err := r.initialize()
	if err != nil {
		return nil
	}
	return r
}

func (r *PythonRuntime) initialize() error {
	runtimeDir := filepath.Join(r.workspace.workDir, "python")

	scriptDir := filepath.Join(runtimeDir, "scripts")
	if err := os.MkdirAll(scriptDir, 0o755); err != nil {
		return err
	}

	cmd := exec.Command(r.config.pythonPath, "load_standard_library.py", runtimeDir)
	// Connect stdout/stderr to this process, so errors show up properly
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run synchronously
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func (r *PythonRuntime) dump(code string) (string, error) {
	data := &PytemplateData{
		Code: code,
	}

	// read template file
	tmpl, err := template.ParseFiles("template/python.tmpl")
	if err != nil {
		return "", err
	}

	// generate a python script
	filename := uuid.New().String() + ".py"
	fullname := r.workspace.scriptDir + "/" + filename
	file, err := os.Create(fullname)
	if err != nil {
		return "", err
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		return "", err
	}
	return fullname, nil
}

func (r *PythonRuntime) Run(code string) (chan []byte, chan []byte, chan bool, error) {
	fullname, err := r.dump(code)
	if err != nil {
		return nil, nil, nil, err
	}

	cmd := exec.Command(r.config.pythonPath, fullname)
	cmd.Dir = r.workspace.workDir

	err = r.capturer.CaptureOutput(cmd)
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
