package icbudropshipping

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDropshippingTokenCreateAPIRequest 国际站dropshipping 选品token 创建 API请求
// alibaba.dropshipping.token.create
//
// 国际站dropshipping 选品token 创建，用于让买家有权限访问我们指定的 商品场馆
type AlibabaDropshippingTokenCreateAPIRequest struct {
	model.Params
}

// NewAlibabaDropshippingTokenCreateRequest 初始化AlibabaDropshippingTokenCreateAPIRequest对象
func NewAlibabaDropshippingTokenCreateRequest() *AlibabaDropshippingTokenCreateAPIRequest {
	return &AlibabaDropshippingTokenCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDropshippingTokenCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.dropshipping.token.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDropshippingTokenCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
