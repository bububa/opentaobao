package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaiworkmcmsgsenddefaultAPIRequest 给注册用户发送消息 API请求
// alibaba.iwork.mc.msg.senddefault
//
// 给神鲸注册用户发送对应操作结果的消息
type AlibabaiworkmcmsgsenddefaultAPIRequest struct {
	model.Params
	// 消息对象
	_messageEvent *DefaultMessageEvent
}

// NewAlibabaiworkmcmsgsenddefaultRequest 初始化AlibabaiworkmcmsgsenddefaultAPIRequest对象
func NewAlibabaiworkmcmsgsenddefaultRequest() *AlibabaiworkmcmsgsenddefaultAPIRequest {
	return &AlibabaiworkmcmsgsenddefaultAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaiworkmcmsgsenddefaultAPIRequest) GetApiMethodName() string {
	return "alibaba.iwork.mc.msg.senddefault"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaiworkmcmsgsenddefaultAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaiworkmcmsgsenddefaultAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMessageEvent is MessageEvent Setter
// 消息对象
func (r *AlibabaiworkmcmsgsenddefaultAPIRequest) SetMessageEvent(_messageEvent *DefaultMessageEvent) error {
	r._messageEvent = _messageEvent
	r.Set("message_event", _messageEvent)
	return nil
}

// GetMessageEvent MessageEvent Getter
func (r AlibabaiworkmcmsgsenddefaultAPIRequest) GetMessageEvent() *DefaultMessageEvent {
	return r._messageEvent
}
