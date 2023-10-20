package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractloginalipayauthAPIRequest 双11到店互动花呗红包获取token鉴权接口 API请求
// alibaba.interact.login.alipayauth
//
// 双11到店互动花呗红包获取token鉴权接口
type AlibabainteractloginalipayauthAPIRequest struct {
	model.Params
}

// NewAlibabainteractloginalipayauthRequest 初始化AlibabainteractloginalipayauthAPIRequest对象
func NewAlibabainteractloginalipayauthRequest() *AlibabainteractloginalipayauthAPIRequest {
	return &AlibabainteractloginalipayauthAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractloginalipayauthAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.login.alipayauth"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractloginalipayauthAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractloginalipayauthAPIRequest) GetRawParams() model.Params {
	return r.Params
}
