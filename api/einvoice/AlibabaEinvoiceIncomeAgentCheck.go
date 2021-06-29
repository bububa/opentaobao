package einvoice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/einvoice"
)

/* 
agent注册校验 
alibaba.einvoice.income.agent.check

agent注册是，需要交易用户填写的agentId是否有效
*/
func AlibabaEinvoiceIncomeAgentCheck(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceIncomeAgentCheckRequest, session string) (*einvoice.AlibabaEinvoiceIncomeAgentCheckAPIResponse, error) {
    var resp einvoice.AlibabaEinvoiceIncomeAgentCheckAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
