package travel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelItemShelveAPIResponse 【API3.0】度假线路商品上下架接口 API返回值
// taobao.alitrip.travel.item.shelve
//
// 旅行度假新商品发布接口 第三版：度假商品上下架接口
// 注意：定时上下架功能，目前只支持接送、租车类目。
type TaobaoAlitripTravelItemShelveAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelItemShelveAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelItemShelveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelItemShelveAPIResponseModel).Reset()
}

// TaobaoAlitripTravelItemShelveAPIResponseModel is 【API3.0】度假线路商品上下架接口 成功返回结果
type TaobaoAlitripTravelItemShelveAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_item_shelve_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品上下架操作是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelItemShelveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolTaobaoAlitripTravelItemShelveAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelItemShelveAPIResponse)
	},
}

// GetTaobaoAlitripTravelItemShelveAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelItemShelveAPIResponse
func GetTaobaoAlitripTravelItemShelveAPIResponse() *TaobaoAlitripTravelItemShelveAPIResponse {
	return poolTaobaoAlitripTravelItemShelveAPIResponse.Get().(*TaobaoAlitripTravelItemShelveAPIResponse)
}

// ReleaseTaobaoAlitripTravelItemShelveAPIResponse 将 TaobaoAlitripTravelItemShelveAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelItemShelveAPIResponse(v *TaobaoAlitripTravelItemShelveAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelItemShelveAPIResponse.Put(v)
}
