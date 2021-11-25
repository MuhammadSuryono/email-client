package utils

import (
	"fmt"
	"strings"
)

func SplitRecipientsEmail(emails string) string {
	emailSplit := strings.Split(emails, ",")
	fmt.Println(emailSplit)

	emailJoin := strings.Join(emailSplit, ",")
	return emailJoin
}
