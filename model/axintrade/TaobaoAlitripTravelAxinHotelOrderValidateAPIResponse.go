package axintrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelOrderValidateAPIResponse 阿信酒店订单数据校验 API返回值
// taobao.alitrip.travel.axin.hotel.order.validate
//
// 阿信酒店订单下单前的数据校验，包括酒店、房型、售卖政策、入离日期等参数的校验
type TaobaoAlitripTravelAxinHotelOrderValidateAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelAxinHotelOrderValidateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelOrderValidateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelAxinHotelOrderValidateAPIResponseModel).Reset()
}

// TaobaoAlitripTravelAxinHotelOrderValidateAPIResponseModel is 阿信酒店订单数据校验 成功返回结果
type TaobaoAlitripTravelAxinHotelOrderValidateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotel_order_validate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *BaseResultApiDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelOrderValidateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripTravelAxinHotelOrderValidateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelAxinHotelOrderValidateAPIResponse)
	},
}

// GetTaobaoAlitripTravelAxinHotelOrderValidateAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelAxinHotelOrderValidateAPIResponse
func GetTaobaoAlitripTravelAxinHotelOrderValidateAPIResponse() *TaobaoAlitripTravelAxinHotelOrderValidateAPIResponse {
	return poolTaobaoAlitripTravelAxinHotelOrderValidateAPIResponse.Get().(*TaobaoAlitripTravelAxinHotelOrderValidateAPIResponse)
}

// ReleaseTaobaoAlitripTravelAxinHotelOrderValidateAPIResponse 将 TaobaoAlitripTravelAxinHotelOrderValidateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelAxinHotelOrderValidateAPIResponse(v *TaobaoAlitripTravelAxinHotelOrderValidateAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelAxinHotelOrderValidateAPIResponse.Put(v)
}
