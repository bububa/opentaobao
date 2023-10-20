package aeusergrowth

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressUsergrowthSearchItemsGetAPIResponse 第三方平台搜索AE商品 API返回值
// aliexpress.usergrowth.search.items.get
//
// 第三方平台的搜索服务   获取AE商品list
type AliexpressUsergrowthSearchItemsGetAPIResponse struct {
	model.CommonResponse
	AliexpressUsergrowthSearchItemsGetAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressUsergrowthSearchItemsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressUsergrowthSearchItemsGetAPIResponseModel).Reset()
}

// AliexpressUsergrowthSearchItemsGetAPIResponseModel is 第三方平台搜索AE商品 成功返回结果
type AliexpressUsergrowthSearchItemsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_usergrowth_search_items_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// response model
	Result *AliexpressUsergrowthSearchItemsGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressUsergrowthSearchItemsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressUsergrowthSearchItemsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressUsergrowthSearchItemsGetAPIResponse)
	},
}

// GetAliexpressUsergrowthSearchItemsGetAPIResponse 从 sync.Pool 获取 AliexpressUsergrowthSearchItemsGetAPIResponse
func GetAliexpressUsergrowthSearchItemsGetAPIResponse() *AliexpressUsergrowthSearchItemsGetAPIResponse {
	return poolAliexpressUsergrowthSearchItemsGetAPIResponse.Get().(*AliexpressUsergrowthSearchItemsGetAPIResponse)
}

// ReleaseAliexpressUsergrowthSearchItemsGetAPIResponse 将 AliexpressUsergrowthSearchItemsGetAPIResponse 保存到 sync.Pool
func ReleaseAliexpressUsergrowthSearchItemsGetAPIResponse(v *AliexpressUsergrowthSearchItemsGetAPIResponse) {
	v.Reset()
	poolAliexpressUsergrowthSearchItemsGetAPIResponse.Put(v)
}
