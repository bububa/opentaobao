package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotmallgeniehotelwelcomeAPIRequest 酒店欢迎词推送 API请求
// taobao.tmallgenie.hotelwelcome
//
// 推送欢迎词，让天猫精灵播放
type TaobaotmallgeniehotelwelcomeAPIRequest struct {
	model.Params
	// 房间号
	_roomNo string
	// 模板变量
	_templateVariable string
	// 模板ID
	_templateId string
	// 酒店ID
	_hotelId int64
}

// NewTaobaotmallgeniehotelwelcomeRequest 初始化TaobaotmallgeniehotelwelcomeAPIRequest对象
func NewTaobaotmallgeniehotelwelcomeRequest() *TaobaotmallgeniehotelwelcomeAPIRequest {
	return &TaobaotmallgeniehotelwelcomeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotmallgeniehotelwelcomeAPIRequest) GetApiMethodName() string {
	return "taobao.tmallgenie.hotelwelcome"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotmallgeniehotelwelcomeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotmallgeniehotelwelcomeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRoomNo is RoomNo Setter
// 房间号
func (r *TaobaotmallgeniehotelwelcomeAPIRequest) SetRoomNo(_roomNo string) error {
	r._roomNo = _roomNo
	r.Set("room_no", _roomNo)
	return nil
}

// GetRoomNo RoomNo Getter
func (r TaobaotmallgeniehotelwelcomeAPIRequest) GetRoomNo() string {
	return r._roomNo
}

// SetTemplateVariable is TemplateVariable Setter
// 模板变量
func (r *TaobaotmallgeniehotelwelcomeAPIRequest) SetTemplateVariable(_templateVariable string) error {
	r._templateVariable = _templateVariable
	r.Set("template_variable", _templateVariable)
	return nil
}

// GetTemplateVariable TemplateVariable Getter
func (r TaobaotmallgeniehotelwelcomeAPIRequest) GetTemplateVariable() string {
	return r._templateVariable
}

// SetTemplateId is TemplateId Setter
// 模板ID
func (r *TaobaotmallgeniehotelwelcomeAPIRequest) SetTemplateId(_templateId string) error {
	r._templateId = _templateId
	r.Set("template_id", _templateId)
	return nil
}

// GetTemplateId TemplateId Getter
func (r TaobaotmallgeniehotelwelcomeAPIRequest) GetTemplateId() string {
	return r._templateId
}

// SetHotelId is HotelId Setter
// 酒店ID
func (r *TaobaotmallgeniehotelwelcomeAPIRequest) SetHotelId(_hotelId int64) error {
	r._hotelId = _hotelId
	r.Set("hotel_id", _hotelId)
	return nil
}

// GetHotelId HotelId Getter
func (r TaobaotmallgeniehotelwelcomeAPIRequest) GetHotelId() int64 {
	return r._hotelId
}
