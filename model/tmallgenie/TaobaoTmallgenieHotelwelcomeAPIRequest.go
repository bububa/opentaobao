package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmallgenieHotelwelcomeAPIRequest 酒店欢迎词推送 API请求
// taobao.tmallgenie.hotelwelcome
//
// 推送欢迎词，让天猫精灵播放
type TaobaoTmallgenieHotelwelcomeAPIRequest struct {
	model.Params
	// 房间号
	_roomNo string
	// 酒店ID
	_hotelId int64
	// 模板ID
	_templateId string
	// 模板变量
	_templateVariable string
}

// NewTaobaoTmallgenieHotelwelcomeRequest 初始化TaobaoTmallgenieHotelwelcomeAPIRequest对象
func NewTaobaoTmallgenieHotelwelcomeRequest() *TaobaoTmallgenieHotelwelcomeAPIRequest {
	return &TaobaoTmallgenieHotelwelcomeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTmallgenieHotelwelcomeAPIRequest) GetApiMethodName() string {
	return "taobao.tmallgenie.hotelwelcome"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTmallgenieHotelwelcomeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRoomNo is RoomNo Setter
// 房间号
func (r *TaobaoTmallgenieHotelwelcomeAPIRequest) SetRoomNo(_roomNo string) error {
	r._roomNo = _roomNo
	r.Set("room_no", _roomNo)
	return nil
}

// GetRoomNo RoomNo Getter
func (r TaobaoTmallgenieHotelwelcomeAPIRequest) GetRoomNo() string {
	return r._roomNo
}

// SetHotelId is HotelId Setter
// 酒店ID
func (r *TaobaoTmallgenieHotelwelcomeAPIRequest) SetHotelId(_hotelId int64) error {
	r._hotelId = _hotelId
	r.Set("hotel_id", _hotelId)
	return nil
}

// GetHotelId HotelId Getter
func (r TaobaoTmallgenieHotelwelcomeAPIRequest) GetHotelId() int64 {
	return r._hotelId
}

// SetTemplateId is TemplateId Setter
// 模板ID
func (r *TaobaoTmallgenieHotelwelcomeAPIRequest) SetTemplateId(_templateId string) error {
	r._templateId = _templateId
	r.Set("template_id", _templateId)
	return nil
}

// GetTemplateId TemplateId Getter
func (r TaobaoTmallgenieHotelwelcomeAPIRequest) GetTemplateId() string {
	return r._templateId
}

// SetTemplateVariable is TemplateVariable Setter
// 模板变量
func (r *TaobaoTmallgenieHotelwelcomeAPIRequest) SetTemplateVariable(_templateVariable string) error {
	r._templateVariable = _templateVariable
	r.Set("template_variable", _templateVariable)
	return nil
}

// GetTemplateVariable TemplateVariable Getter
func (r TaobaoTmallgenieHotelwelcomeAPIRequest) GetTemplateVariable() string {
	return r._templateVariable
}
