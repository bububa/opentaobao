package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelfscrouteapidivisiongetAPIResponse 获取国家城市信息 API返回值
// taobao.alitrip.travel.fsc.route.api.division.get
//
// 获取国家城市信息
type TaobaoalitriptravelfscrouteapidivisiongetAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelfscrouteapidivisiongetAPIResponseModel
}

// TaobaoalitriptravelfscrouteapidivisiongetAPIResponseModel is 获取国家城市信息 成功返回结果
type TaobaoalitriptravelfscrouteapidivisiongetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_division_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoalitriptravelfscrouteapidivisiongetTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}
