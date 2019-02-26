package caerror

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	ParseCsr        = "parse_csr_failed"
	CsrSignature    = "csr_signature_failed"
	FindCsrBase64   = "find_csr_base64_failed"
	DecodeCsrBase64 = "decode_csr_bBase64_failed"
	DBError         = "database_error"
	NotFound        = "not_found"
)

var (
	ParseCsr2 = New(ParseCsr)
)

type CaError struct {
	errorCode string
	errorMsg  string
	file      string
	line      int
}

func New(errorCode string) error {
	return newCaError(errorCode, "")
}

func NewMsg(errorCode string, message string) error {
	return newCaError(errorCode, message)
}

func NewMsgF(errorCode string, format string, v ...interface{}) error {
	return newCaError(errorCode, fmt.Sprintf(format, v...))
}

func newCaError(errorCode string, errorMsg string) error {

	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}

	return &CaError{
		errorCode: errorCode,
		errorMsg:  errorMsg,
		file:      file,
		line:      line,
	}
}

func (e *CaError) Error() string {
	return e.errorCode
}

func (e *CaError) Message() string {
	return e.errorMsg
}
