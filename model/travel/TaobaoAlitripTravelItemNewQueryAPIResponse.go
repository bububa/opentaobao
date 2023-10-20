package travel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelItemNewQueryAPIResponse 【API3.0】新版度假单个商品查询接口 API返回值
// taobao.alitrip.travel.item.new.query
//
// 新版旅行度假新商品查询接口（单个商品查询）
type TaobaoAlitripTravelItemNewQueryAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelItemNewQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelItemNewQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelItemNewQueryAPIResponseModel).Reset()
}

// TaobaoAlitripTravelItemNewQueryAPIResponseModel is 【API3.0】新版度假单个商品查询接口 成功返回结果
type TaobaoAlitripTravelItemNewQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_item_new_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品查询结果
	TravelItem *FullTravelItem `json:"travel_item,omitempty" xml:"travel_item,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelItemNewQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TravelItem = nil
}

var poolTaobaoAlitripTravelItemNewQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelItemNewQueryAPIResponse)
	},
}

// GetTaobaoAlitripTravelItemNewQueryAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelItemNewQueryAPIResponse
func GetTaobaoAlitripTravelItemNewQueryAPIResponse() *TaobaoAlitripTravelItemNewQueryAPIResponse {
	return poolTaobaoAlitripTravelItemNewQueryAPIResponse.Get().(*TaobaoAlitripTravelItemNewQueryAPIResponse)
}

// ReleaseTaobaoAlitripTravelItemNewQueryAPIResponse 将 TaobaoAlitripTravelItemNewQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelItemNewQueryAPIResponse(v *TaobaoAlitripTravelItemNewQueryAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelItemNewQueryAPIResponse.Put(v)
}
