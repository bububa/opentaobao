package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotmallgeniehotelplayerpauseAPIRequest 天猫精灵酒店播放暂停 API请求
// taobao.tmallgenie.hotelplayerpause
//
// 酒店推送指令给天猫精灵停止播放音乐
type TaobaotmallgeniehotelplayerpauseAPIRequest struct {
	model.Params
	// 房间号
	_roomNo string
	// 酒店ID
	_hotelId int64
}

// NewTaobaotmallgeniehotelplayerpauseRequest 初始化TaobaotmallgeniehotelplayerpauseAPIRequest对象
func NewTaobaotmallgeniehotelplayerpauseRequest() *TaobaotmallgeniehotelplayerpauseAPIRequest {
	return &TaobaotmallgeniehotelplayerpauseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotmallgeniehotelplayerpauseAPIRequest) GetApiMethodName() string {
	return "taobao.tmallgenie.hotelplayerpause"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotmallgeniehotelplayerpauseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotmallgeniehotelplayerpauseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRoomNo is RoomNo Setter
// 房间号
func (r *TaobaotmallgeniehotelplayerpauseAPIRequest) SetRoomNo(_roomNo string) error {
	r._roomNo = _roomNo
	r.Set("room_no", _roomNo)
	return nil
}

// GetRoomNo RoomNo Getter
func (r TaobaotmallgeniehotelplayerpauseAPIRequest) GetRoomNo() string {
	return r._roomNo
}

// SetHotelId is HotelId Setter
// 酒店ID
func (r *TaobaotmallgeniehotelplayerpauseAPIRequest) SetHotelId(_hotelId int64) error {
	r._hotelId = _hotelId
	r.Set("hotel_id", _hotelId)
	return nil
}

// GetHotelId HotelId Getter
func (r TaobaotmallgeniehotelplayerpauseAPIRequest) GetHotelId() int64 {
	return r._hotelId
}
