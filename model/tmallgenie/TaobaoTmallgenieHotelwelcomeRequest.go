package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店欢迎词推送 APIRequest
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

func NewTaobaoTmallgenieHotelwelcomeRequest() *TaobaoTmallgenieHotelwelcomeRequest{
    return &TaobaoTmallgenieHotelwelcomeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTmallgenieHotelwelcomeRequest) GetApiMethodName() string {
    return "taobao.tmallgenie.hotelwelcome"
}

func (r TaobaoTmallgenieHotelwelcomeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTmallgenieHotelwelcomeRequest) SetRoomNo(roomNo string) error {
    r.roomNo = roomNo
    r.Set("room_no", roomNo)
    return nil
}

func (r TaobaoTmallgenieHotelwelcomeRequest) GetRoomNo() string {
    return r.roomNo
}

func (r *TaobaoTmallgenieHotelwelcomeRequest) SetHotelId(hotelId int64) error {
    r.hotelId = hotelId
    r.Set("hotel_id", hotelId)
    return nil
}

func (r TaobaoTmallgenieHotelwelcomeRequest) GetHotelId() int64 {
    return r.hotelId
}

func (r *TaobaoTmallgenieHotelwelcomeRequest) SetTemplateId(templateId string) error {
    r.templateId = templateId
    r.Set("template_id", templateId)
    return nil
}

func (r TaobaoTmallgenieHotelwelcomeRequest) GetTemplateId() string {
    return r.templateId
}

func (r *TaobaoTmallgenieHotelwelcomeRequest) SetTemplateVariable(templateVariable string) error {
    r.templateVariable = templateVariable
    r.Set("template_variable", templateVariable)
    return nil
}

func (r TaobaoTmallgenieHotelwelcomeRequest) GetTemplateVariable() string {
    return r.templateVariable
}

