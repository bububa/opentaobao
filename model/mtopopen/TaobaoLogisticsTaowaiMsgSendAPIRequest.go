package mtopopen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsTaowaiMsgSendAPIRequest 淘外包裹物流信息走淘宝发包裹状态通知接口 API请求
// taobao.logistics.taowai.msg.send
//
// 淘外包裹物流信息走淘宝发包裹状态通知接口
type TaobaoLogisticsTaowaiMsgSendAPIRequest struct {
	model.Params
	// 请求对象
	_taowaiMsgSendRequest *TaowaiMsgSendRequest
}

// NewTaobaoLogisticsTaowaiMsgSendRequest 初始化TaobaoLogisticsTaowaiMsgSendAPIRequest对象
func NewTaobaoLogisticsTaowaiMsgSendRequest() *TaobaoLogisticsTaowaiMsgSendAPIRequest {
	return &TaobaoLogisticsTaowaiMsgSendAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsTaowaiMsgSendAPIRequest) Reset() {
	r._taowaiMsgSendRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsTaowaiMsgSendAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.taowai.msg.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsTaowaiMsgSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsTaowaiMsgSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTaowaiMsgSendRequest is TaowaiMsgSendRequest Setter
// 请求对象
func (r *TaobaoLogisticsTaowaiMsgSendAPIRequest) SetTaowaiMsgSendRequest(_taowaiMsgSendRequest *TaowaiMsgSendRequest) error {
	r._taowaiMsgSendRequest = _taowaiMsgSendRequest
	r.Set("taowai_msg_send_request", _taowaiMsgSendRequest)
	return nil
}

// GetTaowaiMsgSendRequest TaowaiMsgSendRequest Getter
func (r TaobaoLogisticsTaowaiMsgSendAPIRequest) GetTaowaiMsgSendRequest() *TaowaiMsgSendRequest {
	return r._taowaiMsgSendRequest
}

var poolTaobaoLogisticsTaowaiMsgSendAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsTaowaiMsgSendRequest()
	},
}

// GetTaobaoLogisticsTaowaiMsgSendRequest 从 sync.Pool 获取 TaobaoLogisticsTaowaiMsgSendAPIRequest
func GetTaobaoLogisticsTaowaiMsgSendAPIRequest() *TaobaoLogisticsTaowaiMsgSendAPIRequest {
	return poolTaobaoLogisticsTaowaiMsgSendAPIRequest.Get().(*TaobaoLogisticsTaowaiMsgSendAPIRequest)
}

// ReleaseTaobaoLogisticsTaowaiMsgSendAPIRequest 将 TaobaoLogisticsTaowaiMsgSendAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsTaowaiMsgSendAPIRequest(v *TaobaoLogisticsTaowaiMsgSendAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsTaowaiMsgSendAPIRequest.Put(v)
}
