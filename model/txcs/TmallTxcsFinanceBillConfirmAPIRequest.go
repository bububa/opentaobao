package txcs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallTxcsFinanceBillConfirmAPIRequest
供应商账单确认 API请求
tmall.txcs.finance.bill.confirm

提供天猫超市外部合作商家：财务账单对账 */
type TmallTxcsFinanceBillConfirmAPIRequest struct {
	model.Params
	// 门店ID
	_ouCode string
	// 系统自动生成
	_statementBillConfirmDTO *StatementBillConfirmDto
}

// New
