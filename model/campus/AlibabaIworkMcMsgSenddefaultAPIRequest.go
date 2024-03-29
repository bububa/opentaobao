package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIworkMcMsgSenddefaultAPIRequest 给注册用户发送消息 API请求
// alibaba.iwork.mc.msg.senddefault
//
// 给神鲸注册用户发送对应操作结果的消息
type AlibabaIworkMcMsgSenddefaultAPIRequest struct {
	model.Params
	// 消息对象
	_messageEvent *DefaultMessageEvent
}

// NewAlibabaIworkMcMsgSenddefaultRequest 初始化AlibabaIworkMcMsgSenddefaultAPIRequest对象
func NewAlibabaIworkMcMsgSenddefaultRequest() *AlibabaIworkMcMsgSenddefaultAPIRequest {
	return &AlibabaIworkMcMsgSenddefaultAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIworkMcMsgSenddefaultAPIRequest) Reset() {
	r._messageEvent = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIworkMcMsgSenddefaultAPIRequest) GetApiMethodName() string {
	return "alibaba.iwork.mc.msg.senddefault"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIworkMcMsgSenddefaultAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIworkMcMsgSenddefaultAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMessageEvent is MessageEvent Setter
// 消息对象
func (r *AlibabaIworkMcMsgSenddefaultAPIRequest) SetMessageEvent(_messageEvent *DefaultMessageEvent) error {
	r._messageEvent = _messageEvent
	r.Set("message_event", _messageEvent)
	return nil
}

// GetMessageEvent MessageEvent Getter
func (r AlibabaIworkMcMsgSenddefaultAPIRequest) GetMessageEvent() *DefaultMessageEvent {
	return r._messageEvent
}

var poolAlibabaIworkMcMsgSenddefaultAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIworkMcMsgSenddefaultRequest()
	},
}

// GetAlibabaIworkMcMsgSenddefaultRequest 从 sync.Pool 获取 AlibabaIworkMcMsgSenddefaultAPIRequest
func GetAlibabaIworkMcMsgSenddefaultAPIRequest() *AlibabaIworkMcMsgSenddefaultAPIRequest {
	return poolAlibabaIworkMcMsgSenddefaultAPIRequest.Get().(*AlibabaIworkMcMsgSenddefaultAPIRequest)
}

// ReleaseAlibabaIworkMcMsgSenddefaultAPIRequest 将 AlibabaIworkMcMsgSenddefaultAPIRequest 放入 sync.Pool
func ReleaseAlibabaIworkMcMsgSenddefaultAPIRequest(v *AlibabaIworkMcMsgSenddefaultAPIRequest) {
	v.Reset()
	poolAlibabaIworkMcMsgSenddefaultAPIRequest.Put(v)
}
