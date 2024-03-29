package txcs

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallTxcsFinanceBillQueryAPIRequest) Reset() {
	r._statementBillQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTxcsFinanceBillQueryAPIRequest) GetApiMethodName() string {
	return "tmall.txcs.finance.bill.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTxcsFinanceBillQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallTxcsFinanceBillQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStatementBillQuery is StatementBillQuery Setter
// 对账单号
func (r *TmallTxcsFinanceBillQueryAPIRequest) SetStatementBillQuery(_statementBillQuery *StatementBillQuery) error {
	r._statementBillQuery = _statementBillQuery
	r.Set("statement_bill_query", _statementBillQuery)
	return nil
}

// GetStatementBillQuery StatementBillQuery Getter
func (r TmallTxcsFinanceBillQueryAPIRequest) GetStatementBillQuery() *StatementBillQuery {
	return r._statementBillQuery
}

var poolTmallTxcsFinanceBillQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallTxcsFinanceBillQueryRequest()
	},
}

// GetTmallTxcsFinanceBillQueryRequest 从 sync.Pool 获取 TmallTxcsFinanceBillQueryAPIRequest
func GetTmallTxcsFinanceBillQueryAPIRequest() *TmallTxcsFinanceBillQueryAPIRequest {
	return poolTmallTxcsFinanceBillQueryAPIRequest.Get().(*TmallTxcsFinanceBillQueryAPIRequest)
}

// ReleaseTmallTxcsFinanceBillQueryAPIRequest 将 TmallTxcsFinanceBillQueryAPIRequest 放入 sync.Pool
func ReleaseTmallTxcsFinanceBillQueryAPIRequest(v *TmallTxcsFinanceBillQueryAPIRequest) {
	v.Reset()
	poolTmallTxcsFinanceBillQueryAPIRequest.Put(v)
}
