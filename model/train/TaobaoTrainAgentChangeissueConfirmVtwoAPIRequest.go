package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotrainagentchangeissueconfirmvtwoAPIRequest 火车票代理商接口-跑腿改签出票回填-含鉴权校验 API请求
// taobao.train.agent.changeissue.confirm.vtwo
//
// 火车票代理商接口-跑腿改签出票回填-含鉴权校验
type TaobaotrainagentchangeissueconfirmvtwoAPIRequest struct {
	model.Params
	// 改签出票回填入参rq
	_changeIssueRq *ChangeIssueRq
}

// NewTaobaotrainagentchangeissueconfirmvtwoRequest 初始化TaobaotrainagentchangeissueconfirmvtwoAPIRequest对象
func NewTaobaotrainagentchangeissueconfirmvtwoRequest() *TaobaotrainagentchangeissueconfirmvtwoAPIRequest {
	return &TaobaotrainagentchangeissueconfirmvtwoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotrainagentchangeissueconfirmvtwoAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.changeissue.confirm.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotrainagentchangeissueconfirmvtwoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotrainagentchangeissueconfirmvtwoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChangeIssueRq is ChangeIssueRq Setter
// 改签出票回填入参rq
func (r *TaobaotrainagentchangeissueconfirmvtwoAPIRequest) SetChangeIssueRq(_changeIssueRq *ChangeIssueRq) error {
	r._changeIssueRq = _changeIssueRq
	r.Set("change_issue_rq", _changeIssueRq)
	return nil
}

// GetChangeIssueRq ChangeIssueRq Getter
func (r TaobaotrainagentchangeissueconfirmvtwoAPIRequest) GetChangeIssueRq() *ChangeIssueRq {
	return r._changeIssueRq
}
