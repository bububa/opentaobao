package train

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentOrderFailAPIRequest 出票失败 API请求
// taobao.train.agent.order.fail
//
// 出票失败
type TaobaoTrainAgentOrderFailAPIRequest struct {
	model.Params
	// rq
	_param *BookTicketFailRq
}

// NewTaobaoTrainAgentOrderFailRequest 初始化TaobaoTrainAgentOrderFailAPIRequest对象
func NewTaobaoTrainAgentOrderFailRequest() *TaobaoTrainAgentOrderFailAPIRequest {
	return &TaobaoTrainAgentOrderFailAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTrainAgentOrderFailAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentOrderFailAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.order.fail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentOrderFailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTrainAgentOrderFailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// rq
func (r *TaobaoTrainAgentOrderFailAPIRequest) SetParam(_param *BookTicketFailRq) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoTrainAgentOrderFailAPIRequest) GetParam() *BookTicketFailRq {
	return r._param
}

var poolTaobaoTrainAgentOrderFailAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTrainAgentOrderFailRequest()
	},
}

// GetTaobaoTrainAgentOrderFailRequest 从 sync.Pool 获取 TaobaoTrainAgentOrderFailAPIRequest
func GetTaobaoTrainAgentOrderFailAPIRequest() *TaobaoTrainAgentOrderFailAPIRequest {
	return poolTaobaoTrainAgentOrderFailAPIRequest.Get().(*TaobaoTrainAgentOrderFailAPIRequest)
}

// ReleaseTaobaoTrainAgentOrderFailAPIRequest 将 TaobaoTrainAgentOrderFailAPIRequest 放入 sync.Pool
func ReleaseTaobaoTrainAgentOrderFailAPIRequest(v *TaobaoTrainAgentOrderFailAPIRequest) {
	v.Reset()
	poolTaobaoTrainAgentOrderFailAPIRequest.Put(v)
}
