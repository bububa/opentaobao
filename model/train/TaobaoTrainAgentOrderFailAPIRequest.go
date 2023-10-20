package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotrainagentorderfailAPIRequest 出票失败 API请求
// taobao.train.agent.order.fail
//
// 出票失败
type TaobaotrainagentorderfailAPIRequest struct {
	model.Params
	// rq
	_param *BookTicketFailRq
}

// NewTaobaotrainagentorderfailRequest 初始化TaobaotrainagentorderfailAPIRequest对象
func NewTaobaotrainagentorderfailRequest() *TaobaotrainagentorderfailAPIRequest {
	return &TaobaotrainagentorderfailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotrainagentorderfailAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.order.fail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotrainagentorderfailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotrainagentorderfailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// rq
func (r *TaobaotrainagentorderfailAPIRequest) SetParam(_param *BookTicketFailRq) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaotrainagentorderfailAPIRequest) GetParam() *BookTicketFailRq {
	return r._param
}
