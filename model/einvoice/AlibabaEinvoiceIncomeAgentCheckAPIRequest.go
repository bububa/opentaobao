package einvoice

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEinvoiceIncomeAgentCheckAPIRequest) Reset() {
	r._agentId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceIncomeAgentCheckAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.income.agent.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceIncomeAgentCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoiceIncomeAgentCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAgentId is AgentId Setter
// 阿里发票平台分配的agentId
func (r *AlibabaEinvoiceIncomeAgentCheckAPIRequest) SetAgentId(_agentId string) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r AlibabaEinvoiceIncomeAgentCheckAPIRequest) GetAgentId() string {
	return r._agentId
}

var poolAlibabaEinvoiceIncomeAgentCheckAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEinvoiceIncomeAgentCheckRequest()
	},
}

// GetAlibabaEinvoiceIncomeAgentCheckRequest 从 sync.Pool 获取 AlibabaEinvoiceIncomeAgentCheckAPIRequest
func GetAlibabaEinvoiceIncomeAgentCheckAPIRequest() *AlibabaEinvoiceIncomeAgentCheckAPIRequest {
	return poolAlibabaEinvoiceIncomeAgentCheckAPIRequest.Get().(*AlibabaEinvoiceIncomeAgentCheckAPIRequest)
}

// ReleaseAlibabaEinvoiceIncomeAgentCheckAPIRequest 将 AlibabaEinvoiceIncomeAgentCheckAPIRequest 放入 sync.Pool
func ReleaseAlibabaEinvoiceIncomeAgentCheckAPIRequest(v *AlibabaEinvoiceIncomeAgentCheckAPIRequest) {
	v.Reset()
	poolAlibabaEinvoiceIncomeAgentCheckAPIRequest.Put(v)
}
