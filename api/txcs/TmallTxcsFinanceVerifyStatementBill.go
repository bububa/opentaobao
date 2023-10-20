package txcs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/txcs"
)

// TmallTxcsFinanceVerifyStatementBill 供应商核销单录入
// tmall.txcs.finance.verify.statement.bill
//
// 供应商核销单录入
func TmallTxcsFinanceVerifyStatementBill(clt *core.SDKClient, req *txcs.TmallTxcsFinanceVerifyStatementBillAPIRequest, resp *txcs.TmallTxcsFinanceVerifyStatementBillAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
