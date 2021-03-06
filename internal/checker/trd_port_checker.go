package checker

import (
	"errors"
	"fmt"
	"net/http"
)

const (
	trdUrl         = "http://localhost:6732/authorized_keys"
	trdSignerTitle = "SignerChecker"
)

type trdSignerChecker struct {
}

func (t *trdSignerChecker) AssertRunning() (err error) {
	resp, err := http.Get(trdUrl)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("%s", resp.Status))
	}
	return nil
}

func (t *trdSignerChecker) GetTitle() string {
	return trdSignerTitle
}

func NewTRDChecker() (c Checker, err error) {
	return &trdSignerChecker{}, nil
}
