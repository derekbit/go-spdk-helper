package types

import (
	"fmt"
	"testing"
)

// The executor reports a timeout with this wording and then SIGKILLs the command,
// which does nothing to one already blocked in the kernel.
func TestErrorIsTimeoutExecuting(t *testing.T) {
	timeoutErr := fmt.Errorf("timeout executing: %v %v",
		"/usr/sbin/nvme", []string{"nvme", "disconnect", "--device", "/dev/nvme0"})
	if !ErrorIsTimeoutExecuting(timeoutErr) {
		t.Fatalf("expected the executor timeout wording to be recognised: %v", timeoutErr)
	}

	if ErrorIsTimeoutExecuting(fmt.Errorf("nvme: Failed to open /dev/nvme0")) {
		t.Fatal("expected an unrelated failure not to be treated as a timeout")
	}

	if ErrorIsTimeoutExecuting(nil) {
		t.Fatal("expected a nil error not to be treated as a timeout")
	}
}
