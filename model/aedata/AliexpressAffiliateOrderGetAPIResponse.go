package aedata

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAffiliateOrderGetAPIResponse AE流量订单详情获取API API返回值
// aliexpress.affiliate.order.get
//
// 联盟推广订单效果获取API
type AliexpressAffiliateOrderGetAPIResponse struct {
	model.CommonResponse
	AliexpressAffiliateOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressAffiliateOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressAffiliateOrderGetAPIResponseModel).Reset()
}

// AliexpressAffiliateOrderGetAPIResponseModel is AE流量订单详情获取API 成功返回结果
type AliexpressAffiliateOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_affiliate_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	RespResult *ResponseDto `json:"resp_result,omitempty" xml:"resp_result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressAffiliateOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RespResult = nil
}

var poolAliexpressAffiliateOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressAffiliateOrderGetAPIResponse)
	},
}

// GetAliexpressAffiliateOrderGetAPIResponse 从 sync.Pool 获取 AliexpressAffiliateOrderGetAPIResponse
func GetAliexpressAffiliateOrderGetAPIResponse() *AliexpressAffiliateOrderGetAPIResponse {
	return poolAliexpressAffiliateOrderGetAPIResponse.Get().(*AliexpressAffiliateOrderGetAPIResponse)
}

// ReleaseAliexpressAffiliateOrderGetAPIResponse 将 AliexpressAffiliateOrderGetAPIResponse 保存到 sync.Pool
func ReleaseAliexpressAffiliateOrderGetAPIResponse(v *AliexpressAffiliateOrderGetAPIResponse) {
	v.Reset()
	poolAliexpressAffiliateOrderGetAPIResponse.Put(v)
}
