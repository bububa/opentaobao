package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelfscrouteapiprojectupdateAPIResponse 更新团期 API返回值
// taobao.alitrip.travel.fsc.route.api.project.update
//
// 更新团期
type TaobaoalitriptravelfscrouteapiprojectupdateAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelfscrouteapiprojectupdateAPIResponseModel
}

// TaobaoalitriptravelfscrouteapiprojectupdateAPIResponseModel is 更新团期 成功返回结果
type TaobaoalitriptravelfscrouteapiprojectupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_project_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoalitriptravelfscrouteapiprojectupdateTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}
