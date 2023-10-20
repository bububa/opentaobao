package axindata

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelPriceGetAPIResponse 酒店报价服务-阿信 API返回值
// taobao.alitrip.travel.axin.hotel.price.get
//
// 酒店报价查询服务
type TaobaoAlitripTravelAxinHotelPriceGetAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelAxinHotelPriceGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelPriceGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelAxinHotelPriceGetAPIResponseModel).Reset()
}

// TaobaoAlitripTravelAxinHotelPriceGetAPIResponseModel is 酒店报价服务-阿信 成功返回结果
type TaobaoAlitripTravelAxinHotelPriceGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotel_price_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoAlitripTravelAxinHotelPriceGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelPriceGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripTravelAxinHotelPriceGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelAxinHotelPriceGetAPIResponse)
	},
}

// GetTaobaoAlitripTravelAxinHotelPriceGetAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelAxinHotelPriceGetAPIResponse
func GetTaobaoAlitripTravelAxinHotelPriceGetAPIResponse() *TaobaoAlitripTravelAxinHotelPriceGetAPIResponse {
	return poolTaobaoAlitripTravelAxinHotelPriceGetAPIResponse.Get().(*TaobaoAlitripTravelAxinHotelPriceGetAPIResponse)
}

// ReleaseTaobaoAlitripTravelAxinHotelPriceGetAPIResponse 将 TaobaoAlitripTravelAxinHotelPriceGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelAxinHotelPriceGetAPIResponse(v *TaobaoAlitripTravelAxinHotelPriceGetAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelAxinHotelPriceGetAPIResponse.Put(v)
}
