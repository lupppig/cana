package cana

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"sync"
)

/*
Package cana  provides a simple HTTP server built from scratch
using the net/tcp package. It supports routing, TLS, and request parsing.

Author: Cu pid
License: MIT
*/

type Cana struct {
	mu        *sync.Mutex
	Addr      string
	quit      chan struct{}
	cListener net.Listener
	conns     map[string]net.Conn
	errCh     chan error
	wg        *sync.WaitGroup
}

type CanaWriter interface {
	Writer(w []byte) error
	Reader(r []byte) error
}

func Canabis(addr string) *Cana {
	if addr == "" {
		addr = ":http"
	}
	return &Cana{
		Addr:  addr,
		mu:    new(sync.Mutex),
		quit:  make(chan struct{}),
		errCh: make(chan error),
		conns: make(map[string]net.Conn),
		wg:    new(sync.WaitGroup),
	}
}

func (c *Cana) ServeCana() error {
	ln, err := net.Listen("tcp", c.Addr)
	if err != nil {
		return err
	}
	c.cListener = ln

	// c.wg.Add(1)
	// go func() {
	// 	defer c.wg.Done()
	c.acceptCana()
	// }()

	return nil
}

func (c *Cana) acceptCana() {
	for {
		select {
		case <-c.quit:
			return
		default:
			conn, err := c.cListener.Accept()
			if err != nil {
				if errors.Is(err, net.ErrClosed) {
					return
				}
				log.Println(err.Error())
				continue
			}
			c.mu.Lock()
			key := conn.RemoteAddr().String()
			c.conns[key] = conn
			c.mu.Unlock()

			c.wg.Add(1)
			go func() {
				defer c.wg.Done()
				c.handleCana(conn)
			}()

		}
	}
}

func (c *Cana) handleCana(conn net.Conn) {
	defer func() {
		c.mu.Lock()
		key := conn.RemoteAddr().String()
		delete(c.conns, key)
		c.mu.Unlock()
	}()

	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return
			}
			if errors.Is(err, net.ErrClosed) {
				return
			}
			log.Println(err.Error())
			continue
		}
		req := newRequest()
		req.httpMethodsParser(buf)

		resp := NewResponse(conn)

		resp.Writer("Hello from this server...\n")
	}

}

func (c *Cana) Shutdown(ctx context.Context) {
	c.cListener.Close()

	c.mu.Lock()
	for _, conn := range c.conns {
		conn.Close()
	}
	c.mu.Unlock()

	done := make(chan struct{})
	go func() {
		close(c.quit)
		c.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		log.Println("server shutdown correctly")
	case <-ctx.Done():
	}
}
