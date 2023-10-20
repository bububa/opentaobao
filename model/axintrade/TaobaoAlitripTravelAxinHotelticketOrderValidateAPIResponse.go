package axintrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelticketOrderValidateAPIResponse 阿信度假业务交易试单接口 API返回值
// taobao.alitrip.travel.axin.hotelticket.order.validate
//
// 阿信度假业务交易试单接口
type TaobaoAlitripTravelAxinHotelticketOrderValidateAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelAxinHotelticketOrderValidateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelticketOrderValidateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelAxinHotelticketOrderValidateAPIResponseModel).Reset()
}

// TaobaoAlitripTravelAxinHotelticketOrderValidateAPIResponseModel is 阿信度假业务交易试单接口 成功返回结果
type TaobaoAlitripTravelAxinHotelticketOrderValidateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotelticket_order_validate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果返回类
	Result *BaseResultApiDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelticketOrderValidateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripTravelAxinHotelticketOrderValidateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelAxinHotelticketOrderValidateAPIResponse)
	},
}

// GetTaobaoAlitripTravelAxinHotelticketOrderValidateAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelAxinHotelticketOrderValidateAPIResponse
func GetTaobaoAlitripTravelAxinHotelticketOrderValidateAPIResponse() *TaobaoAlitripTravelAxinHotelticketOrderValidateAPIResponse {
	return poolTaobaoAlitripTravelAxinHotelticketOrderValidateAPIResponse.Get().(*TaobaoAlitripTravelAxinHotelticketOrderValidateAPIResponse)
}

// ReleaseTaobaoAlitripTravelAxinHotelticketOrderValidateAPIResponse 将 TaobaoAlitripTravelAxinHotelticketOrderValidateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelAxinHotelticketOrderValidateAPIResponse(v *TaobaoAlitripTravelAxinHotelticketOrderValidateAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelAxinHotelticketOrderValidateAPIResponse.Put(v)
}
