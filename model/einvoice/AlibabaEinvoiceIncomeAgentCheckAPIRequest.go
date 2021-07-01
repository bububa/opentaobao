package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceIncomeAgentCheckAPIRequest
agent注册校验 API请求
alibaba.einvoice.income.agent.check

agent注册是，需要交易用户填写的agentId是否有效 */
type AlibabaEinvoiceIncomeAgentCheckAPIRequest struct {
	model.Params
	// 阿里发票平台分配的agentId
	_agentId string
}

// New
