package txcs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallTxcsFinanceBillQueryAPIRequest 天猫超市外部商家财务账单信息查询 API请求
// tmall.txcs.finance.bill.query
//
// 提供天猫超市外部合作商家财务账单对账
type TmallTxcsFinanceBillQueryAPIRequest struct {
	model.Params
	// 对账单号
	_statementBillQuery *StatementBillQuery
}

// NewTmallTxcsFinanceBillQueryRequest 初始化TmallTxcsFinanceBillQueryAPIRequest对象
func NewTmallTxcsFinanceBillQueryRequest() *TmallTxcsFinanceBillQueryAPIRequest {
	return &TmallTxcsFinanceBillQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTxcsFinanceBillQueryAPIRequest) GetApiMethodName() string {
	return "tmall.txcs.finance.bill.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTxcsFinanceBillQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StatementBillQuery Setter
// 对账单号
func (r *TmallTxcsFinanceBillQueryAPIRequest) SetStatementBillQuery(_statementBillQuery *StatementBillQuery) error {
	r._statementBillQuery = _statementBillQuery
	r.Set("statement_bill_query", _statementBillQuery)
	return nil
}

// Get StatementBillQuery Getter
func (r TmallTxcsFinanceBillQueryAPIRequest) GetStatementBillQuery() *StatementBillQuery {
	return r._statementBillQuery
}
