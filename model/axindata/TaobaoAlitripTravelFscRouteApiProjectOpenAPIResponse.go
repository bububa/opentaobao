package axindata

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiProjectOpenAPIResponse 打开团期 API返回值
// taobao.alitrip.travel.fsc.route.api.project.open
//
// 打开团期
type TaobaoAlitripTravelFscRouteApiProjectOpenAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelFscRouteApiProjectOpenAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscRouteApiProjectOpenAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelFscRouteApiProjectOpenAPIResponseModel).Reset()
}

// TaobaoAlitripTravelFscRouteApiProjectOpenAPIResponseModel is 打开团期 成功返回结果
type TaobaoAlitripTravelFscRouteApiProjectOpenAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_project_open_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoAlitripTravelFscRouteApiProjectOpenTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscRouteApiProjectOpenAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TopResult = nil
}

var poolTaobaoAlitripTravelFscRouteApiProjectOpenAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiProjectOpenAPIResponse)
	},
}

// GetTaobaoAlitripTravelFscRouteApiProjectOpenAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelFscRouteApiProjectOpenAPIResponse
func GetTaobaoAlitripTravelFscRouteApiProjectOpenAPIResponse() *TaobaoAlitripTravelFscRouteApiProjectOpenAPIResponse {
	return poolTaobaoAlitripTravelFscRouteApiProjectOpenAPIResponse.Get().(*TaobaoAlitripTravelFscRouteApiProjectOpenAPIResponse)
}

// ReleaseTaobaoAlitripTravelFscRouteApiProjectOpenAPIResponse 将 TaobaoAlitripTravelFscRouteApiProjectOpenAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelFscRouteApiProjectOpenAPIResponse(v *TaobaoAlitripTravelFscRouteApiProjectOpenAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelFscRouteApiProjectOpenAPIResponse.Put(v)
}
