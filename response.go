package cana

type Response struct {
	StatusCode int
	Headers    map[string]string
	Protocol   string
	Body       []byte
}
