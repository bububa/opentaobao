package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiProjectCloseAPIResponse 关闭团期 API返回值
// taobao.alitrip.travel.fsc.route.api.project.close
//
// 关闭团期
type TaobaoAlitripTravelFscRouteApiProjectCloseAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelFscRouteApiProjectCloseAPIResponseModel
}

// TaobaoAlitripTravelFscRouteApiProjectCloseAPIResponseModel is 关闭团期 成功返回结果
type TaobaoAlitripTravelFscRouteApiProjectCloseAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_project_close_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoAlitripTravelFscRouteApiProjectCloseTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}
