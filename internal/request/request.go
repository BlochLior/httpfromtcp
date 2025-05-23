package request

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func parseRequestLine(fullRequest []byte) (*RequestLine, error) {
	str := string(fullRequest)
	partsRaw := strings.Split(str, "\r\n")
	requestLine := partsRaw[0]
	parts := strings.Fields(requestLine)
	method := parts[0]
	if strings.ToUpper(method) != method {
		fmt.Println("error: method is not capitalized")
		return nil, errors.New("method is uncapitalized")
	}

	requestTarget := parts[1]
	httpVersionRaw := parts[2]
	if httpVersionRaw != "HTTP/1.1" {
		fmt.Println("error: http version is not 1.1")
		return nil, errors.New("unfitting http version")
	}
	httpVersion := strings.TrimPrefix(httpVersionRaw, "HTTP/")
	return &RequestLine{
		HttpVersion:   httpVersion,
		RequestTarget: requestTarget,
		Method:        method,
	}, nil

}

func RequestFromReader(reader io.Reader) (*Request, error) {
	entireRequest, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	requestLine, err := parseRequestLine(entireRequest)
	if err != nil {
		return nil, err
	}

	return &Request{
		RequestLine: *requestLine,
	}, nil
}
