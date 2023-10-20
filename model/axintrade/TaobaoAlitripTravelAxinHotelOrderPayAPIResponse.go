package axintrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelOrderPayAPIResponse 阿信酒店分销订单支付 API返回值
// taobao.alitrip.travel.axin.hotel.order.pay
//
// 阿信酒店分销订单支付
type TaobaoAlitripTravelAxinHotelOrderPayAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelAxinHotelOrderPayAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelOrderPayAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelAxinHotelOrderPayAPIResponseModel).Reset()
}

// TaobaoAlitripTravelAxinHotelOrderPayAPIResponseModel is 阿信酒店分销订单支付 成功返回结果
type TaobaoAlitripTravelAxinHotelOrderPayAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotel_order_pay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoAlitripTravelAxinHotelOrderPayResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelOrderPayAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripTravelAxinHotelOrderPayAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelAxinHotelOrderPayAPIResponse)
	},
}

// GetTaobaoAlitripTravelAxinHotelOrderPayAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelAxinHotelOrderPayAPIResponse
func GetTaobaoAlitripTravelAxinHotelOrderPayAPIResponse() *TaobaoAlitripTravelAxinHotelOrderPayAPIResponse {
	return poolTaobaoAlitripTravelAxinHotelOrderPayAPIResponse.Get().(*TaobaoAlitripTravelAxinHotelOrderPayAPIResponse)
}

// ReleaseTaobaoAlitripTravelAxinHotelOrderPayAPIResponse 将 TaobaoAlitripTravelAxinHotelOrderPayAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelAxinHotelOrderPayAPIResponse(v *TaobaoAlitripTravelAxinHotelOrderPayAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelAxinHotelOrderPayAPIResponse.Put(v)
}
