package txcs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商账单确认 APIRequest
tmall.txcs.finance.bill.confirm

提供天猫超市外部合作商家：财务账单对账
*/
type TmallTxcsFinanceBillConfirmRequest struct {
    model.Params

    // 门店ID
    ouCode   string 

    // 系统自动生成
    statementBillConfirmDTO   *StatementBillConfirmDto 

}

func NewTmallTxcsFinanceBillConfirmRequest() *TmallTxcsFinanceBillConfirmRequest{
    return &TmallTxcsFinanceBillConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r TmallTxcsFinanceBillConfirmRequest) GetApiMethodName() string {
    return "tmall.txcs.finance.bill.confirm"
}

func (r TmallTxcsFinanceBillConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallTxcsFinanceBillConfirmRequest) SetOuCode(ouCode string) error {
    r.ouCode = ouCode
    r.Set("ou_code", ouCode)
    return nil
}

func (r TmallTxcsFinanceBillConfirmRequest) GetOuCode() string {
    return r.ouCode
}

func (r *TmallTxcsFinanceBillConfirmRequest) SetStatementBillConfirmDTO(statementBillConfirmDTO *StatementBillConfirmDto) error {
    r.statementBillConfirmDTO = statementBillConfirmDTO
    r.Set("statement_bill_confirm_d_t_o", statementBillConfirmDTO)
    return nil
}

func (r TmallTxcsFinanceBillConfirmRequest) GetStatementBillConfirmDTO() *StatementBillConfirmDto {
    return r.statementBillConfirmDTO
}

