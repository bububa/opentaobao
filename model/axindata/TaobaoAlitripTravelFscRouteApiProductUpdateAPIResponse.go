package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiProductUpdateAPIResponse 更新线路产品基本信息 API返回值
// taobao.alitrip.travel.fsc.route.api.product.update
//
// 更新线路产品基本信息
type TaobaoAlitripTravelFscRouteApiProductUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelFscRouteApiProductUpdateAPIResponseModel
}

// TaobaoAlitripTravelFscRouteApiProductUpdateAPIResponseModel is 更新线路产品基本信息 成功返回结果
type TaobaoAlitripTravelFscRouteApiProductUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_product_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoAlitripTravelFscRouteApiProductUpdateTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}
