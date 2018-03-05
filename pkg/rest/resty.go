package rest

import (
	"net/http"
	"strings"

	"git.containerum.net/ch/kube-client/pkg/cherry"
	resty "github.com/go-resty/resty"
)

var (
	_ REST = &Resty{}
)

// Resty -- resty client,
// implements REST interface
type Resty struct {
	request *resty.Request
}

// NewResty -- Resty constuctor
func NewResty() *Resty {
	return &Resty{
		request: resty.R(),
	}
}

// Get -- http get method
func (re *Resty) Get(result interface{}, params P, path ...string) error {
	resp, err := re.request.
		SetResult(result).
		SetError(cherry.Err{}).
		SetPathParams(params).
		Get(strings.Join(path, ""))
	if err = MapErrors(resp, err, http.StatusOK); err != nil {
		return err
	}
	copyInterface(result, resp.Result())
	return nil
}

// Put -- http put method
func (re *Resty) Put(result, body interface{}, params P, path ...string) error {
	req := re.request.
		SetError(cherry.Err{}).
		SetPathParams(params)
	if result != nil {
		req = req.SetResult(result)
	}
	if body != nil {
		req = req.SetBody(body)
	}
	resp, err := req.Put(strings.Join(path, ""))
	if err = MapErrors(resp, err,
		http.StatusOK,
		http.StatusAccepted,
		http.StatusCreated); err != nil {
		return err
	}
	if result != nil {
		copyInterface(result, resp.Result())
	}
	return nil
}

// Post -- http post method
func (re *Resty) Post(result, body interface{}, params P, path ...string) error {
	req := re.request.
		SetError(cherry.Err{}).
		SetPathParams(params)
	if result != nil {
		req = req.SetResult(result)
	}
	if body != nil {
		req = req.SetBody(body)
	}
	resp, err := req.Post(strings.Join(path, ""))
	if err = MapErrors(resp, err,
		http.StatusOK,
		http.StatusAccepted,
		http.StatusCreated); err != nil {
		return err
	}
	if result != nil {
		copyInterface(result, resp.Result())
	}
	return nil
}

// Delete -- http delete method
func (re *Resty) Delete(params P, path ...string) error {
	resp, err := re.request.
		SetError(cherry.Err{}).
		SetPathParams(params).
		Post(strings.Join(path, ""))
	return MapErrors(resp, err,
		http.StatusOK,
		http.StatusAccepted,
		http.StatusNoContent)
}
