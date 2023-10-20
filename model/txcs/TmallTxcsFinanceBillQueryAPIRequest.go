package txcs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmalltxcsfinancebillqueryAPIRequest 天猫超市外部商家财务账单信息查询 API请求
// tmall.txcs.finance.bill.query
//
// 提供天猫超市外部合作商家财务账单对账
type TmalltxcsfinancebillqueryAPIRequest struct {
	model.Params
	// 对账单号
	_statementBillQuery *StatementBillQuery
}

// NewTmalltxcsfinancebillqueryRequest 初始化TmalltxcsfinancebillqueryAPIRequest对象
func NewTmalltxcsfinancebillqueryRequest() *TmalltxcsfinancebillqueryAPIRequest {
	return &TmalltxcsfinancebillqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmalltxcsfinancebillqueryAPIRequest) GetApiMethodName() string {
	return "tmall.txcs.finance.bill.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmalltxcsfinancebillqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmalltxcsfinancebillqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStatementBillQuery is StatementBillQuery Setter
// 对账单号
func (r *TmalltxcsfinancebillqueryAPIRequest) SetStatementBillQuery(_statementBillQuery *StatementBillQuery) error {
	r._statementBillQuery = _statementBillQuery
	r.Set("statement_bill_query", _statementBillQuery)
	return nil
}

// GetStatementBillQuery StatementBillQuery Getter
func (r TmalltxcsfinancebillqueryAPIRequest) GetStatementBillQuery() *StatementBillQuery {
	return r._statementBillQuery
}
