package axintrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelOrderDetailAPIResponse 阿信酒店分销-订单详情接口 API返回值
// taobao.alitrip.travel.axin.hotel.order.detail
//
// 阿信酒店订单详情的读取接口
type TaobaoAlitripTravelAxinHotelOrderDetailAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelAxinHotelOrderDetailAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelOrderDetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelAxinHotelOrderDetailAPIResponseModel).Reset()
}

// TaobaoAlitripTravelAxinHotelOrderDetailAPIResponseModel is 阿信酒店分销-订单详情接口 成功返回结果
type TaobaoAlitripTravelAxinHotelOrderDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotel_order_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoAlitripTravelAxinHotelOrderDetailResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelOrderDetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripTravelAxinHotelOrderDetailAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelAxinHotelOrderDetailAPIResponse)
	},
}

// GetTaobaoAlitripTravelAxinHotelOrderDetailAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelAxinHotelOrderDetailAPIResponse
func GetTaobaoAlitripTravelAxinHotelOrderDetailAPIResponse() *TaobaoAlitripTravelAxinHotelOrderDetailAPIResponse {
	return poolTaobaoAlitripTravelAxinHotelOrderDetailAPIResponse.Get().(*TaobaoAlitripTravelAxinHotelOrderDetailAPIResponse)
}

// ReleaseTaobaoAlitripTravelAxinHotelOrderDetailAPIResponse 将 TaobaoAlitripTravelAxinHotelOrderDetailAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelAxinHotelOrderDetailAPIResponse(v *TaobaoAlitripTravelAxinHotelOrderDetailAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelAxinHotelOrderDetailAPIResponse.Put(v)
}
