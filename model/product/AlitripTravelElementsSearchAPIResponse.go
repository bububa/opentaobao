package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelElementsSearchAPIResponse 商家元素搜索 API返回值
// alitrip.travel.elements.search
//
// 提供商家维护的景点、酒店、餐饮等元素搜索
type AlitripTravelElementsSearchAPIResponse struct {
	model.CommonResponse
	AlitripTravelElementsSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTravelElementsSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTravelElementsSearchAPIResponseModel).Reset()
}

// AlitripTravelElementsSearchAPIResponseModel is 商家元素搜索 成功返回结果
type AlitripTravelElementsSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_elements_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *ResourceData `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTravelElementsSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripTravelElementsSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTravelElementsSearchAPIResponse)
	},
}

// GetAlitripTravelElementsSearchAPIResponse 从 sync.Pool 获取 AlitripTravelElementsSearchAPIResponse
func GetAlitripTravelElementsSearchAPIResponse() *AlitripTravelElementsSearchAPIResponse {
	return poolAlitripTravelElementsSearchAPIResponse.Get().(*AlitripTravelElementsSearchAPIResponse)
}

// ReleaseAlitripTravelElementsSearchAPIResponse 将 AlitripTravelElementsSearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripTravelElementsSearchAPIResponse(v *AlitripTravelElementsSearchAPIResponse) {
	v.Reset()
	poolAlitripTravelElementsSearchAPIResponse.Put(v)
}
