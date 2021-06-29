package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店欢迎词推送 API请求
taobao.tmallgenie.hotelwelcome

推送欢迎词，让天猫精灵播放
*/
type TaobaoTmallgenieHotelwelcomeRequest struct {
    model.Params
    // 房间号
    _roomNo   string
    // 酒店ID
    _hotelId   int64
    // 模板ID
    _templateId   string
    // 模板变量
    _templateVariable   string
}

// 初始化TaobaoTmallgenieHotelwelcomeRequest对象
func NewTaobaoTmallgenieHotelwelcomeRequest() *TaobaoTmallgenieHotelwelcomeRequest{
    return &TaobaoTmallgenieHotelwelcomeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTmallgenieHotelwelcomeRequest) GetApiMethodName() string {
    return "taobao.tmallgenie.hotelwelcome"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTmallgenieHotelwelcomeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RoomNo Setter
// 房间号
func (r *TaobaoTmallgenieHotelwelcomeRequest) SetRoomNo(_roomNo string) error {
    r._roomNo = _roomNo
    r.Set("room_no", _roomNo)
    return nil
}

// RoomNo Getter
func (r TaobaoTmallgenieHotelwelcomeRequest) GetRoomNo() string {
    return r._roomNo
}
// HotelId Setter
// 酒店ID
func (r *TaobaoTmallgenieHotelwelcomeRequest) SetHotelId(_hotelId int64) error {
    r._hotelId = _hotelId
    r.Set("hotel_id", _hotelId)
    return nil
}

// HotelId Getter
func (r TaobaoTmallgenieHotelwelcomeRequest) GetHotelId() int64 {
    return r._hotelId
}
// TemplateId Setter
// 模板ID
func (r *TaobaoTmallgenieHotelwelcomeRequest) SetTemplateId(_templateId string) error {
    r._templateId = _templateId
    r.Set("template_id", _templateId)
    return nil
}

// TemplateId Getter
func (r TaobaoTmallgenieHotelwelcomeRequest) GetTemplateId() string {
    return r._templateId
}
// TemplateVariable Setter
// 模板变量
func (r *TaobaoTmallgenieHotelwelcomeRequest) SetTemplateVariable(_templateVariable string) error {
    r._templateVariable = _templateVariable
    r.Set("template_variable", _templateVariable)
    return nil
}

// TemplateVariable Getter
func (r TaobaoTmallgenieHotelwelcomeRequest) GetTemplateVariable() string {
    return r._templateVariable
}
