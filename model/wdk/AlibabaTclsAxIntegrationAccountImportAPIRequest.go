package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAxIntegrationAccountImportAPIRequest ISV用户录入 API请求
// alibaba.tcls.ax.integration.account.import
//
// ISV的用户录入翱象
type AlibabaTclsAxIntegrationAccountImportAPIRequest struct {
	model.Params
	// isv编码
	_clientId string
	// 手机号
	_mobile string
	// 邮箱
	_email string
	// 用户名
	_username string
}

// NewAlibabaTclsAxIntegrationAccountImportRequest 初始化AlibabaTclsAxIntegrationAccountImportAPIRequest对象
func NewAlibabaTclsAxIntegrationAccountImportRequest() *AlibabaTclsAxIntegrationAccountImportAPIRequest {
	return &AlibabaTclsAxIntegrationAccountImportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAxIntegrationAccountImportAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.ax.integration.account.import"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAxIntegrationAccountImportAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetClientId is ClientId Setter
// isv编码
func (r *AlibabaTclsAxIntegrationAccountImportAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaTclsAxIntegrationAccountImportAPIRequest) GetClientId() string {
	return r._clientId
}

// SetMobile is Mobile Setter
// 手机号
func (r *AlibabaTclsAxIntegrationAccountImportAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r AlibabaTclsAxIntegrationAccountImportAPIRequest) GetMobile() string {
	return r._mobile
}

// SetEmail is Email Setter
// 邮箱
func (r *AlibabaTclsAxIntegrationAccountImportAPIRequest) SetEmail(_email string) error {
	r._email = _email
	r.Set("email", _email)
	return nil
}

// GetEmail Email Getter
func (r AlibabaTclsAxIntegrationAccountImportAPIRequest) GetEmail() string {
	return r._email
}

// SetUsername is Username Setter
// 用户名
func (r *AlibabaTclsAxIntegrationAccountImportAPIRequest) SetUsername(_username string) error {
	r._username = _username
	r.Set("username", _username)
	return nil
}

// GetUsername Username Getter
func (r AlibabaTclsAxIntegrationAccountImportAPIRequest) GetUsername() string {
	return r._username
}
