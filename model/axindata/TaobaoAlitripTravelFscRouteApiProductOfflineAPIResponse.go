package axindata

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscRouteApiProductOfflineAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelFscRouteApiProductOfflineAPIResponseModel).Reset()
}

// TaobaoAlitripTravelFscRouteApiProductOfflineAPIResponseModel is 产品下线 成功返回结果
type TaobaoAlitripTravelFscRouteApiProductOfflineAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_product_offline_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoAlitripTravelFscRouteApiProductOfflineTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscRouteApiProductOfflineAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TopResult = nil
}

var poolTaobaoAlitripTravelFscRouteApiProductOfflineAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiProductOfflineAPIResponse)
	},
}

// GetTaobaoAlitripTravelFscRouteApiProductOfflineAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelFscRouteApiProductOfflineAPIResponse
func GetTaobaoAlitripTravelFscRouteApiProductOfflineAPIResponse() *TaobaoAlitripTravelFscRouteApiProductOfflineAPIResponse {
	return poolTaobaoAlitripTravelFscRouteApiProductOfflineAPIResponse.Get().(*TaobaoAlitripTravelFscRouteApiProductOfflineAPIResponse)
}

// ReleaseTaobaoAlitripTravelFscRouteApiProductOfflineAPIResponse 将 TaobaoAlitripTravelFscRouteApiProductOfflineAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelFscRouteApiProductOfflineAPIResponse(v *TaobaoAlitripTravelFscRouteApiProductOfflineAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelFscRouteApiProductOfflineAPIResponse.Put(v)
}
