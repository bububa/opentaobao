package train

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentOrderIgnoreAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.order.ignore"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentOrderIgnoreAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
