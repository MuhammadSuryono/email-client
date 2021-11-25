package utils

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	er := SplitRecipientsEmail("msuryono0@gmail.com, suryono@mri-research-ind.com")
	fmt.Println(er)
}
