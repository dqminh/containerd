package supervisor

import "os"

type SignalTask struct {
	baseTask
	ID     string
	Pid    string
	Signal os.Signal
}

func (s *Supervisor) signal(t *SignalTask) error {
	i, ok := s.containers[t.ID]
	if !ok {
		return ErrContainerNotFound
	}
	processes, err := i.container.Processes()
	if err != nil {
		return err
	}
	for _, p := range processes {
		if p.ID() == t.Pid {
			return p.Signal(t.Signal)
		}
	}
	return ErrProcessNotFound
}
