package github

import (
	"fmt"

	"github.com/google/uuid"
)

func prepareKeyValueMessage(key, value string) string {
	delimiter := genDelimiter()
	return fmt.Sprintf("%s<<%s\n%s\n%s\n", key, delimiter, value, delimiter)
}

func genDelimiter() string {
	id, _ := uuid.NewV7()
	return fmt.Sprintf("ghadelimiter_%s", id)
}
