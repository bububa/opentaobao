package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaxintegrationaccountimportAPIRequest ISV用户录入 API请求
// alibaba.tcls.ax.integration.account.import
//
// ISV的用户录入翱象
type AlibabatclsaxintegrationaccountimportAPIRequest struct {
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

// NewAlibabatclsaxintegrationaccountimportRequest 初始化AlibabatclsaxintegrationaccountimportAPIRequest对象
func NewAlibabatclsaxintegrationaccountimportRequest() *AlibabatclsaxintegrationaccountimportAPIRequest {
	return &AlibabatclsaxintegrationaccountimportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatclsaxintegrationaccountimportAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.ax.integration.account.import"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatclsaxintegrationaccountimportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatclsaxintegrationaccountimportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClientId is ClientId Setter
// isv编码
func (r *AlibabatclsaxintegrationaccountimportAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabatclsaxintegrationaccountimportAPIRequest) GetClientId() string {
	return r._clientId
}

// SetMobile is Mobile Setter
// 手机号
func (r *AlibabatclsaxintegrationaccountimportAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r AlibabatclsaxintegrationaccountimportAPIRequest) GetMobile() string {
	return r._mobile
}

// SetEmail is Email Setter
// 邮箱
func (r *AlibabatclsaxintegrationaccountimportAPIRequest) SetEmail(_email string) error {
	r._email = _email
	r.Set("email", _email)
	return nil
}

// GetEmail Email Getter
func (r AlibabatclsaxintegrationaccountimportAPIRequest) GetEmail() string {
	return r._email
}

// SetUsername is Username Setter
// 用户名
func (r *AlibabatclsaxintegrationaccountimportAPIRequest) SetUsername(_username string) error {
	r._username = _username
	r.Set("username", _username)
	return nil
}

// GetUsername Username Getter
func (r AlibabatclsaxintegrationaccountimportAPIRequest) GetUsername() string {
	return r._username
}
