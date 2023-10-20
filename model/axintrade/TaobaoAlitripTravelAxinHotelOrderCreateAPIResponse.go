package axintrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelOrderCreateAPIResponse 酒店分销订单创建服务-阿信 API返回值
// taobao.alitrip.travel.axin.hotel.order.create
//
// 提供酒店分销订单创建服务
type TaobaoAlitripTravelAxinHotelOrderCreateAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelAxinHotelOrderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelOrderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelAxinHotelOrderCreateAPIResponseModel).Reset()
}

// TaobaoAlitripTravelAxinHotelOrderCreateAPIResponseModel is 酒店分销订单创建服务-阿信 成功返回结果
type TaobaoAlitripTravelAxinHotelOrderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotel_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoAlitripTravelAxinHotelOrderCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelOrderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripTravelAxinHotelOrderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelAxinHotelOrderCreateAPIResponse)
	},
}

// GetTaobaoAlitripTravelAxinHotelOrderCreateAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelAxinHotelOrderCreateAPIResponse
func GetTaobaoAlitripTravelAxinHotelOrderCreateAPIResponse() *TaobaoAlitripTravelAxinHotelOrderCreateAPIResponse {
	return poolTaobaoAlitripTravelAxinHotelOrderCreateAPIResponse.Get().(*TaobaoAlitripTravelAxinHotelOrderCreateAPIResponse)
}

// ReleaseTaobaoAlitripTravelAxinHotelOrderCreateAPIResponse 将 TaobaoAlitripTravelAxinHotelOrderCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelAxinHotelOrderCreateAPIResponse(v *TaobaoAlitripTravelAxinHotelOrderCreateAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelAxinHotelOrderCreateAPIResponse.Put(v)
}
