package txcs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/txcs"
)

// TmallTxcsFinanceVerifyStatementBill 供应商核销单录入
// tmall.txcs.finance.verify.statement.bill
//
// 供应商核销单录入
func TmallTxcsFinanceVerifyStatementBill(clt *core.SDKClient, req *txcs.TmallTxcsFinanceVerifyStatementBillAPIRequest, session string) (*txcs.TmallTxcsFinanceVerifyStatementBillAPIResponse, error) {
	var resp txcs.TmallTxcsFinanceVerifyStatementBillAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
