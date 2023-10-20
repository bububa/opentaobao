package axindata

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIResponse 获取业务区域 API返回值
// taobao.alitrip.travel.fsc.route.api.business.area.get
//
// 获取业务区域
type TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIResponseModel).Reset()
}

// TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIResponseModel is 获取业务区域 成功返回结果
type TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_business_area_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoAlitripTravelFscRouteApiBusinessAreaGetTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TopResult = nil
}

var poolTaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIResponse)
	},
}

// GetTaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIResponse
func GetTaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIResponse() *TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIResponse {
	return poolTaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIResponse.Get().(*TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIResponse)
}

// ReleaseTaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIResponse 将 TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIResponse(v *TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIResponse.Put(v)
}
