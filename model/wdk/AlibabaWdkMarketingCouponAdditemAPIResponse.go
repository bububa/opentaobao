package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingcouponadditemAPIResponse 在优惠券活动下添加商品。【注意，此接口暂不支持并发！】 API返回值
// alibaba.wdk.marketing.coupon.additem
//
// 在优惠券活动下添加商品。【注意，此接口暂不支持并发！】
// 如果是商品券，则添加的商品为券适用的商品；
// 如果是品类券，则添加的商品为券排除的商品；
type AlibabawdkmarketingcouponadditemAPIResponse struct {
	model.CommonResponse
	AlibabawdkmarketingcouponadditemAPIResponseModel
}

// AlibabawdkmarketingcouponadditemAPIResponseModel is 在优惠券活动下添加商品。【注意，此接口暂不支持并发！】 成功返回结果
type AlibabawdkmarketingcouponadditemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_coupon_additem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品报名活动的返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
