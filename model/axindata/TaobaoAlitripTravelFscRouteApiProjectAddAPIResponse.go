package axindata

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiProjectAddAPIResponse 新增团期 API返回值
// taobao.alitrip.travel.fsc.route.api.project.add
//
// 新增团期
type TaobaoAlitripTravelFscRouteApiProjectAddAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelFscRouteApiProjectAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscRouteApiProjectAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelFscRouteApiProjectAddAPIResponseModel).Reset()
}

// TaobaoAlitripTravelFscRouteApiProjectAddAPIResponseModel is 新增团期 成功返回结果
type TaobaoAlitripTravelFscRouteApiProjectAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_project_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoAlitripTravelFscRouteApiProjectAddTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscRouteApiProjectAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TopResult = nil
}

var poolTaobaoAlitripTravelFscRouteApiProjectAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiProjectAddAPIResponse)
	},
}

// GetTaobaoAlitripTravelFscRouteApiProjectAddAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelFscRouteApiProjectAddAPIResponse
func GetTaobaoAlitripTravelFscRouteApiProjectAddAPIResponse() *TaobaoAlitripTravelFscRouteApiProjectAddAPIResponse {
	return poolTaobaoAlitripTravelFscRouteApiProjectAddAPIResponse.Get().(*TaobaoAlitripTravelFscRouteApiProjectAddAPIResponse)
}

// ReleaseTaobaoAlitripTravelFscRouteApiProjectAddAPIResponse 将 TaobaoAlitripTravelFscRouteApiProjectAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelFscRouteApiProjectAddAPIResponse(v *TaobaoAlitripTravelFscRouteApiProjectAddAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelFscRouteApiProjectAddAPIResponse.Put(v)
}
