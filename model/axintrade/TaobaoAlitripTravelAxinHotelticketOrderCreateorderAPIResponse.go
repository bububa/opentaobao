package axintrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIResponse 阿信度假业务创单并支付接口 API返回值
// taobao.alitrip.travel.axin.hotelticket.order.createorder
//
// 阿信度假业务创单并支付接口
type TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIResponseModel).Reset()
}

// TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIResponseModel is 阿信度假业务创单并支付接口 成功返回结果
type TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotelticket_order_createorder_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果返回类
	Result *BaseResultApiDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIResponse)
	},
}

// GetTaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIResponse
func GetTaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIResponse() *TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIResponse {
	return poolTaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIResponse.Get().(*TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIResponse)
}

// ReleaseTaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIResponse 将 TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIResponse(v *TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIResponse.Put(v)
}
