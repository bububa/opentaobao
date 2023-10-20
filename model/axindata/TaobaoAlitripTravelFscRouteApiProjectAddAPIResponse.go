package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelfscrouteapiprojectaddAPIResponse 新增团期 API返回值
// taobao.alitrip.travel.fsc.route.api.project.add
//
// 新增团期
type TaobaoalitriptravelfscrouteapiprojectaddAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelfscrouteapiprojectaddAPIResponseModel
}

// TaobaoalitriptravelfscrouteapiprojectaddAPIResponseModel is 新增团期 成功返回结果
type TaobaoalitriptravelfscrouteapiprojectaddAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_project_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoalitriptravelfscrouteapiprojectaddTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}
