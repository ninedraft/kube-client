package cmd

import (
	"git.containerum.net/ch/json-types/user-manager"
	"git.containerum.net/ch/kube-client/pkg/model"
)

const (
	kubeAPIdeploymentPath  = "/namespaces/{namespace}/deployments/{deployment}"
	kubeAPIdeploymentsPath = "/namespaces/{namespace}/deployments"
	resourceDeploymentPath = "namespace/{namespace}/deployment/{deployment}"
)

// GetDeployment -- consumes a namespace and a deployment names,
// returns a Deployment data OR uninitialized struct AND an error
func (client *Client) GetDeployment(namespace, deployment string) (model.Deployment, error) {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"namespace":  namespace,
			"deployment": deployment,
		}).SetResult(model.Deployment{}).
		Get(client.serverURL + kubeAPIdeploymentPath)
	if err != nil {
		return model.Deployment{}, err
	}
	return *resp.Result().(*model.Deployment), nil
}

// GetDeploymentList -- consumes a namespace and a deployment names,
// returns a list of Deployments OR nil slice AND an error
func (client *Client) GetDeploymentList(namespace string) ([]model.Deployment, error) {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"namespace": namespace,
		}).SetResult([]model.Deployment{}).
		Get(client.serverURL + kubeAPIdeploymentsPath)
	if err != nil {
		return nil, err
	}
	return *resp.Result().(*[]model.Deployment), nil
}

// DeleteDeployment -- consumes a namespace, a deployment,
// an user role and an ID
func (client *Client) DeleteDeployment(namespace, deployment, userID, userRole string) error {
	_, err := client.Request.
		SetPathParams(map[string]string{
			"namespace":  namespace,
			"deployment": deployment,
		}).SetHeaders(map[string]string{
		user.UserIDHeader:   userID,
		user.UserRoleHeader: userRole,
	}).Delete(client.resourceServiceAddr + resourceDeploymentPath)
	return err
}

// CreateDeployment -- consumes a namespace, an user ID and a Role,
// returns nil if OK
func (client *Client) CreateDeployment(namespace string, deployment model.Deployment) error {
	_, err := client.Request.
		SetPathParams(map[string]string{
			"namespace": namespace,
		}).SetBody(deployment).
		Post(client.resourceServiceAddr + "/namespace/{namespace}/deployment")
	return err
}
