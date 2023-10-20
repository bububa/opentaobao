package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingcouponadditemAPIRequest 在优惠券活动下添加商品。【注意，此接口暂不支持并发！】 API请求
// alibaba.wdk.marketing.coupon.additem
//
// 在优惠券活动下添加商品。【注意，此接口暂不支持并发！】
// 如果是商品券，则添加的商品为券适用的商品；
// 如果是品类券，则添加的商品为券排除的商品；
type AlibabawdkmarketingcouponadditemAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *ItemCouponSku
	// 活动基本信息
	_param1 *CommonActivityParam
}

// NewAlibabawdkmarketingcouponadditemRequest 初始化AlibabawdkmarketingcouponadditemAPIRequest对象
func NewAlibabawdkmarketingcouponadditemRequest() *AlibabawdkmarketingcouponadditemAPIRequest {
	return &AlibabawdkmarketingcouponadditemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingcouponadditemAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.coupon.additem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingcouponadditemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingcouponadditemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 商品对象
func (r *AlibabawdkmarketingcouponadditemAPIRequest) SetParam0(_param0 *ItemCouponSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabawdkmarketingcouponadditemAPIRequest) GetParam0() *ItemCouponSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 活动基本信息
func (r *AlibabawdkmarketingcouponadditemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabawdkmarketingcouponadditemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}
