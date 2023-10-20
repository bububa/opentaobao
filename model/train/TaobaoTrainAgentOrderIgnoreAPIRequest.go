package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotrainagentorderignoreAPIRequest 忽略订单 API请求
// taobao.train.agent.order.ignore
//
// 忽略订单
type TaobaotrainagentorderignoreAPIRequest struct {
	model.Params
	// rq
	_param0 *IgnoreOrderRq
}

// NewTaobaotrainagentorderignoreRequest 初始化TaobaotrainagentorderignoreAPIRequest对象
func NewTaobaotrainagentorderignoreRequest() *TaobaotrainagentorderignoreAPIRequest {
	return &TaobaotrainagentorderignoreAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotrainagentorderignoreAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.order.ignore"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotrainagentorderignoreAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotrainagentorderignoreAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// rq
func (r *TaobaotrainagentorderignoreAPIRequest) SetParam0(_param0 *IgnoreOrderRq) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaotrainagentorderignoreAPIRequest) GetParam0() *IgnoreOrderRq {
	return r._param0
}
