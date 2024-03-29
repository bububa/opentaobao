package axindata

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiPoiApplyAPIResponse 线路供应商提交新增POI申请 API返回值
// taobao.alitrip.travel.fsc.route.api.poi.apply
//
// 线路供应商提交新增POI申请
type TaobaoAlitripTravelFscRouteApiPoiApplyAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelFscRouteApiPoiApplyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscRouteApiPoiApplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelFscRouteApiPoiApplyAPIResponseModel).Reset()
}

// TaobaoAlitripTravelFscRouteApiPoiApplyAPIResponseModel is 线路供应商提交新增POI申请 成功返回结果
type TaobaoAlitripTravelFscRouteApiPoiApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_poi_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoAlitripTravelFscRouteApiPoiApplyTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscRouteApiPoiApplyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TopResult = nil
}

var poolTaobaoAlitripTravelFscRouteApiPoiApplyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiPoiApplyAPIResponse)
	},
}

// GetTaobaoAlitripTravelFscRouteApiPoiApplyAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelFscRouteApiPoiApplyAPIResponse
func GetTaobaoAlitripTravelFscRouteApiPoiApplyAPIResponse() *TaobaoAlitripTravelFscRouteApiPoiApplyAPIResponse {
	return poolTaobaoAlitripTravelFscRouteApiPoiApplyAPIResponse.Get().(*TaobaoAlitripTravelFscRouteApiPoiApplyAPIResponse)
}

// ReleaseTaobaoAlitripTravelFscRouteApiPoiApplyAPIResponse 将 TaobaoAlitripTravelFscRouteApiPoiApplyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelFscRouteApiPoiApplyAPIResponse(v *TaobaoAlitripTravelFscRouteApiPoiApplyAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelFscRouteApiPoiApplyAPIResponse.Put(v)
}
