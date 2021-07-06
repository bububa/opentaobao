package txcs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallTxcsFinanceBillCheckAPIRequest 天猫超市外部商家财务账单对账 API请求
// tmall.txcs.finance.bill.check
//
// 提供天猫超市外部合作商家财务账单对账
type TmallTxcsFinanceBillCheckAPIRequest struct {
	model.Params
	// 门店编码
	_ouCode string
	// 查询参数
	_statementBillFeeDetailQuery *StatementBillFeeDetailQuery
}

// NewTmallTxcsFinanceBillCheckRequest 初始化TmallTxcsFinanceBillCheckAPIRequest对象
func NewTmallTxcsFinanceBillCheckRequest() *TmallTxcsFinanceBillCheckAPIRequest {
	return &TmallTxcsFinanceBillCheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTxcsFinanceBillCheckAPIRequest) GetApiMethodName() string {
	return "tmall.txcs.finance.bill.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTxcsFinanceBillCheckAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOuCode is OuCode Setter
// 门店编码
func (r *TmallTxcsFinanceBillCheckAPIRequest) SetOuCode(_ouCode string) error {
	r._ouCode = _ouCode
	r.Set("ou_code", _ouCode)
	return nil
}

// GetOuCode OuCode Getter
func (r TmallTxcsFinanceBillCheckAPIRequest) GetOuCode() string {
	return r._ouCode
}

// SetStatementBillFeeDetailQuery is StatementBillFeeDetailQuery Setter
// 查询参数
func (r *TmallTxcsFinanceBillCheckAPIRequest) SetStatementBillFeeDetailQuery(_statementBillFeeDetailQuery *StatementBillFeeDetailQuery) error {
	r._statementBillFeeDetailQuery = _statementBillFeeDetailQuery
	r.Set("statement_bill_fee_detail_query", _statementBillFeeDetailQuery)
	return nil
}

// GetStatementBillFeeDetailQuery StatementBillFeeDetailQuery Getter
func (r TmallTxcsFinanceBillCheckAPIRequest) GetStatementBillFeeDetailQuery() *StatementBillFeeDetailQuery {
	return r._statementBillFeeDetailQuery
}
