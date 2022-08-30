package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentOrderLockAPIRequest 锁单 API请求
// taobao.train.agent.order.lock
//
// 锁单
type TaobaoTrainAgentOrderLockAPIRequest struct {
	model.Params
	// 入参
	_param *LockOrderRQ
}

// NewTaobaoTrainAgentOrderLockRequest 初始化TaobaoTrainAgentOrderLockAPIRequest对象
func NewTaobaoTrainAgentOrderLockRequest() *TaobaoTrainAgentOrderLockAPIRequest {
	return &TaobaoTrainAgentOrderLockAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentOrderLockAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.order.lock"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentOrderLockAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 入参
func (r *TaobaoTrainAgentOrderLockAPIRequest) SetParam(_param *LockOrderRQ) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoTrainAgentOrderLockAPIRequest) GetParam() *LockOrderRQ {
	return r._param
}
