package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiProductLabelGetAPIResponse 获取线路主题 API返回值
// taobao.alitrip.travel.fsc.route.api.product.label.get
//
// 获取线路主题
type TaobaoAlitripTravelFscRouteApiProductLabelGetAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelFscRouteApiProductLabelGetAPIResponseModel
}

// TaobaoAlitripTravelFscRouteApiProductLabelGetAPIResponseModel is 获取线路主题 成功返回结果
type TaobaoAlitripTravelFscRouteApiProductLabelGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_product_label_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoAlitripTravelFscRouteApiProductLabelGetTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}
