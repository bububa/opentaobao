package txcs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫超市外部商家财务账单对账 APIRequest
tmall.txcs.finance.bill.check

提供天猫超市外部合作商家财务账单对账
*/
type TmallTxcsFinanceBillCheckRequest struct {
    model.Params

    // 查询参数
    statementBillFeeDetailQuery   *StatementBillFeeDetailQuery 

    // 门店编码
    ouCode   string 

}

func NewTmallTxcsFinanceBillCheckRequest() *TmallTxcsFinanceBillCheckRequest{
    return &TmallTxcsFinanceBillCheckRequest{
        Params: model.NewParams(),
    }
}

func (r TmallTxcsFinanceBillCheckRequest) GetApiMethodName() string {
    return "tmall.txcs.finance.bill.check"
}

func (r TmallTxcsFinanceBillCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallTxcsFinanceBillCheckRequest) SetStatementBillFeeDetailQuery(statementBillFeeDetailQuery *StatementBillFeeDetailQuery) error {
    r.statementBillFeeDetailQuery = statementBillFeeDetailQuery
    r.Set("statement_bill_fee_detail_query", statementBillFeeDetailQuery)
    return nil
}

func (r TmallTxcsFinanceBillCheckRequest) GetStatementBillFeeDetailQuery() *StatementBillFeeDetailQuery {
    return r.statementBillFeeDetailQuery
}

func (r *TmallTxcsFinanceBillCheckRequest) SetOuCode(ouCode string) error {
    r.ouCode = ouCode
    r.Set("ou_code", ouCode)
    return nil
}

func (r TmallTxcsFinanceBillCheckRequest) GetOuCode() string {
    return r.ouCode
}

