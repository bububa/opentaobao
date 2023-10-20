package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelfscrouteapibusinessareagetAPIResponse 获取业务区域 API返回值
// taobao.alitrip.travel.fsc.route.api.business.area.get
//
// 获取业务区域
type TaobaoalitriptravelfscrouteapibusinessareagetAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelfscrouteapibusinessareagetAPIResponseModel
}

// TaobaoalitriptravelfscrouteapibusinessareagetAPIResponseModel is 获取业务区域 成功返回结果
type TaobaoalitriptravelfscrouteapibusinessareagetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_business_area_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoalitriptravelfscrouteapibusinessareagetTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}
