package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIworkMcMsgSendmobileAPIRequest
发送消息给手机用户 API请求
alibaba.iwork.mc.msg.sendmobile

给手机用户发送对应操作结果的消息 */
type AlibabaIworkMcMsgSendmobileAPIRequest struct {
	model.Params
	// 消息对象
	_mobileReceiverMessageEvent *MobileReceiverMessageEvent
}

// NewAlibabaIworkMcMsgSendmobileRequest 初始化AlibabaIworkMcMsgSendmobileAPIRequest对象
func NewAlibabaIworkMcMsgSendmobileRequest() *AlibabaIworkMcMsgSendmobileAPIRequest {
	return &AlibabaIworkMcMsgSendmobileAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIworkMcMsgSendmobileAPIRequest) GetApiMethodName() string {
	return "alibaba.iwork.mc.msg.sendmobile"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIworkMcMsgSendmobileAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MobileReceiverMessageEvent Setter
// 消息对象
func (r *AlibabaIworkMcMsgSendmobileAPIRequest) SetMobileReceiverMessageEvent(_mobileReceiverMessageEvent *MobileReceiverMessageEvent) error {
	r._mobileReceiverMessageEvent = _mobileReceiverMessageEvent
	r.Set("mobile_receiver_message_event", _mobileReceiverMessageEvent)
	return nil
}

// Get MobileReceiverMessageEvent Getter
func (r AlibabaIworkMcMsgSendmobileAPIRequest) GetMobileReceiverMessageEvent() *MobileReceiverMessageEvent {
	return r._mobileReceiverMessageEvent
}
