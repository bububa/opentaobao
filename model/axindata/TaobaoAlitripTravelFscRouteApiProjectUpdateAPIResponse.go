package axindata

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiProjectUpdateAPIResponse 更新团期 API返回值
// taobao.alitrip.travel.fsc.route.api.project.update
//
// 更新团期
type TaobaoAlitripTravelFscRouteApiProjectUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelFscRouteApiProjectUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscRouteApiProjectUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelFscRouteApiProjectUpdateAPIResponseModel).Reset()
}

// TaobaoAlitripTravelFscRouteApiProjectUpdateAPIResponseModel is 更新团期 成功返回结果
type TaobaoAlitripTravelFscRouteApiProjectUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_project_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoAlitripTravelFscRouteApiProjectUpdateTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscRouteApiProjectUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TopResult = nil
}

var poolTaobaoAlitripTravelFscRouteApiProjectUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiProjectUpdateAPIResponse)
	},
}

// GetTaobaoAlitripTravelFscRouteApiProjectUpdateAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelFscRouteApiProjectUpdateAPIResponse
func GetTaobaoAlitripTravelFscRouteApiProjectUpdateAPIResponse() *TaobaoAlitripTravelFscRouteApiProjectUpdateAPIResponse {
	return poolTaobaoAlitripTravelFscRouteApiProjectUpdateAPIResponse.Get().(*TaobaoAlitripTravelFscRouteApiProjectUpdateAPIResponse)
}

// ReleaseTaobaoAlitripTravelFscRouteApiProjectUpdateAPIResponse 将 TaobaoAlitripTravelFscRouteApiProjectUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelFscRouteApiProjectUpdateAPIResponse(v *TaobaoAlitripTravelFscRouteApiProjectUpdateAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelFscRouteApiProjectUpdateAPIResponse.Put(v)
}
