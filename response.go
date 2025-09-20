package cana

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"time"
)

const TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"

type WriteResponse interface {
	Writer(data interface{}, statusCode int) error
	write(data interface{}, statusCode int) error
}

type Headers interface {
	SetHeader(key, value string)
	AddHeader(key, value string)
	DelHeader(key string)
	GetHeader(key string) []string
}

type Header map[string]string
type Response struct {
	StatusCode Status
	Headers    Header
	Protocol   string
	Body       []byte
	writer     io.WriteCloser
}

func NewResponse(conn net.Conn) *Response {
	statusCode := StatusOK
	return &Response{
		StatusCode: statusCode,
		Headers:    make(Header),
		Protocol:   "HTTP/1.1",
		Body:       make([]byte, 0),
		writer:     conn,
	}
}

// Set header function used for updating header value
func (h Header) SetHeader(key string, value string) {
	if value, ok := h[MakeHeaderCanonical(key)]; ok {
		val := fmt.Sprintf("%s,%s", h[MakeHeaderCanonical(key)], value)
		h[key] = val
	}
}

// Add header function for setting new header value
func (h Header) AddHeader(key string, value string) {
	h[key] = value
}

func (h Header) AllHeader() Header {
	return h
}
func writeHeader(header Header, key, value string) {
	for key, value := range header {
		header[MakeHeaderCanonical(key)] = value
	}
}

func (r *Response) WriteStatus(statusCode int) {
	r.StatusCode = Status(statusCode)
}

func (r *Response) Writer(data interface{}) error {
	var body []byte
	switch v := data.(type) {
	case []byte:
		body = v
		if _, ok := r.Headers["Content-Type"]; !ok {
			r.Headers["Content-Type"] = "application/octet-stream"
		}
	case string:
		body = []byte(v)
		if _, ok := r.Headers["Content-Type"]; !ok {
			r.Headers["Content-Type"] = "text/plain; charset=utf-8"
		}
	default:
		bod, err := json.Marshal(data)
		if err != nil {
			return err
		}
		body = bod
	}
	r.Body = body
	if err := r.responseParser(); err != nil {
		return err
	}
	return nil
}

// Del function for deleting new header value in key
func (h Header) DelHeader(key string) {
	delete(h, key)
}

// return an array of headers...
func (h Header) GetHeader(key string) []string {
	if value, ok := h[MakeHeaderCanonical(key)]; ok {
		return strings.Split(value, ",")
	}
	return []string{}
}

// parse response to the client
func (r *Response) responseParser() error {
	var resp bytes.Buffer
	resp.WriteString(fmt.Sprintf("HTTP/1.1 %s\r\n", r.StatusCode.StatusText()))
	r.constructHeaders()

	for key, value := range r.Headers {
		resp.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}
	resp.WriteString("\r\n")
	resp.Write(r.Body)

	if err := r.responseWriter(resp.Bytes()); err != nil {
		return err
	}

	return nil
}

func (r *Response) constructHeaders() {
	content_length := strconv.Itoa(len(r.Body))
	r.Headers[MakeHeaderCanonical("date")] = time.Now().UTC().Format(TimeFormat)
	r.Headers[MakeHeaderCanonical("content-length")] = content_length
}

func (r *Response) responseWriter(data []byte) error {
	defer r.writer.Close()
	if _, err := r.writer.Write(data); err != nil {
		return err
	}
	return nil
}

func MakeHeaderCanonical(s string) string {
	if s == "" {
		return ""
	}
	var headerBuilder strings.Builder
	headerBuilder.Grow(len(s))
	upperNext := true
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if upperNext && ch >= 'a' && ch <= 'z' {
			headerBuilder.WriteByte(ch - 'a' + 'A')
		} else if !upperNext && ch >= 'A' && ch <= 'Z' {
			headerBuilder.WriteByte(ch - 'A' + 'a')
		} else {
			headerBuilder.WriteByte(ch)
		}
		upperNext = ch == '-'
	}
	return headerBuilder.String()
}
