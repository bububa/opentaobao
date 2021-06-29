package txcs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商核销单录入 APIRequest
tmall.txcs.finance.verify.statement.bill

供应商核销单录入
*/
type TmallTxcsFinanceVerifyStatementBillRequest struct {
    model.Params

    // 门店ID
    ouCode   string 

    // 核销单内容
    verificationBillDTO   *VerificationBillDto 

}

func NewTmallTxcsFinanceVerifyStatementBillRequest() *TmallTxcsFinanceVerifyStatementBillRequest{
    return &TmallTxcsFinanceVerifyStatementBillRequest{
        Params: model.NewParams(),
    }
}

func (r TmallTxcsFinanceVerifyStatementBillRequest) GetApiMethodName() string {
    return "tmall.txcs.finance.verify.statement.bill"
}

func (r TmallTxcsFinanceVerifyStatementBillRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallTxcsFinanceVerifyStatementBillRequest) SetOuCode(ouCode string) error {
    r.ouCode = ouCode
    r.Set("ou_code", ouCode)
    return nil
}

func (r TmallTxcsFinanceVerifyStatementBillRequest) GetOuCode() string {
    return r.ouCode
}

func (r *TmallTxcsFinanceVerifyStatementBillRequest) SetVerificationBillDTO(verificationBillDTO *VerificationBillDto) error {
    r.verificationBillDTO = verificationBillDTO
    r.Set("verification_bill_d_t_o", verificationBillDTO)
    return nil
}

func (r TmallTxcsFinanceVerifyStatementBillRequest) GetVerificationBillDTO() *VerificationBillDto {
    return r.verificationBillDTO
}

