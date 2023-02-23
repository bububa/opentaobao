package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentDirectCompensateAPIRequest 火车票代理商接口——订单关闭实际出票成功审计接口 API请求
// taobao.train.agent.direct.compensate
//
// 代购直连订单平台关单但是代理商出票成功补偿接口
type TaobaoTrainAgentDirectCompensateAPIRequest struct {
	model.Params
	// 出票成功补偿入参
	_compensateParam *CompensateParam
}

// NewTaobaoTrainAgentDirectCompensateRequest 初始化TaobaoTrainAgentDirectCompensateAPIRequest对象
func NewTaobaoTrainAgentDirectCompensateRequest() *TaobaoTrainAgentDirectCompensateAPIRequest {
	return &TaobaoTrainAgentDirectCompensateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentDirectCompensateAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.direct.compensate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentDirectCompensateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTrainAgentDirectCompensateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCompensateParam is CompensateParam Setter
// 出票成功补偿入参
func (r *TaobaoTrainAgentDirectCompensateAPIRequest) SetCompensateParam(_compensateParam *CompensateParam) error {
	r._compensateParam = _compensateParam
	r.Set("compensate_param", _compensateParam)
	return nil
}

// GetCompensateParam CompensateParam Getter
func (r TaobaoTrainAgentDirectCompensateAPIRequest) GetCompensateParam() *CompensateParam {
	return r._compensateParam
}
