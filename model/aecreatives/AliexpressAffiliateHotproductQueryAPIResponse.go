package aecreatives

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAffiliateHotproductQueryAPIResponse 查询联盟爆品数据 API返回值
// aliexpress.affiliate.hotproduct.query
//
// 查询联盟爆品API
type AliexpressAffiliateHotproductQueryAPIResponse struct {
	model.CommonResponse
	AliexpressAffiliateHotproductQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressAffiliateHotproductQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressAffiliateHotproductQueryAPIResponseModel).Reset()
}

// AliexpressAffiliateHotproductQueryAPIResponseModel is 查询联盟爆品数据 成功返回结果
type AliexpressAffiliateHotproductQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_affiliate_hotproduct_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	RespResult *ResponseDto `json:"resp_result,omitempty" xml:"resp_result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressAffiliateHotproductQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RespResult = nil
}

var poolAliexpressAffiliateHotproductQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressAffiliateHotproductQueryAPIResponse)
	},
}

// GetAliexpressAffiliateHotproductQueryAPIResponse 从 sync.Pool 获取 AliexpressAffiliateHotproductQueryAPIResponse
func GetAliexpressAffiliateHotproductQueryAPIResponse() *AliexpressAffiliateHotproductQueryAPIResponse {
	return poolAliexpressAffiliateHotproductQueryAPIResponse.Get().(*AliexpressAffiliateHotproductQueryAPIResponse)
}

// ReleaseAliexpressAffiliateHotproductQueryAPIResponse 将 AliexpressAffiliateHotproductQueryAPIResponse 保存到 sync.Pool
func ReleaseAliexpressAffiliateHotproductQueryAPIResponse(v *AliexpressAffiliateHotproductQueryAPIResponse) {
	v.Reset()
	poolAliexpressAffiliateHotproductQueryAPIResponse.Put(v)
}
