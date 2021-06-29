package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
agent注册校验 APIResponse
alibaba.einvoice.income.agent.check

agent注册是，需要交易用户填写的agentId是否有效
*/
type AlibabaEinvoiceIncomeAgentCheckAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceIncomeAgentCheckResponse
}

type AlibabaEinvoiceIncomeAgentCheckResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_income_agent_check_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否调用成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
