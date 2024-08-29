package secret

import "fmt"

func SecretP(a ...any) (n int, err error) {
	return fmt.Println(a...)
}
