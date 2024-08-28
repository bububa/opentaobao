package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAccountBudgetGet 查询日消耗预算
// alibaba.scbp.account.budget.get
//
// 查询日消耗预算
func AlibabaScbpAccountBudgetGet(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAccountBudgetGetAPIRequest, resp *scbp.AlibabaScbpAccountBudgetGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
