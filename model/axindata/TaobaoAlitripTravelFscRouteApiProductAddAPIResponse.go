package axindata

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscRouteApiProductAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelFscRouteApiProductAddAPIResponseModel).Reset()
}

// TaobaoAlitripTravelFscRouteApiProductAddAPIResponseModel is 新增线路产品基本信息 成功返回结果
type TaobaoAlitripTravelFscRouteApiProductAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_product_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoAlitripTravelFscRouteApiProductAddTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscRouteApiProductAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TopResult = nil
}

var poolTaobaoAlitripTravelFscRouteApiProductAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiProductAddAPIResponse)
	},
}

// GetTaobaoAlitripTravelFscRouteApiProductAddAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelFscRouteApiProductAddAPIResponse
func GetTaobaoAlitripTravelFscRouteApiProductAddAPIResponse() *TaobaoAlitripTravelFscRouteApiProductAddAPIResponse {
	return poolTaobaoAlitripTravelFscRouteApiProductAddAPIResponse.Get().(*TaobaoAlitripTravelFscRouteApiProductAddAPIResponse)
}

// ReleaseTaobaoAlitripTravelFscRouteApiProductAddAPIResponse 将 TaobaoAlitripTravelFscRouteApiProductAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelFscRouteApiProductAddAPIResponse(v *TaobaoAlitripTravelFscRouteApiProductAddAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelFscRouteApiProductAddAPIResponse.Put(v)
}
