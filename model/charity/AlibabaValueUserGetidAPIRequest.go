package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabavalueusergetidAPIRequest 获取用户userId API请求
// alibaba.value.user.getid
//
// 获取用户userId
type AlibabavalueusergetidAPIRequest struct {
	model.Params
	// 登录凭证
	_authCode string
	// 渠道编码
	_channelCode string
}

// NewAlibabavalueusergetidRequest 初始化AlibabavalueusergetidAPIRequest对象
func NewAlibabavalueusergetidRequest() *AlibabavalueusergetidAPIRequest {
	return &AlibabavalueusergetidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabavalueusergetidAPIRequest) GetApiMethodName() string {
	return "alibaba.value.user.getid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabavalueusergetidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabavalueusergetidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAuthCode is AuthCode Setter
// 登录凭证
func (r *AlibabavalueusergetidAPIRequest) SetAuthCode(_authCode string) error {
	r._authCode = _authCode
	r.Set("auth_code", _authCode)
	return nil
}

// GetAuthCode AuthCode Getter
func (r AlibabavalueusergetidAPIRequest) GetAuthCode() string {
	return r._authCode
}

// SetChannelCode is ChannelCode Setter
// 渠道编码
func (r *AlibabavalueusergetidAPIRequest) SetChannelCode(_channelCode string) error {
	r._channelCode = _channelCode
	r.Set("channel_code", _channelCode)
	return nil
}

// GetChannelCode ChannelCode Getter
func (r AlibabavalueusergetidAPIRequest) GetChannelCode() string {
	return r._channelCode
}
