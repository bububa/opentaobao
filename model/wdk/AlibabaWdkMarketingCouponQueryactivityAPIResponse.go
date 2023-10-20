package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingcouponqueryactivityAPIResponse 查询优惠券活动 API返回值
// alibaba.wdk.marketing.coupon.queryactivity
//
// 查询优惠券活动
type AlibabawdkmarketingcouponqueryactivityAPIResponse struct {
	model.CommonResponse
	AlibabawdkmarketingcouponqueryactivityAPIResponseModel
}

// AlibabawdkmarketingcouponqueryactivityAPIResponseModel is 查询优惠券活动 成功返回结果
type AlibabawdkmarketingcouponqueryactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_coupon_queryactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
