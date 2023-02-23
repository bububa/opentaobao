package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelRoomMatchAPIRequest 阿信酒店房型匹配 API请求
// taobao.alitrip.travel.axin.hotel.room.match
//
// 阿信酒店房型匹配
type TaobaoAlitripTravelAxinHotelRoomMatchAPIRequest struct {
	model.Params
	// 参数
	_req *MatchedRoomRequestDto
}

// NewTaobaoAlitripTravelAxinHotelRoomMatchRequest 初始化TaobaoAlitripTravelAxinHotelRoomMatchAPIRequest对象
func NewTaobaoAlitripTravelAxinHotelRoomMatchRequest() *TaobaoAlitripTravelAxinHotelRoomMatchAPIRequest {
	return &TaobaoAlitripTravelAxinHotelRoomMatchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelAxinHotelRoomMatchAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.room.match"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelAxinHotelRoomMatchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelAxinHotelRoomMatchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReq is Req Setter
// 参数
func (r *TaobaoAlitripTravelAxinHotelRoomMatchAPIRequest) SetReq(_req *MatchedRoomRequestDto) error {
	r._req = _req
	r.Set("req", _req)
	return nil
}

// GetReq Req Getter
func (r TaobaoAlitripTravelAxinHotelRoomMatchAPIRequest) GetReq() *MatchedRoomRequestDto {
	return r._req
}
