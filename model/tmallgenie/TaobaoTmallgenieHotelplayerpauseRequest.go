package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵酒店播放暂停 API请求
taobao.tmallgenie.hotelplayerpause

酒店推送指令给天猫精灵停止播放音乐
*/
type TaobaoTmallgenieHotelplayerpauseRequest struct {
    model.Params
    // 房间号
    roomNo   string
    // 酒店ID
    hotelId   int64
}

// 初始化TaobaoTmallgenieHotelplayerpauseRequest对象
func NewTaobaoTmallgenieHotelplayerpauseRequest() *TaobaoTmallgenieHotelplayerpauseRequest{
    return &TaobaoTmallgenieHotelplayerpauseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTmallgenieHotelplayerpauseRequest) GetApiMethodName() string {
    return "taobao.tmallgenie.hotelplayerpause"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTmallgenieHotelplayerpauseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RoomNo Setter
// 房间号
func (r *TaobaoTmallgenieHotelplayerpauseRequest) SetRoomNo(roomNo string) error {
    r.roomNo = roomNo
    r.Set("room_no", roomNo)
    return nil
}

// RoomNo Getter
func (r TaobaoTmallgenieHotelplayerpauseRequest) GetRoomNo() string {
    return r.roomNo
}
// HotelId Setter
// 酒店ID
func (r *TaobaoTmallgenieHotelplayerpauseRequest) SetHotelId(hotelId int64) error {
    r.hotelId = hotelId
    r.Set("hotel_id", hotelId)
    return nil
}

// HotelId Getter
func (r TaobaoTmallgenieHotelplayerpauseRequest) GetHotelId() int64 {
    return r.hotelId
}
