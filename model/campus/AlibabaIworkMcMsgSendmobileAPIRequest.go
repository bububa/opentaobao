package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaiworkmcmsgsendmobileAPIRequest 发送消息给手机用户 API请求
// alibaba.iwork.mc.msg.sendmobile
//
// 给手机用户发送对应操作结果的消息
type AlibabaiworkmcmsgsendmobileAPIRequest struct {
	model.Params
	// 消息对象
	_mobileReceiverMessageEvent *MobileReceiverMessageEvent
}

// NewAlibabaiworkmcmsgsendmobileRequest 初始化AlibabaiworkmcmsgsendmobileAPIRequest对象
func NewAlibabaiworkmcmsgsendmobileRequest() *AlibabaiworkmcmsgsendmobileAPIRequest {
	return &AlibabaiworkmcmsgsendmobileAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaiworkmcmsgsendmobileAPIRequest) GetApiMethodName() string {
	return "alibaba.iwork.mc.msg.sendmobile"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaiworkmcmsgsendmobileAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaiworkmcmsgsendmobileAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMobileReceiverMessageEvent is MobileReceiverMessageEvent Setter
// 消息对象
func (r *AlibabaiworkmcmsgsendmobileAPIRequest) SetMobileReceiverMessageEvent(_mobileReceiverMessageEvent *MobileReceiverMessageEvent) error {
	r._mobileReceiverMessageEvent = _mobileReceiverMessageEvent
	r.Set("mobile_receiver_message_event", _mobileReceiverMessageEvent)
	return nil
}

// GetMobileReceiverMessageEvent MobileReceiverMessageEvent Getter
func (r AlibabaiworkmcmsgsendmobileAPIRequest) GetMobileReceiverMessageEvent() *MobileReceiverMessageEvent {
	return r._mobileReceiverMessageEvent
}
