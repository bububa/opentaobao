package txcs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallTxcsFinanceBillCheckAPIRequest
天猫超市外部商家财务账单对账 API请求
tmall.txcs.finance.bill.check

提供天猫超市外部合作商家财务账单对账 */
type TmallTxcsFinanceBillCheckAPIRequest struct {
	model.Params
	// 查询参数
	_statementBillFeeDetailQuery *StatementBillFeeDetailQuery
	// 门店编码
	_ouCode string
}

// New
