package axindata

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscRouteApiProductUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelFscRouteApiProductUpdateAPIResponseModel).Reset()
}

// TaobaoAlitripTravelFscRouteApiProductUpdateAPIResponseModel is 更新线路产品基本信息 成功返回结果
type TaobaoAlitripTravelFscRouteApiProductUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_product_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoAlitripTravelFscRouteApiProductUpdateTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscRouteApiProductUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TopResult = nil
}

var poolTaobaoAlitripTravelFscRouteApiProductUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiProductUpdateAPIResponse)
	},
}

// GetTaobaoAlitripTravelFscRouteApiProductUpdateAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelFscRouteApiProductUpdateAPIResponse
func GetTaobaoAlitripTravelFscRouteApiProductUpdateAPIResponse() *TaobaoAlitripTravelFscRouteApiProductUpdateAPIResponse {
	return poolTaobaoAlitripTravelFscRouteApiProductUpdateAPIResponse.Get().(*TaobaoAlitripTravelFscRouteApiProductUpdateAPIResponse)
}

// ReleaseTaobaoAlitripTravelFscRouteApiProductUpdateAPIResponse 将 TaobaoAlitripTravelFscRouteApiProductUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelFscRouteApiProductUpdateAPIResponse(v *TaobaoAlitripTravelFscRouteApiProductUpdateAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelFscRouteApiProductUpdateAPIResponse.Put(v)
}
