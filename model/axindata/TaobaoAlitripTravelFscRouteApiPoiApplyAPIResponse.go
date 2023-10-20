package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelfscrouteapipoiapplyAPIResponse 线路供应商提交新增POI申请 API返回值
// taobao.alitrip.travel.fsc.route.api.poi.apply
//
// 线路供应商提交新增POI申请
type TaobaoalitriptravelfscrouteapipoiapplyAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelfscrouteapipoiapplyAPIResponseModel
}

// TaobaoalitriptravelfscrouteapipoiapplyAPIResponseModel is 线路供应商提交新增POI申请 成功返回结果
type TaobaoalitriptravelfscrouteapipoiapplyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_poi_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoalitriptravelfscrouteapipoiapplyTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}
