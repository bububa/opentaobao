package axindata

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiDivisionGetAPIResponse 获取国家城市信息 API返回值
// taobao.alitrip.travel.fsc.route.api.division.get
//
// 获取国家城市信息
type TaobaoAlitripTravelFscRouteApiDivisionGetAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelFscRouteApiDivisionGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscRouteApiDivisionGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelFscRouteApiDivisionGetAPIResponseModel).Reset()
}

// TaobaoAlitripTravelFscRouteApiDivisionGetAPIResponseModel is 获取国家城市信息 成功返回结果
type TaobaoAlitripTravelFscRouteApiDivisionGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_division_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoAlitripTravelFscRouteApiDivisionGetTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscRouteApiDivisionGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TopResult = nil
}

var poolTaobaoAlitripTravelFscRouteApiDivisionGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiDivisionGetAPIResponse)
	},
}

// GetTaobaoAlitripTravelFscRouteApiDivisionGetAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelFscRouteApiDivisionGetAPIResponse
func GetTaobaoAlitripTravelFscRouteApiDivisionGetAPIResponse() *TaobaoAlitripTravelFscRouteApiDivisionGetAPIResponse {
	return poolTaobaoAlitripTravelFscRouteApiDivisionGetAPIResponse.Get().(*TaobaoAlitripTravelFscRouteApiDivisionGetAPIResponse)
}

// ReleaseTaobaoAlitripTravelFscRouteApiDivisionGetAPIResponse 将 TaobaoAlitripTravelFscRouteApiDivisionGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelFscRouteApiDivisionGetAPIResponse(v *TaobaoAlitripTravelFscRouteApiDivisionGetAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelFscRouteApiDivisionGetAPIResponse.Put(v)
}
