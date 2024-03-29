package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingCouponAdditemAPIRequest 在优惠券活动下添加商品。【注意，此接口暂不支持并发！】 API请求
// alibaba.wdk.marketing.coupon.additem
//
// 在优惠券活动下添加商品。【注意，此接口暂不支持并发！】
// 如果是商品券，则添加的商品为券适用的商品；
// 如果是品类券，则添加的商品为券排除的商品；
type AlibabaWdkMarketingCouponAdditemAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *ItemCouponSku
	// 活动基本信息
	_param1 *CommonActivityParam
}

// NewAlibabaWdkMarketingCouponAdditemRequest 初始化AlibabaWdkMarketingCouponAdditemAPIRequest对象
func NewAlibabaWdkMarketingCouponAdditemRequest() *AlibabaWdkMarketingCouponAdditemAPIRequest {
	return &AlibabaWdkMarketingCouponAdditemAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingCouponAdditemAPIRequest) Reset() {
	r._param0 = nil
	r._param1 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingCouponAdditemAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.coupon.additem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingCouponAdditemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingCouponAdditemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 商品对象
func (r *AlibabaWdkMarketingCouponAdditemAPIRequest) SetParam0(_param0 *ItemCouponSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaWdkMarketingCouponAdditemAPIRequest) GetParam0() *ItemCouponSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 活动基本信息
func (r *AlibabaWdkMarketingCouponAdditemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaWdkMarketingCouponAdditemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}

var poolAlibabaWdkMarketingCouponAdditemAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingCouponAdditemRequest()
	},
}

// GetAlibabaWdkMarketingCouponAdditemRequest 从 sync.Pool 获取 AlibabaWdkMarketingCouponAdditemAPIRequest
func GetAlibabaWdkMarketingCouponAdditemAPIRequest() *AlibabaWdkMarketingCouponAdditemAPIRequest {
	return poolAlibabaWdkMarketingCouponAdditemAPIRequest.Get().(*AlibabaWdkMarketingCouponAdditemAPIRequest)
}

// ReleaseAlibabaWdkMarketingCouponAdditemAPIRequest 将 AlibabaWdkMarketingCouponAdditemAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingCouponAdditemAPIRequest(v *AlibabaWdkMarketingCouponAdditemAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingCouponAdditemAPIRequest.Put(v)
}
