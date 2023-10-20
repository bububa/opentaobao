package axintrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponse 阿信度假交易订单查询接口 API返回值
// taobao.alitrip.travel.axin.hotelticket.order.query
//
// 阿信度假交易订单查询接口
type TaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponseModel).Reset()
}

// TaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponseModel is 阿信度假交易订单查询接口 成功返回结果
type TaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotelticket_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *BaseResultApiDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponse)
	},
}

// GetTaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponse
func GetTaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponse() *TaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponse {
	return poolTaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponse.Get().(*TaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponse)
}

// ReleaseTaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponse 将 TaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponse(v *TaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponse.Put(v)
}
