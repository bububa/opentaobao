package txcs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallTxcsFinanceBillQueryAPIRequest
天猫超市外部商家财务账单信息查询 API请求
tmall.txcs.finance.bill.query

提供天猫超市外部合作商家财务账单对账 */
type TmallTxcsFinanceBillQueryAPIRequest struct {
	model.Params
	// 对账单号
	_statementBillQuery *StatementBillQuery
}

// New
