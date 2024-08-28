package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceIncomeAgentCheck agent注册校验
// alibaba.einvoice.income.agent.check
//
// agent注册是，需要交易用户填写的agentId是否有效
func AlibabaEinvoiceIncomeAgentCheck(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceIncomeAgentCheckAPIRequest, resp *einvoice.AlibabaEinvoiceIncomeAgentCheckAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
