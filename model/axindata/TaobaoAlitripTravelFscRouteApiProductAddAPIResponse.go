package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiProductAddAPIResponse 新增线路产品基本信息 API返回值
// taobao.alitrip.travel.fsc.route.api.product.add
//
// 新增线路产品基本信息
type TaobaoAlitripTravelFscRouteApiProductAddAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelFscRouteApiProductAddAPIResponseModel
}

// TaobaoAlitripTravelFscRouteApiProductAddAPIResponseModel is 新增线路产品基本信息 成功返回结果
type TaobaoAlitripTravelFscRouteApiProductAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_product_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoAlitripTravelFscRouteApiProductAddTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}
