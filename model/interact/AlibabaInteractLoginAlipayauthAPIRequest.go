package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractLoginAlipayauthAPIRequest 双11到店互动花呗红包获取token鉴权接口 API请求
// alibaba.interact.login.alipayauth
//
// 双11到店互动花呗红包获取token鉴权接口
type AlibabaInteractLoginAlipayauthAPIRequest struct {
	model.Params
}

// NewAlibabaInteractLoginAlipayauthRequest 初始化AlibabaInteractLoginAlipayauthAPIRequest对象
func NewAlibabaInteractLoginAlipayauthRequest() *AlibabaInteractLoginAlipayauthAPIRequest {
	return &AlibabaInteractLoginAlipayauthAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractLoginAlipayauthAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractLoginAlipayauthAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.login.alipayauth"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractLoginAlipayauthAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractLoginAlipayauthAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractLoginAlipayauthAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractLoginAlipayauthRequest()
	},
}

// GetAlibabaInteractLoginAlipayauthRequest 从 sync.Pool 获取 AlibabaInteractLoginAlipayauthAPIRequest
func GetAlibabaInteractLoginAlipayauthAPIRequest() *AlibabaInteractLoginAlipayauthAPIRequest {
	return poolAlibabaInteractLoginAlipayauthAPIRequest.Get().(*AlibabaInteractLoginAlipayauthAPIRequest)
}

// ReleaseAlibabaInteractLoginAlipayauthAPIRequest 将 AlibabaInteractLoginAlipayauthAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractLoginAlipayauthAPIRequest(v *AlibabaInteractLoginAlipayauthAPIRequest) {
	v.Reset()
	poolAlibabaInteractLoginAlipayauthAPIRequest.Put(v)
}
