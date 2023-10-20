package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsFengchaoMsgSendAPIRequest 丰巢走淘宝发包裹状态通知接口 API请求
// taobao.logistics.fengchao.msg.send
//
// 丰巢走淘宝发包裹状态通知接口
type TaobaoLogisticsFengchaoMsgSendAPIRequest struct {
	model.Params
	// 请求对象
	_msgSendRequest *MsgSendRequest
}

// NewTaobaoLogisticsFengchaoMsgSendRequest 初始化TaobaoLogisticsFengchaoMsgSendAPIRequest对象
func NewTaobaoLogisticsFengchaoMsgSendRequest() *TaobaoLogisticsFengchaoMsgSendAPIRequest {
	return &TaobaoLogisticsFengchaoMsgSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsFengchaoMsgSendAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.fengchao.msg.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsFengchaoMsgSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsFengchaoMsgSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMsgSendRequest is MsgSendRequest Setter
// 请求对象
func (r *TaobaoLogisticsFengchaoMsgSendAPIRequest) SetMsgSendRequest(_msgSendRequest *MsgSendRequest) error {
	r._msgSendRequest = _msgSendRequest
	r.Set("msg_send_request", _msgSendRequest)
	return nil
}

// GetMsgSendRequest MsgSendRequest Getter
func (r TaobaoLogisticsFengchaoMsgSendAPIRequest) GetMsgSendRequest() *MsgSendRequest {
	return r._msgSendRequest
}
