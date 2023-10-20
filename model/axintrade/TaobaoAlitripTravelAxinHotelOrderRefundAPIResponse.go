package axintrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelOrderRefundAPIResponse 阿信酒店分销订单退款API API返回值
// taobao.alitrip.travel.axin.hotel.order.refund
//
// 阿信酒店分销订单退款
type TaobaoAlitripTravelAxinHotelOrderRefundAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelAxinHotelOrderRefundAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelOrderRefundAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelAxinHotelOrderRefundAPIResponseModel).Reset()
}

// TaobaoAlitripTravelAxinHotelOrderRefundAPIResponseModel is 阿信酒店分销订单退款API 成功返回结果
type TaobaoAlitripTravelAxinHotelOrderRefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotel_order_refund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *BaseResultApiDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelOrderRefundAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripTravelAxinHotelOrderRefundAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelAxinHotelOrderRefundAPIResponse)
	},
}

// GetTaobaoAlitripTravelAxinHotelOrderRefundAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelAxinHotelOrderRefundAPIResponse
func GetTaobaoAlitripTravelAxinHotelOrderRefundAPIResponse() *TaobaoAlitripTravelAxinHotelOrderRefundAPIResponse {
	return poolTaobaoAlitripTravelAxinHotelOrderRefundAPIResponse.Get().(*TaobaoAlitripTravelAxinHotelOrderRefundAPIResponse)
}

// ReleaseTaobaoAlitripTravelAxinHotelOrderRefundAPIResponse 将 TaobaoAlitripTravelAxinHotelOrderRefundAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelAxinHotelOrderRefundAPIResponse(v *TaobaoAlitripTravelAxinHotelOrderRefundAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelAxinHotelOrderRefundAPIResponse.Put(v)
}
