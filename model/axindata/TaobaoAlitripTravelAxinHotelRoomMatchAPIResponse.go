package axindata

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelRoomMatchAPIResponse 阿信酒店房型匹配 API返回值
// taobao.alitrip.travel.axin.hotel.room.match
//
// 阿信酒店房型匹配
type TaobaoAlitripTravelAxinHotelRoomMatchAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelAxinHotelRoomMatchAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelRoomMatchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelAxinHotelRoomMatchAPIResponseModel).Reset()
}

// TaobaoAlitripTravelAxinHotelRoomMatchAPIResponseModel is 阿信酒店房型匹配 成功返回结果
type TaobaoAlitripTravelAxinHotelRoomMatchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotel_room_match_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *BaseResultApiDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelRoomMatchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripTravelAxinHotelRoomMatchAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelAxinHotelRoomMatchAPIResponse)
	},
}

// GetTaobaoAlitripTravelAxinHotelRoomMatchAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelAxinHotelRoomMatchAPIResponse
func GetTaobaoAlitripTravelAxinHotelRoomMatchAPIResponse() *TaobaoAlitripTravelAxinHotelRoomMatchAPIResponse {
	return poolTaobaoAlitripTravelAxinHotelRoomMatchAPIResponse.Get().(*TaobaoAlitripTravelAxinHotelRoomMatchAPIResponse)
}

// ReleaseTaobaoAlitripTravelAxinHotelRoomMatchAPIResponse 将 TaobaoAlitripTravelAxinHotelRoomMatchAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelAxinHotelRoomMatchAPIResponse(v *TaobaoAlitripTravelAxinHotelRoomMatchAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelAxinHotelRoomMatchAPIResponse.Put(v)
}
