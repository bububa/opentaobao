package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstsmsoaidmessagesendAPIRequest 基于OAID的短信发送接口 API请求
// taobao.jst.sms.oaid.message.send
//
// 基于OAID的短信发送接口
type TaobaojstsmsoaidmessagesendAPIRequest struct {
	model.Params
	// 单个OAID发送短信的入参
	_paramSendMessageByOAIDRequest *SendMessageByOaidRequest
}

// NewTaobaojstsmsoaidmessagesendRequest 初始化TaobaojstsmsoaidmessagesendAPIRequest对象
func NewTaobaojstsmsoaidmessagesendRequest() *TaobaojstsmsoaidmessagesendAPIRequest {
	return &TaobaojstsmsoaidmessagesendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstsmsoaidmessagesendAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.oaid.message.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstsmsoaidmessagesendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstsmsoaidmessagesendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamSendMessageByOAIDRequest is ParamSendMessageByOAIDRequest Setter
// 单个OAID发送短信的入参
func (r *TaobaojstsmsoaidmessagesendAPIRequest) SetParamSendMessageByOAIDRequest(_paramSendMessageByOAIDRequest *SendMessageByOaidRequest) error {
	r._paramSendMessageByOAIDRequest = _paramSendMessageByOAIDRequest
	r.Set("param_send_message_by_o_a_i_d_request", _paramSendMessageByOAIDRequest)
	return nil
}

// GetParamSendMessageByOAIDRequest ParamSendMessageByOAIDRequest Getter
func (r TaobaojstsmsoaidmessagesendAPIRequest) GetParamSendMessageByOAIDRequest() *SendMessageByOaidRequest {
	return r._paramSendMessageByOAIDRequest
}
