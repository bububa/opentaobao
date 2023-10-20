package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsfengchaomsgsendAPIRequest 丰巢走淘宝发包裹状态通知接口 API请求
// taobao.logistics.fengchao.msg.send
//
// 丰巢走淘宝发包裹状态通知接口
type TaobaologisticsfengchaomsgsendAPIRequest struct {
	model.Params
	// 请求对象
	_msgSendRequest *MsgSendRequest
}

// NewTaobaologisticsfengchaomsgsendRequest 初始化TaobaologisticsfengchaomsgsendAPIRequest对象
func NewTaobaologisticsfengchaomsgsendRequest() *TaobaologisticsfengchaomsgsendAPIRequest {
	return &TaobaologisticsfengchaomsgsendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsfengchaomsgsendAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.fengchao.msg.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsfengchaomsgsendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsfengchaomsgsendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMsgSendRequest is MsgSendRequest Setter
// 请求对象
func (r *TaobaologisticsfengchaomsgsendAPIRequest) SetMsgSendRequest(_msgSendRequest *MsgSendRequest) error {
	r._msgSendRequest = _msgSendRequest
	r.Set("msg_send_request", _msgSendRequest)
	return nil
}

// GetMsgSendRequest MsgSendRequest Getter
func (r TaobaologisticsfengchaomsgsendAPIRequest) GetMsgSendRequest() *MsgSendRequest {
	return r._msgSendRequest
}
