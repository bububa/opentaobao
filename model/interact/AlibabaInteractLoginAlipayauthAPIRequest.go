package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractLoginAlipayauthAPIRequest
双11到店互动花呗红包获取token鉴权接口 API请求
alibaba.interact.login.alipayauth

双11到店互动花呗红包获取token鉴权接口 */
type AlibabaInteractLoginAlipayauthAPIRequest struct {
	model.Params
}

// NewAlibabaInteractLoginAlipayauthRequest 初始化AlibabaInteractLoginAlipayauthAPIRequest对象
func NewAlibabaInteractLoginAlipayauthRequest() *AlibabaInteractLoginAlipayauthAPIRequest {
	return &AlibabaInteractLoginAlipayauthAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractLoginAlipayauthAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.login.alipayauth"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractLoginAlipayauthAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
