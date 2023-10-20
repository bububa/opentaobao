package b2bcert

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaauthcertgetAPIRequest 获取证书数据 API请求
// alibaba.auth.cert.get
//
// 获取证书数据
type AlibabaauthcertgetAPIRequest struct {
	model.Params
	// 认证商
	_provider string
	// 证书数据
	_receiveInfo string
}

// NewAlibabaauthcertgetRequest 初始化AlibabaauthcertgetAPIRequest对象
func NewAlibabaauthcertgetRequest() *AlibabaauthcertgetAPIRequest {
	return &AlibabaauthcertgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaauthcertgetAPIRequest) GetApiMethodName() string {
	return "alibaba.auth.cert.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaauthcertgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaauthcertgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProvider is Provider Setter
// 认证商
func (r *AlibabaauthcertgetAPIRequest) SetProvider(_provider string) error {
	r._provider = _provider
	r.Set("provider", _provider)
	return nil
}

// GetProvider Provider Getter
func (r AlibabaauthcertgetAPIRequest) GetProvider() string {
	return r._provider
}

// SetReceiveInfo is ReceiveInfo Setter
// 证书数据
func (r *AlibabaauthcertgetAPIRequest) SetReceiveInfo(_receiveInfo string) error {
	r._receiveInfo = _receiveInfo
	r.Set("receive_info", _receiveInfo)
	return nil
}

// GetReceiveInfo ReceiveInfo Getter
func (r AlibabaauthcertgetAPIRequest) GetReceiveInfo() string {
	return r._receiveInfo
}
