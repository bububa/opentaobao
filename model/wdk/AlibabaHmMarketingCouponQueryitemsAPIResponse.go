package wdk

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaHmMarketingCouponQueryitemsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingCouponQueryitemsAPIResponseModel).Reset()
}

// AlibabaHmMarketingCouponQueryitemsAPIResponseModel is 查询优惠券活动下的商品 成功返回结果
type AlibabaHmMarketingCouponQueryitemsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_coupon_queryitems_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询返回结果
	Result *MarketPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingCouponQueryitemsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingCouponQueryitemsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingCouponQueryitemsAPIResponse)
	},
}

// GetAlibabaHmMarketingCouponQueryitemsAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingCouponQueryitemsAPIResponse
func GetAlibabaHmMarketingCouponQueryitemsAPIResponse() *AlibabaHmMarketingCouponQueryitemsAPIResponse {
	return poolAlibabaHmMarketingCouponQueryitemsAPIResponse.Get().(*AlibabaHmMarketingCouponQueryitemsAPIResponse)
}

// ReleaseAlibabaHmMarketingCouponQueryitemsAPIResponse 将 AlibabaHmMarketingCouponQueryitemsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingCouponQueryitemsAPIResponse(v *AlibabaHmMarketingCouponQueryitemsAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingCouponQueryitemsAPIResponse.Put(v)
}
