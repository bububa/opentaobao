package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingCouponQueryitemsAPIResponse 查询优惠券活动下的商品 API返回值
// alibaba.hm.marketing.coupon.queryitems
//
// 查询优惠券活动下面的商品
type AlibabaHmMarketingCouponQueryitemsAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingCouponQueryitemsAPIResponseModel
}

// AlibabaHmMarketingCouponQueryitemsAPIResponseModel is 查询优惠券活动下的商品 成功返回结果
type AlibabaHmMarketingCouponQueryitemsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_coupon_queryitems_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询返回结果
	Result *MarketPageResult `json:"result,omitempty" xml:"result,omitempty"`
}
