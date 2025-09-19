package cana

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"io"
	"strconv"
	"strings"
)

type Request struct {
	Headers        map[string]string
	Protocol       string
	Method         string
	Body           []byte
	Routes         string
	Path           string
	QueryParameter string
	Ctx            context.Context
}

func newRequest() *Request {
	return &Request{
		Headers: make(map[string]string),
	}
}

func (r *Request) httpMethodsParser(data []byte) error {
	var headerBuilder strings.Builder
	var contentBuilder bytes.Buffer
	var method string
	var err error

	delim := []byte("\r\n\r\n")
	idx := bytes.Index(data, delim)
	if idx != -1 {
		_, err = headerBuilder.Write(data[:idx])
		if err != nil {
			return err
		}
	} else {
		_, err = headerBuilder.Write(data)
		if err != nil {
			return err
		}
	}
	request_headers := strings.Split(headerBuilder.String(), "\r\n")
	for _, headers := range request_headers {
		if strings.Contains(headers, "POST") ||
			strings.Contains(headers, "GET") ||
			strings.Contains(headers, "DELETE") ||
			strings.Contains(headers, "PATCH") ||
			strings.Contains(headers, "PUT") ||
			strings.Contains(headers, "HEADS") ||
			strings.Contains(headers, "OPTIONS") {
			r.Method, r.Protocol, r.Path, err = parse_request_method(headers)
			if err != nil {
				return err
			}
			method = r.Method
			continue
		}
		if strings.Contains(headers, ":") {
			spl_head := strings.SplitN(headers, ":", 2)
			key, value := spl_head[0], spl_head[1]
			key = strings.ToLower(strings.TrimSpace(key))
			r.Headers[key] = strings.TrimSpace(value)
			continue
		}
		if headers != " " && headers != "" {
			headers = strings.ToLower(strings.TrimSpace(headers))
			r.Headers[headers] = ""
		}
	}

	if t_encode, ok := r.Headers["transfer-encoding"]; ok {
		var reader = bufio.NewReader(strings.NewReader(string(data[idx+len(delim):])))
		switch t_encode {
		case "chunked":
			for {
				line, err := reader.ReadString('\n')
				if err != nil {
					return err
				}
				hexStr := strings.TrimSpace(line)
				size, _ := strconv.ParseInt(hexStr, 16, 64)

				if size == 0 {
					reader.ReadString('\n')
					break
				}

				buf := make([]byte, size)
				if _, err := io.ReadFull(reader, buf); err != nil {
					return err
				}
				if _, err := contentBuilder.Write(buf); err != nil {
					return err
				}
				if _, err := reader.ReadString('\n'); err != nil {
					return err
				}
			}
		}
	} else if content_len, ok := r.Headers["content-length"]; ok {
		n, _ := strconv.Atoi(content_len)
		data = data[idx+len(delim):]
		for _, byt := range data[:n] {
			contentBuilder.WriteByte(byt)
		}
	}

	r.Body = contentBuilder.Bytes()

	contentBuilder.Reset()
	switch method {
	case "GET":
	case "POST":
	case "PUT":
	case "PATCH":
	case "DELETE":
	}
	return nil
}

func parse_request_method(headers string) (string, string, string, error) {
	split_str := strings.Split(headers, " ")

	if len(split_str) != 3 {
		return "", "", "", errors.New("invalid request target header")
	}
	method := split_str[0]

	if method != "POST" &&
		method != "PUT" &&
		method != "PATCH" &&
		method != "GET" &&
		method != "OPTION" &&
		method != "HEADS" && method != "DELETE" {
		return "", "", "", errors.New("missing proper http method")
	}
	path := split_str[1]

	if !strings.Contains(path, "/") {
		return "", "", "", errors.New("missing path value in request target")
	}

	http_protocol := split_str[2]

	if !strings.Contains(http_protocol, "HTTP/1.1") {
		return "", "", "", errors.New("missing proper http protocol")
	}
	return method, path, http_protocol, nil
}
