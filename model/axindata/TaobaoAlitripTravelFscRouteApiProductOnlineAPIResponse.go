package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelfscrouteapiproductonlineAPIResponse 产品上线 API返回值
// taobao.alitrip.travel.fsc.route.api.product.online
//
// 产品上线
type TaobaoalitriptravelfscrouteapiproductonlineAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelfscrouteapiproductonlineAPIResponseModel
}

// TaobaoalitriptravelfscrouteapiproductonlineAPIResponseModel is 产品上线 成功返回结果
type TaobaoalitriptravelfscrouteapiproductonlineAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_product_online_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoalitriptravelfscrouteapiproductonlineTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}
