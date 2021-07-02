package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceIncomeAgentCheckAPIRequest agent注册校验 API请求
// alibaba.einvoice.income.agent.check
//
// agent注册是，需要交易用户填写的agentId是否有效
type AlibabaEinvoiceIncomeAgentCheckAPIRequest struct {
	model.Params
	// 阿里发票平台分配的agentId
	_agentId string
}

// NewAlibabaEinvoiceIncomeAgentCheckRequest 初始化AlibabaEinvoiceIncomeAgentCheckAPIRequest对象
func NewAlibabaEinvoiceIncomeAgentCheckRequest() *AlibabaEinvoiceIncomeAgentCheckAPIRequest {
	return &AlibabaEinvoiceIncomeAgentCheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceIncomeAgentCheckAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.income.agent.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceIncomeAgentCheckAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AgentId Setter
// 阿里发票平台分配的agentId
func (r *AlibabaEinvoiceIncomeAgentCheckAPIRequest) SetAgentId(_agentId string) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// Get AgentId Getter
func (r AlibabaEinvoiceIncomeAgentCheckAPIRequest) GetAgentId() string {
	return r._agentId
}
