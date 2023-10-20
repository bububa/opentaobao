package axindata

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiDivisionApplyAPIResponse 线路供应商提交新增城市申请 API返回值
// taobao.alitrip.travel.fsc.route.api.division.apply
//
// 线路供应商提交新增城市申请
type TaobaoAlitripTravelFscRouteApiDivisionApplyAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelFscRouteApiDivisionApplyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscRouteApiDivisionApplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelFscRouteApiDivisionApplyAPIResponseModel).Reset()
}

// TaobaoAlitripTravelFscRouteApiDivisionApplyAPIResponseModel is 线路供应商提交新增城市申请 成功返回结果
type TaobaoAlitripTravelFscRouteApiDivisionApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_division_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoAlitripTravelFscRouteApiDivisionApplyTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscRouteApiDivisionApplyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TopResult = nil
}

var poolTaobaoAlitripTravelFscRouteApiDivisionApplyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiDivisionApplyAPIResponse)
	},
}

// GetTaobaoAlitripTravelFscRouteApiDivisionApplyAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelFscRouteApiDivisionApplyAPIResponse
func GetTaobaoAlitripTravelFscRouteApiDivisionApplyAPIResponse() *TaobaoAlitripTravelFscRouteApiDivisionApplyAPIResponse {
	return poolTaobaoAlitripTravelFscRouteApiDivisionApplyAPIResponse.Get().(*TaobaoAlitripTravelFscRouteApiDivisionApplyAPIResponse)
}

// ReleaseTaobaoAlitripTravelFscRouteApiDivisionApplyAPIResponse 将 TaobaoAlitripTravelFscRouteApiDivisionApplyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelFscRouteApiDivisionApplyAPIResponse(v *TaobaoAlitripTravelFscRouteApiDivisionApplyAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelFscRouteApiDivisionApplyAPIResponse.Put(v)
}
