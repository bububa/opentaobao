package aeusergrowth

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressUsergrowthRecommendItemsGetAPIResponse 第三方平台推荐商品 API返回值
// aliexpress.usergrowth.recommend.items.get
//
// 第三方平台的推荐AE商品   场景：skin 、底部推荐等
type AliexpressUsergrowthRecommendItemsGetAPIResponse struct {
	model.CommonResponse
	AliexpressUsergrowthRecommendItemsGetAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressUsergrowthRecommendItemsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressUsergrowthRecommendItemsGetAPIResponseModel).Reset()
}

// AliexpressUsergrowthRecommendItemsGetAPIResponseModel is 第三方平台推荐商品 成功返回结果
type AliexpressUsergrowthRecommendItemsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_usergrowth_recommend_items_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// response model
	Result *AliexpressUsergrowthRecommendItemsGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressUsergrowthRecommendItemsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressUsergrowthRecommendItemsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressUsergrowthRecommendItemsGetAPIResponse)
	},
}

// GetAliexpressUsergrowthRecommendItemsGetAPIResponse 从 sync.Pool 获取 AliexpressUsergrowthRecommendItemsGetAPIResponse
func GetAliexpressUsergrowthRecommendItemsGetAPIResponse() *AliexpressUsergrowthRecommendItemsGetAPIResponse {
	return poolAliexpressUsergrowthRecommendItemsGetAPIResponse.Get().(*AliexpressUsergrowthRecommendItemsGetAPIResponse)
}

// ReleaseAliexpressUsergrowthRecommendItemsGetAPIResponse 将 AliexpressUsergrowthRecommendItemsGetAPIResponse 保存到 sync.Pool
func ReleaseAliexpressUsergrowthRecommendItemsGetAPIResponse(v *AliexpressUsergrowthRecommendItemsGetAPIResponse) {
	v.Reset()
	poolAliexpressUsergrowthRecommendItemsGetAPIResponse.Put(v)
}
