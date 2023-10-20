package charity

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaValueUserGetidAPIRequest 获取用户userId API请求
// alibaba.value.user.getid
//
// 获取用户userId
type AlibabaValueUserGetidAPIRequest struct {
	model.Params
	// 登录凭证
	_authCode string
	// 渠道编码
	_channelCode string
}

// NewAlibabaValueUserGetidRequest 初始化AlibabaValueUserGetidAPIRequest对象
func NewAlibabaValueUserGetidRequest() *AlibabaValueUserGetidAPIRequest {
	return &AlibabaValueUserGetidAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaValueUserGetidAPIRequest) Reset() {
	r._authCode = ""
	r._channelCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaValueUserGetidAPIRequest) GetApiMethodName() string {
	return "alibaba.value.user.getid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaValueUserGetidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaValueUserGetidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAuthCode is AuthCode Setter
// 登录凭证
func (r *AlibabaValueUserGetidAPIRequest) SetAuthCode(_authCode string) error {
	r._authCode = _authCode
	r.Set("auth_code", _authCode)
	return nil
}

// GetAuthCode AuthCode Getter
func (r AlibabaValueUserGetidAPIRequest) GetAuthCode() string {
	return r._authCode
}

// SetChannelCode is ChannelCode Setter
// 渠道编码
func (r *AlibabaValueUserGetidAPIRequest) SetChannelCode(_channelCode string) error {
	r._channelCode = _channelCode
	r.Set("channel_code", _channelCode)
	return nil
}

// GetChannelCode ChannelCode Getter
func (r AlibabaValueUserGetidAPIRequest) GetChannelCode() string {
	return r._channelCode
}

var poolAlibabaValueUserGetidAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaValueUserGetidRequest()
	},
}

// GetAlibabaValueUserGetidRequest 从 sync.Pool 获取 AlibabaValueUserGetidAPIRequest
func GetAlibabaValueUserGetidAPIRequest() *AlibabaValueUserGetidAPIRequest {
	return poolAlibabaValueUserGetidAPIRequest.Get().(*AlibabaValueUserGetidAPIRequest)
}

// ReleaseAlibabaValueUserGetidAPIRequest 将 AlibabaValueUserGetidAPIRequest 放入 sync.Pool
func ReleaseAlibabaValueUserGetidAPIRequest(v *AlibabaValueUserGetidAPIRequest) {
	v.Reset()
	poolAlibabaValueUserGetidAPIRequest.Put(v)
}
