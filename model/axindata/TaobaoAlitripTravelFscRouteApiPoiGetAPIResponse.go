package axindata

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiPoiGetAPIResponse 获取景点（POI）信息 API返回值
// taobao.alitrip.travel.fsc.route.api.poi.get
//
// 获取景点（POI）信息
type TaobaoAlitripTravelFscRouteApiPoiGetAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelFscRouteApiPoiGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscRouteApiPoiGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelFscRouteApiPoiGetAPIResponseModel).Reset()
}

// TaobaoAlitripTravelFscRouteApiPoiGetAPIResponseModel is 获取景点（POI）信息 成功返回结果
type TaobaoAlitripTravelFscRouteApiPoiGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_poi_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoAlitripTravelFscRouteApiPoiGetTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscRouteApiPoiGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TopResult = nil
}

var poolTaobaoAlitripTravelFscRouteApiPoiGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiPoiGetAPIResponse)
	},
}

// GetTaobaoAlitripTravelFscRouteApiPoiGetAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelFscRouteApiPoiGetAPIResponse
func GetTaobaoAlitripTravelFscRouteApiPoiGetAPIResponse() *TaobaoAlitripTravelFscRouteApiPoiGetAPIResponse {
	return poolTaobaoAlitripTravelFscRouteApiPoiGetAPIResponse.Get().(*TaobaoAlitripTravelFscRouteApiPoiGetAPIResponse)
}

// ReleaseTaobaoAlitripTravelFscRouteApiPoiGetAPIResponse 将 TaobaoAlitripTravelFscRouteApiPoiGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelFscRouteApiPoiGetAPIResponse(v *TaobaoAlitripTravelFscRouteApiPoiGetAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelFscRouteApiPoiGetAPIResponse.Put(v)
}
