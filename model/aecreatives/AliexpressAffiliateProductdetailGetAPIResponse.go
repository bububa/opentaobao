package aecreatives

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAffiliateProductdetailGetAPIResponse 联盟商品详情获取接口 API返回值
// aliexpress.affiliate.productdetail.get
//
// 联盟推广商品搜索接口，用于搜索联盟推广商品数据
type AliexpressAffiliateProductdetailGetAPIResponse struct {
	model.CommonResponse
	AliexpressAffiliateProductdetailGetAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressAffiliateProductdetailGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressAffiliateProductdetailGetAPIResponseModel).Reset()
}

// AliexpressAffiliateProductdetailGetAPIResponseModel is 联盟商品详情获取接口 成功返回结果
type AliexpressAffiliateProductdetailGetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_affiliate_productdetail_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	RespResult *ResponseDto `json:"resp_result,omitempty" xml:"resp_result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressAffiliateProductdetailGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RespResult = nil
}

var poolAliexpressAffiliateProductdetailGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressAffiliateProductdetailGetAPIResponse)
	},
}

// GetAliexpressAffiliateProductdetailGetAPIResponse 从 sync.Pool 获取 AliexpressAffiliateProductdetailGetAPIResponse
func GetAliexpressAffiliateProductdetailGetAPIResponse() *AliexpressAffiliateProductdetailGetAPIResponse {
	return poolAliexpressAffiliateProductdetailGetAPIResponse.Get().(*AliexpressAffiliateProductdetailGetAPIResponse)
}

// ReleaseAliexpressAffiliateProductdetailGetAPIResponse 将 AliexpressAffiliateProductdetailGetAPIResponse 保存到 sync.Pool
func ReleaseAliexpressAffiliateProductdetailGetAPIResponse(v *AliexpressAffiliateProductdetailGetAPIResponse) {
	v.Reset()
	poolAliexpressAffiliateProductdetailGetAPIResponse.Put(v)
}
