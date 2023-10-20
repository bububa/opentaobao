package axintrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelticketProductDetailAPIResponse 阿信酒景套餐产品详情查询 API返回值
// taobao.alitrip.travel.axin.hotelticket.product.detail
//
// 阿信酒景套餐产品详情查询
type TaobaoAlitripTravelAxinHotelticketProductDetailAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelAxinHotelticketProductDetailAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelticketProductDetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelAxinHotelticketProductDetailAPIResponseModel).Reset()
}

// TaobaoAlitripTravelAxinHotelticketProductDetailAPIResponseModel is 阿信酒景套餐产品详情查询 成功返回结果
type TaobaoAlitripTravelAxinHotelticketProductDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotelticket_product_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *BaseResultApiDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelticketProductDetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripTravelAxinHotelticketProductDetailAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelAxinHotelticketProductDetailAPIResponse)
	},
}

// GetTaobaoAlitripTravelAxinHotelticketProductDetailAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelAxinHotelticketProductDetailAPIResponse
func GetTaobaoAlitripTravelAxinHotelticketProductDetailAPIResponse() *TaobaoAlitripTravelAxinHotelticketProductDetailAPIResponse {
	return poolTaobaoAlitripTravelAxinHotelticketProductDetailAPIResponse.Get().(*TaobaoAlitripTravelAxinHotelticketProductDetailAPIResponse)
}

// ReleaseTaobaoAlitripTravelAxinHotelticketProductDetailAPIResponse 将 TaobaoAlitripTravelAxinHotelticketProductDetailAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelAxinHotelticketProductDetailAPIResponse(v *TaobaoAlitripTravelAxinHotelticketProductDetailAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelAxinHotelticketProductDetailAPIResponse.Put(v)
}
