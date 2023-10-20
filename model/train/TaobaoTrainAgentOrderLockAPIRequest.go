package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotrainagentorderlockAPIRequest 锁单 API请求
// taobao.train.agent.order.lock
//
// 锁单
type TaobaotrainagentorderlockAPIRequest struct {
	model.Params
	// 入参
	_param *LockOrderRq
}

// NewTaobaotrainagentorderlockRequest 初始化TaobaotrainagentorderlockAPIRequest对象
func NewTaobaotrainagentorderlockRequest() *TaobaotrainagentorderlockAPIRequest {
	return &TaobaotrainagentorderlockAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotrainagentorderlockAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.order.lock"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotrainagentorderlockAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotrainagentorderlockAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *TaobaotrainagentorderlockAPIRequest) SetParam(_param *LockOrderRq) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaotrainagentorderlockAPIRequest) GetParam() *LockOrderRq {
	return r._param
}
