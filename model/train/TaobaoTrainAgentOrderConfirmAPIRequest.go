package train

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentOrderConfirmAPIRequest 确认出票 API请求
// taobao.train.agent.order.confirm
//
// 确认出票
type TaobaoTrainAgentOrderConfirmAPIRequest struct {
	model.Params
	// 入参
	_param *BookTicketConfirmRq
}

// NewTaobaoTrainAgentOrderConfirmRequest 初始化TaobaoTrainAgentOrderConfirmAPIRequest对象
func NewTaobaoTrainAgentOrderConfirmRequest() *TaobaoTrainAgentOrderConfirmAPIRequest {
	return &TaobaoTrainAgentOrderConfirmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTrainAgentOrderConfirmAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentOrderConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.order.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentOrderConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTrainAgentOrderConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *TaobaoTrainAgentOrderConfirmAPIRequest) SetParam(_param *BookTicketConfirmRq) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoTrainAgentOrderConfirmAPIRequest) GetParam() *BookTicketConfirmRq {
	return r._param
}

var poolTaobaoTrainAgentOrderConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTrainAgentOrderConfirmRequest()
	},
}

// GetTaobaoTrainAgentOrderConfirmRequest 从 sync.Pool 获取 TaobaoTrainAgentOrderConfirmAPIRequest
func GetTaobaoTrainAgentOrderConfirmAPIRequest() *TaobaoTrainAgentOrderConfirmAPIRequest {
	return poolTaobaoTrainAgentOrderConfirmAPIRequest.Get().(*TaobaoTrainAgentOrderConfirmAPIRequest)
}

// ReleaseTaobaoTrainAgentOrderConfirmAPIRequest 将 TaobaoTrainAgentOrderConfirmAPIRequest 放入 sync.Pool
func ReleaseTaobaoTrainAgentOrderConfirmAPIRequest(v *TaobaoTrainAgentOrderConfirmAPIRequest) {
	v.Reset()
	poolTaobaoTrainAgentOrderConfirmAPIRequest.Put(v)
}
