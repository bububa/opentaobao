package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiProductOnlineAPIResponse 产品上线 API返回值
// taobao.alitrip.travel.fsc.route.api.product.online
//
// 产品上线
type TaobaoAlitripTravelFscRouteApiProductOnlineAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelFscRouteApiProductOnlineAPIResponseModel
}

// TaobaoAlitripTravelFscRouteApiProductOnlineAPIResponseModel is 产品上线 成功返回结果
type TaobaoAlitripTravelFscRouteApiProductOnlineAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_product_online_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoAlitripTravelFscRouteApiProductOnlineTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}
