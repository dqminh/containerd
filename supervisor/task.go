package supervisor

import (
	"sync"

	"github.com/docker/containerd/runtime"
)

type StartResponse struct {
	Container runtime.Container
}

type Task interface {
	Error() chan error
}

type baseTask struct {
	err chan error
	m   sync.Mutex
}

func (t *baseTask) Error() chan error {
	t.m.Lock()
	defer t.m.Unlock()
	if t.err == nil {
		t.err = make(chan error, 1)
	}
	return t.err
}
