package python

import (
	"os"
	"path/filepath"
)

type Workspace struct {
	workDir   string
	scriptDir string
	LibDir    string
}

func NewWorkspace(workDir string) *Workspace {
	w := &Workspace{
		workDir: workDir,
	}
	err := w.setup()
	if err != nil {
		return nil
	}
	return w
}

func (w *Workspace) setup() error {

	baseDir := filepath.Join(w.workDir, "python")

	// create script directory
	scriptDir := filepath.Join(baseDir, "scripts")
	if err := os.MkdirAll(scriptDir, 0o755); err != nil {
		return err
	}
	w.scriptDir = scriptDir
	// create script directory
	libDir := filepath.Join(baseDir, "lib")
	if err := os.MkdirAll(libDir, 0o755); err != nil {
		return err
	}
	w.LibDir = libDir
	return nil
}

func (w *Workspace) GetWorkDir() string {
	return w.workDir
}
