package python

import (
	"os"
)

type Workspace struct {
	workDir   string
	scriptDir string
	libDir    string
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
	if err := os.MkdirAll(w.workDir, 0o755); err != nil {
		return err
	}
	return nil
}

func (w *Workspace) GetWorkDir() string {
	return w.workDir
}
