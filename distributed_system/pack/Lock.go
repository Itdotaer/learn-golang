package pack

import "time"

type Lock struct {
	c chan struct{}
}

func NewLock() Lock {
	var lock Lock
	lock.c = make(chan struct{}, 1)
	lock.c <- struct{}{}
	return lock
}

func (lock Lock) Lock(d time.Duration) bool {
	lockRs := false
	select {
	case <-lock.c:
		lockRs = true
	case <-time.After(d):
		println("time out")
	default:
	}

	return lockRs
}

func (lock Lock) UnLock() {
	lock.c <- struct{}{}
}
