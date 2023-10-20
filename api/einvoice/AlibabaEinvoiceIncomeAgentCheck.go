package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoiceincomeagentcheck agent注册校验
// alibaba.einvoice.income.agent.check
//
// agent注册是，需要交易用户填写的agentId是否有效
func Alibabaeinvoiceincomeagentcheck(clt *core.SDKClient, req *einvoice.AlibabaeinvoiceincomeagentcheckAPIRequest, session string) (*einvoice.AlibabaeinvoiceincomeagentcheckAPIResponse, error) {
	var resp einvoice.AlibabaeinvoiceincomeagentcheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
