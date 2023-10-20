package axindata

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscRouteApiProductLabelGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelFscRouteApiProductLabelGetAPIResponseModel).Reset()
}

// TaobaoAlitripTravelFscRouteApiProductLabelGetAPIResponseModel is 获取线路主题 成功返回结果
type TaobaoAlitripTravelFscRouteApiProductLabelGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_product_label_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoAlitripTravelFscRouteApiProductLabelGetTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscRouteApiProductLabelGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TopResult = nil
}

var poolTaobaoAlitripTravelFscRouteApiProductLabelGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiProductLabelGetAPIResponse)
	},
}

// GetTaobaoAlitripTravelFscRouteApiProductLabelGetAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelFscRouteApiProductLabelGetAPIResponse
func GetTaobaoAlitripTravelFscRouteApiProductLabelGetAPIResponse() *TaobaoAlitripTravelFscRouteApiProductLabelGetAPIResponse {
	return poolTaobaoAlitripTravelFscRouteApiProductLabelGetAPIResponse.Get().(*TaobaoAlitripTravelFscRouteApiProductLabelGetAPIResponse)
}

// ReleaseTaobaoAlitripTravelFscRouteApiProductLabelGetAPIResponse 将 TaobaoAlitripTravelFscRouteApiProductLabelGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelFscRouteApiProductLabelGetAPIResponse(v *TaobaoAlitripTravelFscRouteApiProductLabelGetAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelFscRouteApiProductLabelGetAPIResponse.Put(v)
}
