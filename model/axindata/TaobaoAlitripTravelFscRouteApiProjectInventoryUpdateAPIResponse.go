package axindata

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIResponse 更新团期库存 API返回值
// taobao.alitrip.travel.fsc.route.api.project.inventory.update
//
// 更新团期库存
type TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIResponseModel).Reset()
}

// TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIResponseModel is 更新团期库存 成功返回结果
type TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_project_inventory_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TopResult = nil
}

var poolTaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIResponse)
	},
}

// GetTaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIResponse
func GetTaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIResponse() *TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIResponse {
	return poolTaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIResponse.Get().(*TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIResponse)
}

// ReleaseTaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIResponse 将 TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIResponse(v *TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIResponse.Put(v)
}
