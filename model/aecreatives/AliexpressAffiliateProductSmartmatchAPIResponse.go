package aecreatives

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAffiliateProductSmartmatchAPIResponse 联盟物料智能推荐api API返回值
// aliexpress.affiliate.product.smartmatch
//
// 联盟物料算法智能推荐
type AliexpressAffiliateProductSmartmatchAPIResponse struct {
	model.CommonResponse
	AliexpressAffiliateProductSmartmatchAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressAffiliateProductSmartmatchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressAffiliateProductSmartmatchAPIResponseModel).Reset()
}

// AliexpressAffiliateProductSmartmatchAPIResponseModel is 联盟物料智能推荐api 成功返回结果
type AliexpressAffiliateProductSmartmatchAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_affiliate_product_smartmatch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	RespResult *ResponseDto `json:"resp_result,omitempty" xml:"resp_result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressAffiliateProductSmartmatchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RespResult = nil
}

var poolAliexpressAffiliateProductSmartmatchAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressAffiliateProductSmartmatchAPIResponse)
	},
}

// GetAliexpressAffiliateProductSmartmatchAPIResponse 从 sync.Pool 获取 AliexpressAffiliateProductSmartmatchAPIResponse
func GetAliexpressAffiliateProductSmartmatchAPIResponse() *AliexpressAffiliateProductSmartmatchAPIResponse {
	return poolAliexpressAffiliateProductSmartmatchAPIResponse.Get().(*AliexpressAffiliateProductSmartmatchAPIResponse)
}

// ReleaseAliexpressAffiliateProductSmartmatchAPIResponse 将 AliexpressAffiliateProductSmartmatchAPIResponse 保存到 sync.Pool
func ReleaseAliexpressAffiliateProductSmartmatchAPIResponse(v *AliexpressAffiliateProductSmartmatchAPIResponse) {
	v.Reset()
	poolAliexpressAffiliateProductSmartmatchAPIResponse.Put(v)
}
