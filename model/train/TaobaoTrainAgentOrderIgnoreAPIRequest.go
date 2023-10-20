package train

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentOrderIgnoreAPIRequest 忽略订单 API请求
// taobao.train.agent.order.ignore
//
// 忽略订单
type TaobaoTrainAgentOrderIgnoreAPIRequest struct {
	model.Params
	// rq
	_param0 *IgnoreOrderRq
}

// NewTaobaoTrainAgentOrderIgnoreRequest 初始化TaobaoTrainAgentOrderIgnoreAPIRequest对象
func NewTaobaoTrainAgentOrderIgnoreRequest() *TaobaoTrainAgentOrderIgnoreAPIRequest {
	return &TaobaoTrainAgentOrderIgnoreAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTrainAgentOrderIgnoreAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentOrderIgnoreAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.order.ignore"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentOrderIgnoreAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTrainAgentOrderIgnoreAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// rq
func (r *TaobaoTrainAgentOrderIgnoreAPIRequest) SetParam0(_param0 *IgnoreOrderRq) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoTrainAgentOrderIgnoreAPIRequest) GetParam0() *IgnoreOrderRq {
	return r._param0
}

var poolTaobaoTrainAgentOrderIgnoreAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTrainAgentOrderIgnoreRequest()
	},
}

// GetTaobaoTrainAgentOrderIgnoreRequest 从 sync.Pool 获取 TaobaoTrainAgentOrderIgnoreAPIRequest
func GetTaobaoTrainAgentOrderIgnoreAPIRequest() *TaobaoTrainAgentOrderIgnoreAPIRequest {
	return poolTaobaoTrainAgentOrderIgnoreAPIRequest.Get().(*TaobaoTrainAgentOrderIgnoreAPIRequest)
}

// ReleaseTaobaoTrainAgentOrderIgnoreAPIRequest 将 TaobaoTrainAgentOrderIgnoreAPIRequest 放入 sync.Pool
func ReleaseTaobaoTrainAgentOrderIgnoreAPIRequest(v *TaobaoTrainAgentOrderIgnoreAPIRequest) {
	v.Reset()
	poolTaobaoTrainAgentOrderIgnoreAPIRequest.Put(v)
}
