package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingcouponsendmaAPIRequest 发放匿名码 API请求
// alibaba.hm.marketing.coupon.sendma
//
// 根据优惠券活动id打印单个匿名码
type AlibabahmmarketingcouponsendmaAPIRequest struct {
	model.Params
	// 发放匿名码入参
	_param0 *CommonActivityParam
}

// NewAlibabahmmarketingcouponsendmaRequest 初始化AlibabahmmarketingcouponsendmaAPIRequest对象
func NewAlibabahmmarketingcouponsendmaRequest() *AlibabahmmarketingcouponsendmaAPIRequest {
	return &AlibabahmmarketingcouponsendmaAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingcouponsendmaAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.coupon.sendma"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingcouponsendmaAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingcouponsendmaAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 发放匿名码入参
func (r *AlibabahmmarketingcouponsendmaAPIRequest) SetParam0(_param0 *CommonActivityParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabahmmarketingcouponsendmaAPIRequest) GetParam0() *CommonActivityParam {
	return r._param0
}
