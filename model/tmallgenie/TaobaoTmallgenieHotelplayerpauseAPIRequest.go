package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTmallgenieHotelplayerpauseAPIRequest
天猫精灵酒店播放暂停 API请求
taobao.tmallgenie.hotelplayerpause

酒店推送指令给天猫精灵停止播放音乐 */
type TaobaoTmallgenieHotelplayerpauseAPIRequest struct {
	model.Params
	// 房间号
	_roomNo string
	// 酒店ID
	_hotelId int64
}

// NewTaobaoTmallgenieHotelplayerpauseRequest 初始化TaobaoTmallgenieHotelplayerpauseAPIRequest对象
func NewTaobaoTmallgenieHotelplayerpauseRequest() *TaobaoTmallgenieHotelplayerpauseAPIRequest {
	return &TaobaoTmallgenieHotelplayerpauseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTmallgenieHotelplayerpauseAPIRequest) GetApiMethodName() string {
	return "taobao.tmallgenie.hotelplayerpause"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTmallgenieHotelplayerpauseAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RoomNo Setter
// 房间号
func (r *TaobaoTmallgenieHotelplayerpauseAPIRequest) SetRoomNo(_roomNo string) error {
	r._roomNo = _roomNo
	r.Set("room_no", _roomNo)
	return nil
}

// Get RoomNo Getter
func (r TaobaoTmallgenieHotelplayerpauseAPIRequest) GetRoomNo() string {
	return r._roomNo
}

// Set is HotelId Setter
// 酒店ID
func (r *TaobaoTmallgenieHotelplayerpauseAPIRequest) SetHotelId(_hotelId int64) error {
	r._hotelId = _hotelId
	r.Set("hotel_id", _hotelId)
	return nil
}

// Get HotelId Getter
func (r TaobaoTmallgenieHotelplayerpauseAPIRequest) GetHotelId() int64 {
	return r._hotelId
}
