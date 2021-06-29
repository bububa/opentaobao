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
    roomNo   string
    // 酒店ID
    hotelId   int64
    // 模板ID
    templateId   string
    // 模板变量
    templateVariable   string
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
func (r *TaobaoTmallgenieHotelwelcomeRequest) SetRoomNo(roomNo string) error {
    r.roomNo = roomNo
    r.Set("room_no", roomNo)
    return nil
}

// RoomNo Getter
func (r TaobaoTmallgenieHotelwelcomeRequest) GetRoomNo() string {
    return r.roomNo
}
// HotelId Setter
// 酒店ID
func (r *TaobaoTmallgenieHotelwelcomeRequest) SetHotelId(hotelId int64) error {
    r.hotelId = hotelId
    r.Set("hotel_id", hotelId)
    return nil
}

// HotelId Getter
func (r TaobaoTmallgenieHotelwelcomeRequest) GetHotelId() int64 {
    return r.hotelId
}
// TemplateId Setter
// 模板ID
func (r *TaobaoTmallgenieHotelwelcomeRequest) SetTemplateId(templateId string) error {
    r.templateId = templateId
    r.Set("template_id", templateId)
    return nil
}

// TemplateId Getter
func (r TaobaoTmallgenieHotelwelcomeRequest) GetTemplateId() string {
    return r.templateId
}
// TemplateVariable Setter
// 模板变量
func (r *TaobaoTmallgenieHotelwelcomeRequest) SetTemplateVariable(templateVariable string) error {
    r.templateVariable = templateVariable
    r.Set("template_variable", templateVariable)
    return nil
}

// TemplateVariable Getter
func (r TaobaoTmallgenieHotelwelcomeRequest) GetTemplateVariable() string {
    return r.templateVariable
}
