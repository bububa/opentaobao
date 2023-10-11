package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiProductOfflineAPIResponse 产品下线 API返回值
// taobao.alitrip.travel.fsc.route.api.product.offline
//
// 产品下线
type TaobaoAlitripTravelFscRouteApiProductOfflineAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelFscRouteApiProductOfflineAPIResponseModel
}

// TaobaoAlitripTravelFscRouteApiProductOfflineAPIResponseModel is 产品下线 成功返回结果
type TaobaoAlitripTravelFscRouteApiProductOfflineAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_product_offline_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoAlitripTravelFscRouteApiProductOfflineTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}
