package txcs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallTxcsFinanceVerifyStatementBillAPIRequest
供应商核销单录入 API请求
tmall.txcs.finance.verify.statement.bill

供应商核销单录入 */
type TmallTxcsFinanceVerifyStatementBillAPIRequest struct {
	model.Params
	// 门店ID
	_ouCode string
	// 核销单内容
	_verificationBillDTO *VerificationBillDto
}

// New
