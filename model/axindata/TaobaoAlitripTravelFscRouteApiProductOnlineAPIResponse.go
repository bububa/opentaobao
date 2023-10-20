package axindata

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscRouteApiProductOnlineAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelFscRouteApiProductOnlineAPIResponseModel).Reset()
}

// TaobaoAlitripTravelFscRouteApiProductOnlineAPIResponseModel is 产品上线 成功返回结果
type TaobaoAlitripTravelFscRouteApiProductOnlineAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_product_online_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoAlitripTravelFscRouteApiProductOnlineTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscRouteApiProductOnlineAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TopResult = nil
}

var poolTaobaoAlitripTravelFscRouteApiProductOnlineAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiProductOnlineAPIResponse)
	},
}

// GetTaobaoAlitripTravelFscRouteApiProductOnlineAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelFscRouteApiProductOnlineAPIResponse
func GetTaobaoAlitripTravelFscRouteApiProductOnlineAPIResponse() *TaobaoAlitripTravelFscRouteApiProductOnlineAPIResponse {
	return poolTaobaoAlitripTravelFscRouteApiProductOnlineAPIResponse.Get().(*TaobaoAlitripTravelFscRouteApiProductOnlineAPIResponse)
}

// ReleaseTaobaoAlitripTravelFscRouteApiProductOnlineAPIResponse 将 TaobaoAlitripTravelFscRouteApiProductOnlineAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelFscRouteApiProductOnlineAPIResponse(v *TaobaoAlitripTravelFscRouteApiProductOnlineAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelFscRouteApiProductOnlineAPIResponse.Put(v)
}
