package axindata

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelMatchAPIResponse 酒店匹配接口-阿信 API返回值
// taobao.alitrip.travel.axin.hotel.match
//
// 酒店匹配接口-阿信
type TaobaoAlitripTravelAxinHotelMatchAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelAxinHotelMatchAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelMatchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelAxinHotelMatchAPIResponseModel).Reset()
}

// TaobaoAlitripTravelAxinHotelMatchAPIResponseModel is 酒店匹配接口-阿信 成功返回结果
type TaobaoAlitripTravelAxinHotelMatchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotel_match_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果result
	Result *BaseResultApiDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelMatchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripTravelAxinHotelMatchAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelAxinHotelMatchAPIResponse)
	},
}

// GetTaobaoAlitripTravelAxinHotelMatchAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelAxinHotelMatchAPIResponse
func GetTaobaoAlitripTravelAxinHotelMatchAPIResponse() *TaobaoAlitripTravelAxinHotelMatchAPIResponse {
	return poolTaobaoAlitripTravelAxinHotelMatchAPIResponse.Get().(*TaobaoAlitripTravelAxinHotelMatchAPIResponse)
}

// ReleaseTaobaoAlitripTravelAxinHotelMatchAPIResponse 将 TaobaoAlitripTravelAxinHotelMatchAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelAxinHotelMatchAPIResponse(v *TaobaoAlitripTravelAxinHotelMatchAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelAxinHotelMatchAPIResponse.Put(v)
}
