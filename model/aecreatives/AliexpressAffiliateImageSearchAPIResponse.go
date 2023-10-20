package aecreatives

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAffiliateImageSearchAPIResponse 图搜 API返回值
// aliexpress.affiliate.image.search
//
// 图片搜索接口
type AliexpressAffiliateImageSearchAPIResponse struct {
	model.CommonResponse
	AliexpressAffiliateImageSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressAffiliateImageSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressAffiliateImageSearchAPIResponseModel).Reset()
}

// AliexpressAffiliateImageSearchAPIResponseModel is 图搜 成功返回结果
type AliexpressAffiliateImageSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_affiliate_image_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AliexpressAffiliateImageSearchResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressAffiliateImageSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressAffiliateImageSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressAffiliateImageSearchAPIResponse)
	},
}

// GetAliexpressAffiliateImageSearchAPIResponse 从 sync.Pool 获取 AliexpressAffiliateImageSearchAPIResponse
func GetAliexpressAffiliateImageSearchAPIResponse() *AliexpressAffiliateImageSearchAPIResponse {
	return poolAliexpressAffiliateImageSearchAPIResponse.Get().(*AliexpressAffiliateImageSearchAPIResponse)
}

// ReleaseAliexpressAffiliateImageSearchAPIResponse 将 AliexpressAffiliateImageSearchAPIResponse 保存到 sync.Pool
func ReleaseAliexpressAffiliateImageSearchAPIResponse(v *AliexpressAffiliateImageSearchAPIResponse) {
	v.Reset()
	poolAliexpressAffiliateImageSearchAPIResponse.Put(v)
}
