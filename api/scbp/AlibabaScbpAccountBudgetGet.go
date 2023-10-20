package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAccountBudgetGet 查询日消耗预算
// alibaba.scbp.account.budget.get
//
// 查询日消耗预算
func AlibabaScbpAccountBudgetGet(clt *core.SDKClient, req *scbp.AlibabaScbpAccountBudgetGetAPIRequest, resp *scbp.AlibabaScbpAccountBudgetGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
