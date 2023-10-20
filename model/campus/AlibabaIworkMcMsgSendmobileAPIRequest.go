package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIworkMcMsgSendmobileAPIRequest 发送消息给手机用户 API请求
// alibaba.iwork.mc.msg.sendmobile
//
// 给手机用户发送对应操作结果的消息
type AlibabaIworkMcMsgSendmobileAPIRequest struct {
	model.Params
	// 消息对象
	_mobileReceiverMessageEvent *MobileReceiverMessageEvent
}

// NewAlibabaIworkMcMsgSendmobileRequest 初始化AlibabaIworkMcMsgSendmobileAPIRequest对象
func NewAlibabaIworkMcMsgSendmobileRequest() *AlibabaIworkMcMsgSendmobileAPIRequest {
	return &AlibabaIworkMcMsgSendmobileAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIworkMcMsgSendmobileAPIRequest) Reset() {
	r._mobileReceiverMessageEvent = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIworkMcMsgSendmobileAPIRequest) GetApiMethodName() string {
	return "alibaba.iwork.mc.msg.sendmobile"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIworkMcMsgSendmobileAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIworkMcMsgSendmobileAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMobileReceiverMessageEvent is MobileReceiverMessageEvent Setter
// 消息对象
func (r *AlibabaIworkMcMsgSendmobileAPIRequest) SetMobileReceiverMessageEvent(_mobileReceiverMessageEvent *MobileReceiverMessageEvent) error {
	r._mobileReceiverMessageEvent = _mobileReceiverMessageEvent
	r.Set("mobile_receiver_message_event", _mobileReceiverMessageEvent)
	return nil
}

// GetMobileReceiverMessageEvent MobileReceiverMessageEvent Getter
func (r AlibabaIworkMcMsgSendmobileAPIRequest) GetMobileReceiverMessageEvent() *MobileReceiverMessageEvent {
	return r._mobileReceiverMessageEvent
}

var poolAlibabaIworkMcMsgSendmobileAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIworkMcMsgSendmobileRequest()
	},
}

// GetAlibabaIworkMcMsgSendmobileRequest 从 sync.Pool 获取 AlibabaIworkMcMsgSendmobileAPIRequest
func GetAlibabaIworkMcMsgSendmobileAPIRequest() *AlibabaIworkMcMsgSendmobileAPIRequest {
	return poolAlibabaIworkMcMsgSendmobileAPIRequest.Get().(*AlibabaIworkMcMsgSendmobileAPIRequest)
}

// ReleaseAlibabaIworkMcMsgSendmobileAPIRequest 将 AlibabaIworkMcMsgSendmobileAPIRequest 放入 sync.Pool
func ReleaseAlibabaIworkMcMsgSendmobileAPIRequest(v *AlibabaIworkMcMsgSendmobileAPIRequest) {
	v.Reset()
	poolAlibabaIworkMcMsgSendmobileAPIRequest.Put(v)
}
