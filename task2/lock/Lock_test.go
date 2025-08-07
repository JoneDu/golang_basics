package lock

import "testing"

func TestGoroutineInc(t *testing.T) {
	GoroutineInc()
}

func TestGoroutineFreeLockCounter(t *testing.T) {
	GoroutineFreeLockCounter()
}
