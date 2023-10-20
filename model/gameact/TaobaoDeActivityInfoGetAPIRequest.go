package gameact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaodeactivityinfogetAPIRequest 获取活动信息 API请求
// taobao.de.activity.info.get
//
// 根据appKey和活动id获取活动
type TaobaodeactivityinfogetAPIRequest struct {
	model.Params
	// 事件唯一标识
	_eventKey string
}

// NewTaobaodeactivityinfogetRequest 初始化TaobaodeactivityinfogetAPIRequest对象
func NewTaobaodeactivityinfogetRequest() *TaobaodeactivityinfogetAPIRequest {
	return &TaobaodeactivityinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaodeactivityinfogetAPIRequest) GetApiMethodName() string {
	return "taobao.de.activity.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaodeactivityinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaodeactivityinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEventKey is EventKey Setter
// 事件唯一标识
func (r *TaobaodeactivityinfogetAPIRequest) SetEventKey(_eventKey string) error {
	r._eventKey = _eventKey
	r.Set("event_key", _eventKey)
	return nil
}

// GetEventKey EventKey Getter
func (r TaobaodeactivityinfogetAPIRequest) GetEventKey() string {
	return r._eventKey
}
