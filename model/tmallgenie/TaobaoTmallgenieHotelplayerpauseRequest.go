package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵酒店播放暂停 APIRequest
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

func NewTaobaoTmallgenieHotelplayerpauseRequest() *TaobaoTmallgenieHotelplayerpauseRequest{
    return &TaobaoTmallgenieHotelplayerpauseRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTmallgenieHotelplayerpauseRequest) GetApiMethodName() string {
    return "taobao.tmallgenie.hotelplayerpause"
}

func (r TaobaoTmallgenieHotelplayerpauseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTmallgenieHotelplayerpauseRequest) SetRoomNo(roomNo string) error {
    r.roomNo = roomNo
    r.Set("room_no", roomNo)
    return nil
}

func (r TaobaoTmallgenieHotelplayerpauseRequest) GetRoomNo() string {
    return r.roomNo
}

func (r *TaobaoTmallgenieHotelplayerpauseRequest) SetHotelId(hotelId int64) error {
    r.hotelId = hotelId
    r.Set("hotel_id", hotelId)
    return nil
}

func (r TaobaoTmallgenieHotelplayerpauseRequest) GetHotelId() int64 {
    return r.hotelId
}

