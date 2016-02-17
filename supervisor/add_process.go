package supervisor

import (
	"time"

	"github.com/docker/containerd/runtime"
	"github.com/opencontainers/specs"
)

type AddProcessTask struct {
	baseTask
	ID            string
	Pid           string
	Stdout        string
	Stderr        string
	Stdin         string
	ProcessSpec   *specs.Process
	StartResponse chan StartResponse
}

func (s *Supervisor) addProcess(t *AddProcessTask) error {
	start := time.Now()
	ci, ok := s.containers[t.ID]
	if !ok {
		return ErrContainerNotFound
	}
	process, err := ci.container.Exec(t.Pid, *t.ProcessSpec, runtime.NewStdio(t.Stdin, t.Stdout, t.Stderr))
	if err != nil {
		return err
	}
	if err := s.monitorProcess(process); err != nil {
		return err
	}
	ExecProcessTimer.UpdateSince(start)
	t.StartResponse <- StartResponse{}
	s.notifySubscribers(Event{
		Timestamp: time.Now(),
		Type:      "start-process",
		Pid:       t.Pid,
		ID:        t.ID,
	})
	return nil
}
