package cana

type Status int

const (
	StatusContinue          Status = 100
	StatusSwitchingProtocol Status = 101
	StatusProcessing        Status = 102
	StatusEarlyHints        Status = 103

	// successuful response
	StatusOK                        Status = 200
	StatusCreated                   Status = 201
	StatusAccepted                  Status = 202
	StatusNonAuthorativeInformation Status = 203
	StatusNoContent                 Status = 204
	StatusResetContent              Status = 205
	StatusPartialContent            Status = 206
	StatusMultiStatus               Status = 207
	StatusAlreadyReported           Status = 208
	StatusIMUsed                    Status = 226

	// redirection messages
	StatusMultipleChoices   Status = 300
	StatusMovedPermanently  Status = 301
	StatusFound             Status = 302
	StatusSeeOther          Status = 303
	StatusNotModified       Status = 304
	StatusUseProxy          Status = 305
	StatusUnUsed            Status = 306
	StatusTemporaryRedirect Status = 307
	StatusPermanentRedirect Status = 308

	// client error responses

	StatusBadRequest                  Status = 400
	StatusUnAuthorized                Status = 401
	StatusPaymentRequired             Status = 402
	StatusForbidden                   Status = 403
	StatusNotFound                    Status = 404
	StatusMethodNotAllowed            Status = 405
	StatusNotAcceptable               Status = 406
	StatusProxyAuthenticationRequired Status = 407
	StatusRequestTimeOut              Status = 408
	StatusConflict                    Status = 409
	StatusGone                        Status = 410
	StatusLengthRequired              Status = 411
	StatusPreconditionFailed          Status = 412
	StatusContentTooLarge             Status = 413
	StatusURITooLong                  Status = 414
	StatusUnSupportedMediaType        Status = 415
	StatusRangeNotSatisfiable         Status = 416
	StatusExpectationFailed           Status = 417
	StatusImATeapot                   Status = 418
	StatusMisdirectedRequest          Status = 421
	StatusUnprocessableContent        Status = 422
	StatusLocked                      Status = 423
	StatusFailedDependency            Status = 424
	StatusTooEarly                    Status = 425
	StatusUpgradeRequired             Status = 426
	StatusPreconditionRequired        Status = 428
	StatusTooManyRequest              Status = 429
	StatusRequestHeaderFieldsTooLarge Status = 431
	StatusUnavailableForLegalReasons  Status = 451

	// server  error responses
	StatusInternalServerError           Status = 500
	StatusNotImplemented                Status = 501
	StatusBadGateWay                    Status = 502
	StatusServiceUnavailable            Status = 503
	StatusGateWayTimeOut                Status = 504
	StatusHttpVersionNotSupported       Status = 505
	StatusVariantAlsoNegotiates         Status = 506
	StatusInsufficientStorage           Status = 507
	StatusLoopDetected                  Status = 508
	StatusNotExtended                   Status = 510
	StatusNetworkAuthenticationRequired Status = 511
)

func (s Status) StatusText() string {
	switch s {
	case StatusContinue:
		return "100 Continue"
	case StatusSwitchingProtocol:
		return "100 Switching Protcols"
	case StatusProcessing:
		return "102 Processing"
	case StatusEarlyHints:
		return "103 Early Hints"

	case StatusOK:
		return "200 OK"
	case StatusCreated:
		return "201 Created"
	case StatusAccepted:
		return "202 Accepted"
	case StatusNonAuthorativeInformation:
		return "203 Non-Authoritative Information"
	case StatusNoContent:
		return "204 No Content"
	case StatusResetContent:
		return "205 Reset Content"
	case StatusPartialContent:
		return "206 Partial Content"
	case StatusMultiStatus:
		return "207 Multi-Status"
	case StatusIMUsed:
		return "226 IM Used"

	case StatusMultipleChoices:
		return "300 Multiple Choices"
	case StatusMovedPermanently:
		return "301 Moved Permanently"
	case StatusFound:
		return "302 Found"
	case StatusSeeOther:
		return "303 See Other"
	case StatusNotModified:
		return "304 Not Modified"
	case StatusUseProxy:
		return "305 Use Proxy"
	case StatusUnUsed:
		return "306 unused"
	case StatusTemporaryRedirect:
		return "307 Temporary Redirect"
	case StatusPermanentRedirect:
		return "308 Permanent Redirect"

	case StatusBadRequest:
		return "400 Bad Request"
	case StatusUnAuthorized:
		return "402 Payment Required"
	case StatusForbidden:
		return "403 Forbidden"
	case StatusNotFound:
		return "404 Not Found"
	case StatusMethodNotAllowed:
		return "405 Method Not Allowed"
	case StatusNotAcceptable:
		return "406 Not Acceptable"
	case StatusProxyAuthenticationRequired:
		return "407 Proxy Authentication Required"
	case StatusRequestTimeOut:
		return "408 Request Timeout"
	case StatusConflict:
		return "409 Conflict"
	case StatusGone:
		return "410 Gone"
	case StatusLengthRequired:
		return "411 Length Required"
	case StatusPreconditionFailed:
		return "412 Preconditioned Failed"
	case StatusContentTooLarge:
		return "413 Content Too Large"
	case StatusURITooLong:
		return "414 URI Too Long"
	case StatusUnSupportedMediaType:
		return "415 Unsupported Media Type"
	case StatusRangeNotSatisfiable:
		return "416 Range Not Satisfiable"
	case StatusExpectationFailed:
		return "417 Expectation Failed"
	case StatusImATeapot:
		return "418 I'm a teapot"
	case StatusMisdirectedRequest:
		return "421 Misdirected Request"
	case StatusUnprocessableContent:
		return "422 Unprocessable Content"
	case StatusLocked:
		return "423 Locked"
	case StatusFailedDependency:
		return "424 Failed Dependency"
	case StatusTooEarly:
		return "425 Too Early"
	case StatusUpgradeRequired:
		return "426 Upgrade Required"
	case StatusPreconditionRequired:
		return "428 Precondition Required"
	case StatusTooManyRequest:
		return "429 Too Many Request"
	case StatusRequestHeaderFieldsTooLarge:
		return "431 Request Header Fields Too Large"
	case StatusUnavailableForLegalReasons:
		return "451 Unavailable For Legal Reasons"

	case StatusInternalServerError:
		return "500 Internal Server Error"
	case StatusNotImplemented:
		return "501 Not Implemented"
	case StatusBadGateWay:
		return "502 Bad Gateway"
	case StatusServiceUnavailable:
		return "503 Service Unavailable"
	case StatusGateWayTimeOut:
		return "504 Gateway Timeout"
	case StatusHttpVersionNotSupported:
		return "505 HTTP Version Not Supported"
	case StatusVariantAlsoNegotiates:
		return "506 Variant Also Negotiates"
	case StatusInsufficientStorage:
		return "507 Insufficient Storages"
	case StatusLoopDetected:
		return "508 Loop Detected"
	case StatusNotExtended:
		return "510 Not Extended"
	case StatusNetworkAuthenticationRequired:
		return "511 Network Authentication Required"
	}
	return ""
}
