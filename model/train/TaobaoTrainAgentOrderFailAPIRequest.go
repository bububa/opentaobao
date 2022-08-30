package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentOrderFailAPIRequest 出票失败 API请求
// taobao.train.agent.order.fail
//
// 出票失败
type TaobaoTrainAgentOrderFailAPIRequest struct {
	model.Params
	// rq
	_param *BookTicketFailRQ
}

// NewTaobaoTrainAgentOrderFailRequest 初始化TaobaoTrainAgentOrderFailAPIRequest对象
func NewTaobaoTrainAgentOrderFailRequest() *TaobaoTrainAgentOrderFailAPIRequest {
	return &TaobaoTrainAgentOrderFailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentOrderFailAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.order.fail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentOrderFailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// rq
func (r *TaobaoTrainAgentOrderFailAPIRequest) SetParam(_param *BookTicketFailRQ) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoTrainAgentOrderFailAPIRequest) GetParam() *BookTicketFailRQ {
	return r._param
}
