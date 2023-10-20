package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotrainagentorderconfirmAPIRequest 确认出票 API请求
// taobao.train.agent.order.confirm
//
// 确认出票
type TaobaotrainagentorderconfirmAPIRequest struct {
	model.Params
	// 入参
	_param *BookTicketConfirmRq
}

// NewTaobaotrainagentorderconfirmRequest 初始化TaobaotrainagentorderconfirmAPIRequest对象
func NewTaobaotrainagentorderconfirmRequest() *TaobaotrainagentorderconfirmAPIRequest {
	return &TaobaotrainagentorderconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotrainagentorderconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.order.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotrainagentorderconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotrainagentorderconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *TaobaotrainagentorderconfirmAPIRequest) SetParam(_param *BookTicketConfirmRq) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaotrainagentorderconfirmAPIRequest) GetParam() *BookTicketConfirmRq {
	return r._param
}
