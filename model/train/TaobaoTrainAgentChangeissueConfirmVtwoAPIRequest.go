package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentChangeissueConfirmVtwoAPIRequest 火车票代理商接口-跑腿改签出票回填-含鉴权校验 API请求
// taobao.train.agent.changeissue.confirm.vtwo
//
// 火车票代理商接口-跑腿改签出票回填-含鉴权校验
type TaobaoTrainAgentChangeissueConfirmVtwoAPIRequest struct {
	model.Params
	// 改签出票回填入参rq
	_changeIssueRq *ChangeIssueRq
}

// NewTaobaoTrainAgentChangeissueConfirmVtwoRequest 初始化TaobaoTrainAgentChangeissueConfirmVtwoAPIRequest对象
func NewTaobaoTrainAgentChangeissueConfirmVtwoRequest() *TaobaoTrainAgentChangeissueConfirmVtwoAPIRequest {
	return &TaobaoTrainAgentChangeissueConfirmVtwoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentChangeissueConfirmVtwoAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.changeissue.confirm.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentChangeissueConfirmVtwoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTrainAgentChangeissueConfirmVtwoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChangeIssueRq is ChangeIssueRq Setter
// 改签出票回填入参rq
func (r *TaobaoTrainAgentChangeissueConfirmVtwoAPIRequest) SetChangeIssueRq(_changeIssueRq *ChangeIssueRq) error {
	r._changeIssueRq = _changeIssueRq
	r.Set("change_issue_rq", _changeIssueRq)
	return nil
}

// GetChangeIssueRq ChangeIssueRq Getter
func (r TaobaoTrainAgentChangeissueConfirmVtwoAPIRequest) GetChangeIssueRq() *ChangeIssueRq {
	return r._changeIssueRq
}
