package txcs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫超市外部商家财务账单信息查询 APIRequest
tmall.txcs.finance.bill.query

提供天猫超市外部合作商家财务账单对账
*/
type TmallTxcsFinanceBillQueryRequest struct {
    model.Params

    // 对账单号
    statementBillQuery   *StatementBillQuery 

}

func NewTmallTxcsFinanceBillQueryRequest() *TmallTxcsFinanceBillQueryRequest{
    return &TmallTxcsFinanceBillQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TmallTxcsFinanceBillQueryRequest) GetApiMethodName() string {
    return "tmall.txcs.finance.bill.query"
}

func (r TmallTxcsFinanceBillQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallTxcsFinanceBillQueryRequest) SetStatementBillQuery(statementBillQuery *StatementBillQuery) error {
    r.statementBillQuery = statementBillQuery
    r.Set("statement_bill_query", statementBillQuery)
    return nil
}

func (r TmallTxcsFinanceBillQueryRequest) GetStatementBillQuery() *StatementBillQuery {
    return r.statementBillQuery
}

