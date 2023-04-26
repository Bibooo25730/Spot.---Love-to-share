package text

import (
	"github.com/google/uuid"
	"strings"
)

func GetUUID() string {
	uuid := uuid.New()
	return strings.ReplaceAll(uuid.String(), "-", "")
}
