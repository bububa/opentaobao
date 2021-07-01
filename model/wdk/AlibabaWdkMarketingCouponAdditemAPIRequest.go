package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingCouponAdditemAPIRequest
在优惠券活动下添加商品。【注意，此接口暂不支持并发！】 API请求
alibaba.wdk.marketing.coupon.additem

在优惠券活动下添加商品。【注意，此接口暂不支持并发！】
如果是商品券，则添加的商品为券适用的商品；
如果是品类券，则添加的商品为券排除的商品； */
type AlibabaWdkMarketingCouponAdditemAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *ItemCouponSku
	// 活动基本信息
	_param1 *CommonActivityParam
}

// New
