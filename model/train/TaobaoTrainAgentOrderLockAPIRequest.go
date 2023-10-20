package train

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentOrderLockAPIRequest 锁单 API请求
// taobao.train.agent.order.lock
//
// 锁单
type TaobaoTrainAgentOrderLockAPIRequest struct {
	model.Params
	// 入参
	_param *LockOrderRq
}

// NewTaobaoTrainAgentOrderLockRequest 初始化TaobaoTrainAgentOrderLockAPIRequest对象
func NewTaobaoTrainAgentOrderLockRequest() *TaobaoTrainAgentOrderLockAPIRequest {
	return &TaobaoTrainAgentOrderLockAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTrainAgentOrderLockAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentOrderLockAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.order.lock"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentOrderLockAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTrainAgentOrderLockAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *TaobaoTrainAgentOrderLockAPIRequest) SetParam(_param *LockOrderRq) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoTrainAgentOrderLockAPIRequest) GetParam() *LockOrderRq {
	return r._param
}

var poolTaobaoTrainAgentOrderLockAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTrainAgentOrderLockRequest()
	},
}

// GetTaobaoTrainAgentOrderLockRequest 从 sync.Pool 获取 TaobaoTrainAgentOrderLockAPIRequest
func GetTaobaoTrainAgentOrderLockAPIRequest() *TaobaoTrainAgentOrderLockAPIRequest {
	return poolTaobaoTrainAgentOrderLockAPIRequest.Get().(*TaobaoTrainAgentOrderLockAPIRequest)
}

// ReleaseTaobaoTrainAgentOrderLockAPIRequest 将 TaobaoTrainAgentOrderLockAPIRequest 放入 sync.Pool
func ReleaseTaobaoTrainAgentOrderLockAPIRequest(v *TaobaoTrainAgentOrderLockAPIRequest) {
	v.Reset()
	poolTaobaoTrainAgentOrderLockAPIRequest.Put(v)
}
