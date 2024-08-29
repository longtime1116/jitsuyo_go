package main

import (
	"errors"
	"fmt"

	"github.com/longtime1116/jitsuyo_go/u"
	"golang.org/x/xerrors"
)

var ErrNotFound = errors.New("not found")

type HTTPError struct {
	StatusCode int
	URL        string
}

func (he *HTTPError) Error() string {
	return fmt.Sprintf("http status code: %d, url: %s", he.StatusCode, he.URL)
}

func invokeHTTPError() error {
	return &HTTPError{StatusCode: 500, URL: "https://example.com"}

}

func main() {

	u.PublicSecretP("hoge")
	u.P("hoge")
	// errors.New
	u.P(ErrNotFound)

	// fmt.Errorf
	u.P(fmt.Errorf("length must be greater than 0"))

	// error interface を満たす独自のエラー型を定義する
	if err := invokeHTTPError(); err != nil {
		u.P(err)
	}

	// スタックトレースを出すには xerrors が便利
	err := xerrors.New("xerrors error")
	fmt.Printf("%+v\n", err)

	// %w で wrap する
	baseErr := errors.New("original error")
	wrappedErr := fmt.Errorf("an error occurred: %w", baseErr)

	fmt.Println(wrappedErr) // "an error occurred: original error"

	if errors.Is(wrappedErr, baseErr) {
		fmt.Println("wrappedErr contains baseErr")
	}
}
