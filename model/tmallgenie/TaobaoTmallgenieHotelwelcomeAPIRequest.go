package tmallgenie

import (
	"net/url"
	"sync"

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
	// 模板变量
	_templateVariable string
	// 模板ID
	_templateId string
	// 酒店ID
	_hotelId int64
}

// NewTaobaoTmallgenieHotelwelcomeRequest 初始化TaobaoTmallgenieHotelwelcomeAPIRequest对象
func NewTaobaoTmallgenieHotelwelcomeRequest() *TaobaoTmallgenieHotelwelcomeAPIRequest {
	return &TaobaoTmallgenieHotelwelcomeAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTmallgenieHotelwelcomeAPIRequest) Reset() {
	r._roomNo = ""
	r._templateVariable = ""
	r._templateId = ""
	r._hotelId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTmallgenieHotelwelcomeAPIRequest) GetApiMethodName() string {
	return "taobao.tmallgenie.hotelwelcome"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTmallgenieHotelwelcomeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTmallgenieHotelwelcomeAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoTmallgenieHotelwelcomeAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTmallgenieHotelwelcomeRequest()
	},
}

// GetTaobaoTmallgenieHotelwelcomeRequest 从 sync.Pool 获取 TaobaoTmallgenieHotelwelcomeAPIRequest
func GetTaobaoTmallgenieHotelwelcomeAPIRequest() *TaobaoTmallgenieHotelwelcomeAPIRequest {
	return poolTaobaoTmallgenieHotelwelcomeAPIRequest.Get().(*TaobaoTmallgenieHotelwelcomeAPIRequest)
}

// ReleaseTaobaoTmallgenieHotelwelcomeAPIRequest 将 TaobaoTmallgenieHotelwelcomeAPIRequest 放入 sync.Pool
func ReleaseTaobaoTmallgenieHotelwelcomeAPIRequest(v *TaobaoTmallgenieHotelwelcomeAPIRequest) {
	v.Reset()
	poolTaobaoTmallgenieHotelwelcomeAPIRequest.Put(v)
}
