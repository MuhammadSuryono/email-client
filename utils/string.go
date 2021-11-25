package utils

import (
	"fmt"
	"strings"
)

func SplitRecipientsEmail(emails string) string {
	emailSplit := strings.Split(emails, ",")

	var emailJoin string
	for i, email := range emailSplit {
		emailJoin += fmt.Sprintf(`"%s"`, email)
		if i != 0 && i != (len(emailSplit)-1) {
			emailJoin += ","
		}
	}
	return emailJoin
}
