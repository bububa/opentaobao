package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelfscrouteapiprojectcloseAPIResponse 关闭团期 API返回值
// taobao.alitrip.travel.fsc.route.api.project.close
//
// 关闭团期
type TaobaoalitriptravelfscrouteapiprojectcloseAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelfscrouteapiprojectcloseAPIResponseModel
}

// TaobaoalitriptravelfscrouteapiprojectcloseAPIResponseModel is 关闭团期 成功返回结果
type TaobaoalitriptravelfscrouteapiprojectcloseAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_project_close_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoalitriptravelfscrouteapiprojectcloseTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}
