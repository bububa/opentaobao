package txcs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫超市外部商家财务账单信息查询 API请求
tmall.txcs.finance.bill.query

提供天猫超市外部合作商家财务账单对账
*/
type TmallTxcsFinanceBillQueryRequest struct {
    model.Params
    // 对账单号
    _statementBillQuery   *StatementBillQuery
}

// 初始化TmallTxcsFinanceBillQueryRequest对象
func NewTmallTxcsFinanceBillQueryRequest() *TmallTxcsFinanceBillQueryRequest{
    return &TmallTxcsFinanceBillQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallTxcsFinanceBillQueryRequest) GetApiMethodName() string {
    return "tmall.txcs.finance.bill.query"
}

// IRequest interface 方法, 获取API参数
func (r TmallTxcsFinanceBillQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StatementBillQuery Setter
// 对账单号
func (r *TmallTxcsFinanceBillQueryRequest) SetStatementBillQuery(_statementBillQuery *StatementBillQuery) error {
    r._statementBillQuery = _statementBillQuery
    r.Set("statement_bill_query", _statementBillQuery)
    return nil
}

// StatementBillQuery Getter
func (r TmallTxcsFinanceBillQueryRequest) GetStatementBillQuery() *StatementBillQuery {
    return r._statementBillQuery
}
