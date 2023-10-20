package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelfscrouteapiprojectopenAPIResponse 打开团期 API返回值
// taobao.alitrip.travel.fsc.route.api.project.open
//
// 打开团期
type TaobaoalitriptravelfscrouteapiprojectopenAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelfscrouteapiprojectopenAPIResponseModel
}

// TaobaoalitriptravelfscrouteapiprojectopenAPIResponseModel is 打开团期 成功返回结果
type TaobaoalitriptravelfscrouteapiprojectopenAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_project_open_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoalitriptravelfscrouteapiprojectopenTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}
