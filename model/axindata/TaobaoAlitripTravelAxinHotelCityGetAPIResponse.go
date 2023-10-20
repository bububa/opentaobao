package axindata

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelCityGetAPIResponse 城市列表信息查询-阿信 API返回值
// taobao.alitrip.travel.axin.hotel.city.get
//
// 阿信城市列表信息查询
type TaobaoAlitripTravelAxinHotelCityGetAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelAxinHotelCityGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelCityGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelAxinHotelCityGetAPIResponseModel).Reset()
}

// TaobaoAlitripTravelAxinHotelCityGetAPIResponseModel is 城市列表信息查询-阿信 成功返回结果
type TaobaoAlitripTravelAxinHotelCityGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotel_city_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoAlitripTravelAxinHotelCityGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelCityGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripTravelAxinHotelCityGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelAxinHotelCityGetAPIResponse)
	},
}

// GetTaobaoAlitripTravelAxinHotelCityGetAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelAxinHotelCityGetAPIResponse
func GetTaobaoAlitripTravelAxinHotelCityGetAPIResponse() *TaobaoAlitripTravelAxinHotelCityGetAPIResponse {
	return poolTaobaoAlitripTravelAxinHotelCityGetAPIResponse.Get().(*TaobaoAlitripTravelAxinHotelCityGetAPIResponse)
}

// ReleaseTaobaoAlitripTravelAxinHotelCityGetAPIResponse 将 TaobaoAlitripTravelAxinHotelCityGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelAxinHotelCityGetAPIResponse(v *TaobaoAlitripTravelAxinHotelCityGetAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelAxinHotelCityGetAPIResponse.Put(v)
}
