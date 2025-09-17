package cana

import (
	"bytes"
	"context"
	"errors"
	"strconv"
	"strings"
)

type Request struct {
	Headers        map[string]interface{}
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
		Headers: make(map[string]interface{}),
	}
}

func (r *Request) httpMethods(data []byte) error {
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
			spl_head := strings.Split(headers, ":")
			key, value := spl_head[0], spl_head[1]
			r.Headers[key] = value
			continue
		}
		if headers != " " && headers != "" {
			r.Headers[headers] = struct{}{}
		}
	}

	if content_len, ok := r.Headers["Content-Length"]; ok {
		n, _ := strconv.Atoi(content_len.(string))
		data = data[idx+len(delim):]
		_, err = contentBuilder.Write(data[:n])
		if err != nil {
			return err
		}
		r.Body = contentBuilder.Bytes()
	}

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
