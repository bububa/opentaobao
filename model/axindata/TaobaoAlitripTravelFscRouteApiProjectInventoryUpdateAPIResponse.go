package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelfscrouteapiprojectinventoryupdateAPIResponse 更新团期库存 API返回值
// taobao.alitrip.travel.fsc.route.api.project.inventory.update
//
// 更新团期库存
type TaobaoalitriptravelfscrouteapiprojectinventoryupdateAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelfscrouteapiprojectinventoryupdateAPIResponseModel
}

// TaobaoalitriptravelfscrouteapiprojectinventoryupdateAPIResponseModel is 更新团期库存 成功返回结果
type TaobaoalitriptravelfscrouteapiprojectinventoryupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_project_inventory_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoalitriptravelfscrouteapiprojectinventoryupdateTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}
