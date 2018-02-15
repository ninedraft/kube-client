package cmd

import (
	"net/http"

	"git.containerum.net/ch/kube-client/pkg/model"
)

const (
	userInfoPath           = "/user/info"
	userPasswordChangePath = "/password/change"
)

// GetProfileInfo -- returns user info
func (client *Client) GetProfileInfo() (model.User, error) {
	resp, err := client.Request.
		SetResult(model.User{}).
		SetError(model.ResourceError{}).
		Get(client.UserManagerURL + userInfoPath)
	if err := catchErr(err, resp, http.StatusOK); err != nil {
		return model.User{}, err
	}
	return *resp.Result().(*model.User), nil
}

// ChangePassword -- changes user password, returns access and refresh tokens
func (client *Client) ChangePassword(currentPassword, newPassword string) (model.Tokens, error) {
	resp, err := client.Request.
		SetResult(model.Tokens{}).
		SetError(model.ResourceError{}).
		Get(client.UserManagerURL + userPasswordChangePath)
	if err := catchErr(err, resp, http.StatusAccepted, http.StatusOK); err != nil {
		return model.Tokens{}, err
	}
	return *resp.Error().(*model.Tokens), nil
}