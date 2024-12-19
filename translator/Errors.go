package translator

import "fmt"

type TranslationError struct {
	Code int
	Msg  string
}

func (e *TranslationError) Error() string {
	return fmt.Sprintf("Code: %d, Msg: %s", e.Code, e.Msg)
}

var TranslationCode = map[string]int{
	"TranslationNotFound":   1,
	"FileLoadingError":      2,
	"FileUnmarshalingError": 3,
}
