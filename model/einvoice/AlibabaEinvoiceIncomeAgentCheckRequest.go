package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
agent注册校验 APIRequest
alibaba.einvoice.income.agent.check

agent注册是，需要交易用户填写的agentId是否有效
*/
type AlibabaEinvoiceIncomeAgentCheckRequest struct {
    model.Params

    // 阿里发票平台分配的agentId
    agentId   string 

}

func NewAlibabaEinvoiceIncomeAgentCheckRequest() *AlibabaEinvoiceIncomeAgentCheckRequest{
    return &AlibabaEinvoiceIncomeAgentCheckRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceIncomeAgentCheckRequest) GetApiMethodName() string {
    return "alibaba.einvoice.income.agent.check"
}

func (r AlibabaEinvoiceIncomeAgentCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceIncomeAgentCheckRequest) SetAgentId(agentId string) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r AlibabaEinvoiceIncomeAgentCheckRequest) GetAgentId() string {
    return r.agentId
}

