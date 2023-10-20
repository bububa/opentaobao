package travel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelItemElementQueryAPIResponse 【API3.0】资源元素查询接口 API返回值
// taobao.alitrip.travel.item.element.query
//
// 提供资源元素查询接口，支持商家查询已经发布过的资源元素
type TaobaoAlitripTravelItemElementQueryAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelItemElementQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelItemElementQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelItemElementQueryAPIResponseModel).Reset()
}

// TaobaoAlitripTravelItemElementQueryAPIResponseModel is 【API3.0】资源元素查询接口 成功返回结果
type TaobaoAlitripTravelItemElementQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_item_element_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 资源元素列表
	Results []TopElementParam `json:"results,omitempty" xml:"results>top_element_param,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelItemElementQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolTaobaoAlitripTravelItemElementQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelItemElementQueryAPIResponse)
	},
}

// GetTaobaoAlitripTravelItemElementQueryAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelItemElementQueryAPIResponse
func GetTaobaoAlitripTravelItemElementQueryAPIResponse() *TaobaoAlitripTravelItemElementQueryAPIResponse {
	return poolTaobaoAlitripTravelItemElementQueryAPIResponse.Get().(*TaobaoAlitripTravelItemElementQueryAPIResponse)
}

// ReleaseTaobaoAlitripTravelItemElementQueryAPIResponse 将 TaobaoAlitripTravelItemElementQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelItemElementQueryAPIResponse(v *TaobaoAlitripTravelItemElementQueryAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelItemElementQueryAPIResponse.Put(v)
}
