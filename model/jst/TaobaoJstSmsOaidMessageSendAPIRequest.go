package jst

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsOaidMessageSendAPIRequest 基于OAID的短信发送接口 API请求
// taobao.jst.sms.oaid.message.send
//
// 基于OAID的短信发送接口
type TaobaoJstSmsOaidMessageSendAPIRequest struct {
	model.Params
	// 单个OAID发送短信的入参
	_paramSendMessageByOAIDRequest *SendMessageByOaidRequest
}

// NewTaobaoJstSmsOaidMessageSendRequest 初始化TaobaoJstSmsOaidMessageSendAPIRequest对象
func NewTaobaoJstSmsOaidMessageSendRequest() *TaobaoJstSmsOaidMessageSendAPIRequest {
	return &TaobaoJstSmsOaidMessageSendAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJstSmsOaidMessageSendAPIRequest) Reset() {
	r._paramSendMessageByOAIDRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsOaidMessageSendAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.oaid.message.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsOaidMessageSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstSmsOaidMessageSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamSendMessageByOAIDRequest is ParamSendMessageByOAIDRequest Setter
// 单个OAID发送短信的入参
func (r *TaobaoJstSmsOaidMessageSendAPIRequest) SetParamSendMessageByOAIDRequest(_paramSendMessageByOAIDRequest *SendMessageByOaidRequest) error {
	r._paramSendMessageByOAIDRequest = _paramSendMessageByOAIDRequest
	r.Set("param_send_message_by_o_a_i_d_request", _paramSendMessageByOAIDRequest)
	return nil
}

// GetParamSendMessageByOAIDRequest ParamSendMessageByOAIDRequest Getter
func (r TaobaoJstSmsOaidMessageSendAPIRequest) GetParamSendMessageByOAIDRequest() *SendMessageByOaidRequest {
	return r._paramSendMessageByOAIDRequest
}

var poolTaobaoJstSmsOaidMessageSendAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJstSmsOaidMessageSendRequest()
	},
}

// GetTaobaoJstSmsOaidMessageSendRequest 从 sync.Pool 获取 TaobaoJstSmsOaidMessageSendAPIRequest
func GetTaobaoJstSmsOaidMessageSendAPIRequest() *TaobaoJstSmsOaidMessageSendAPIRequest {
	return poolTaobaoJstSmsOaidMessageSendAPIRequest.Get().(*TaobaoJstSmsOaidMessageSendAPIRequest)
}

// ReleaseTaobaoJstSmsOaidMessageSendAPIRequest 将 TaobaoJstSmsOaidMessageSendAPIRequest 放入 sync.Pool
func ReleaseTaobaoJstSmsOaidMessageSendAPIRequest(v *TaobaoJstSmsOaidMessageSendAPIRequest) {
	v.Reset()
	poolTaobaoJstSmsOaidMessageSendAPIRequest.Put(v)
}
