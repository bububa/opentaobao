package mtopopen

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsFengchaoMsgSendAPIRequest) Reset() {
	r._msgSendRequest = nil
	r.Params.ToZero()
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

var poolTaobaoLogisticsFengchaoMsgSendAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsFengchaoMsgSendRequest()
	},
}

// GetTaobaoLogisticsFengchaoMsgSendRequest 从 sync.Pool 获取 TaobaoLogisticsFengchaoMsgSendAPIRequest
func GetTaobaoLogisticsFengchaoMsgSendAPIRequest() *TaobaoLogisticsFengchaoMsgSendAPIRequest {
	return poolTaobaoLogisticsFengchaoMsgSendAPIRequest.Get().(*TaobaoLogisticsFengchaoMsgSendAPIRequest)
}

// ReleaseTaobaoLogisticsFengchaoMsgSendAPIRequest 将 TaobaoLogisticsFengchaoMsgSendAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsFengchaoMsgSendAPIRequest(v *TaobaoLogisticsFengchaoMsgSendAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsFengchaoMsgSendAPIRequest.Put(v)
}
