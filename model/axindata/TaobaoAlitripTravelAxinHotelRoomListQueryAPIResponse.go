package axindata

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelRoomListQueryAPIResponse 阿信酒店分销-标准酒店房型列表查询 API返回值
// taobao.alitrip.travel.axin.hotel.room.list.query
//
// 标准酒店房型列表查询
type TaobaoAlitripTravelAxinHotelRoomListQueryAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelAxinHotelRoomListQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelRoomListQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelAxinHotelRoomListQueryAPIResponseModel).Reset()
}

// TaobaoAlitripTravelAxinHotelRoomListQueryAPIResponseModel is 阿信酒店分销-标准酒店房型列表查询 成功返回结果
type TaobaoAlitripTravelAxinHotelRoomListQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotel_room_list_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回模型
	Result *BaseResultApiDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelRoomListQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripTravelAxinHotelRoomListQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelAxinHotelRoomListQueryAPIResponse)
	},
}

// GetTaobaoAlitripTravelAxinHotelRoomListQueryAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelAxinHotelRoomListQueryAPIResponse
func GetTaobaoAlitripTravelAxinHotelRoomListQueryAPIResponse() *TaobaoAlitripTravelAxinHotelRoomListQueryAPIResponse {
	return poolTaobaoAlitripTravelAxinHotelRoomListQueryAPIResponse.Get().(*TaobaoAlitripTravelAxinHotelRoomListQueryAPIResponse)
}

// ReleaseTaobaoAlitripTravelAxinHotelRoomListQueryAPIResponse 将 TaobaoAlitripTravelAxinHotelRoomListQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelAxinHotelRoomListQueryAPIResponse(v *TaobaoAlitripTravelAxinHotelRoomListQueryAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelAxinHotelRoomListQueryAPIResponse.Put(v)
}
