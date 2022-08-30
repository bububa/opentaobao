package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentOrderConfirmAPIRequest 确认出票 API请求
// taobao.train.agent.order.confirm
//
// 确认出票
type TaobaoTrainAgentOrderConfirmAPIRequest struct {
	model.Params
	// 入参
	_param *BookTicketConfirmRQ
}

// NewTaobaoTrainAgentOrderConfirmRequest 初始化TaobaoTrainAgentOrderConfirmAPIRequest对象
func NewTaobaoTrainAgentOrderConfirmRequest() *TaobaoTrainAgentOrderConfirmAPIRequest {
	return &TaobaoTrainAgentOrderConfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentOrderConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.order.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentOrderConfirmAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 入参
func (r *TaobaoTrainAgentOrderConfirmAPIRequest) SetParam(_param *BookTicketConfirmRQ) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoTrainAgentOrderConfirmAPIRequest) GetParam() *BookTicketConfirmRQ {
	return r._param
}
