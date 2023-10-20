package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoiceincomeagentcheckAPIRequest agent注册校验 API请求
// alibaba.einvoice.income.agent.check
//
// agent注册是，需要交易用户填写的agentId是否有效
type AlibabaeinvoiceincomeagentcheckAPIRequest struct {
	model.Params
	// 阿里发票平台分配的agentId
	_agentId string
}

// NewAlibabaeinvoiceincomeagentcheckRequest 初始化AlibabaeinvoiceincomeagentcheckAPIRequest对象
func NewAlibabaeinvoiceincomeagentcheckRequest() *AlibabaeinvoiceincomeagentcheckAPIRequest {
	return &AlibabaeinvoiceincomeagentcheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoiceincomeagentcheckAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.income.agent.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoiceincomeagentcheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoiceincomeagentcheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAgentId is AgentId Setter
// 阿里发票平台分配的agentId
func (r *AlibabaeinvoiceincomeagentcheckAPIRequest) SetAgentId(_agentId string) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r AlibabaeinvoiceincomeagentcheckAPIRequest) GetAgentId() string {
	return r._agentId
}
