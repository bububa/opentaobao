package axintrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIResponse 阿信度假业务申请退款 API返回值
// taobao.alitrip.travel.axin.hotelticket.refund.orderrefund
//
// 阿信度假业务申请退款
type TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIResponseModel).Reset()
}

// TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIResponseModel is 阿信度假业务申请退款 成功返回结果
type TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotelticket_refund_orderrefund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果返回类
	Result *BaseResultApiDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIResponse)
	},
}

// GetTaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIResponse
func GetTaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIResponse() *TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIResponse {
	return poolTaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIResponse.Get().(*TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIResponse)
}

// ReleaseTaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIResponse 将 TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIResponse(v *TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIResponse.Put(v)
}
