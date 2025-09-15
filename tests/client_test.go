package tests

import (
	"context"
	"fmt"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/lupppig/cana"
)

func TestClientConnection(t *testing.T) {
	var wg sync.WaitGroup
	portArr := []int{8801, 8800, 8399, 4011, 2031}

	servers := make([]*cana.Cana, len(portArr))
	for i, port := range portArr {
		c := cana.Canabis(fmt.Sprintf("localhost:%d", port))
		servers[i] = c

		wg.Add(1)
		go func(c *cana.Cana) {
			defer wg.Done()
			if err := c.ServeCana(); err != nil {
				t.Errorf("server failed on %s: %v", c.Addr, err)
			}
		}(c)
	}

	time.Sleep(100 * time.Millisecond)

	for _, port := range portArr {
		dial, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
		if err != nil {
			t.Errorf("failed to dial port %d: %v", port, err)
			continue
		}
		defer dial.Close()

		_, err = dial.Write([]byte(fmt.Sprintf("hey this is from your test client: %d", port)))
		if err != nil {
			t.Errorf("failed to write to server %d: %v", port, err)
			continue
		}

		buf := make([]byte, 1024)
		n, err := dial.Read(buf)
		if err != nil {
			t.Errorf("failed to read from server %d: %v", port, err)
			continue
		}
		t.Logf("server %d responded: %s", port, string(buf[:n]))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	for _, s := range servers {
		s.Shutdown(ctx)
	}

	wg.Wait()
}
