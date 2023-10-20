package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractcouponapplyAPIRequest 优惠券领取鉴权接口 API请求
// alibaba.interact.coupon.apply
//
// 鉴权接口，为coupon.apply接口鉴权
type AlibabainteractcouponapplyAPIRequest struct {
	model.Params
}

// NewAlibabainteractcouponapplyRequest 初始化AlibabainteractcouponapplyAPIRequest对象
func NewAlibabainteractcouponapplyRequest() *AlibabainteractcouponapplyAPIRequest {
	return &AlibabainteractcouponapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractcouponapplyAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.coupon.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractcouponapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractcouponapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}
