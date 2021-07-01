package txcs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商核销单录入 API请求
tmall.txcs.finance.verify.statement.bill

供应商核销单录入
*/
type TmallTxcsFinanceVerifyStatementBillAPIRequest struct {
    model.Params
    // 门店ID
    _ouCode   string
    // 核销单内容
    _verificationBillDTO   *VerificationBillDTO
}

// 初始化TmallTxcsFinanceVerifyStatementBillAPIRequest对象
func NewTmallTxcsFinanceVerifyStatementBillRequest() *TmallTxcsFinanceVerifyStatementBillAPIRequest{
    return &TmallTxcsFinanceVerifyStatementBillAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallTxcsFinanceVerifyStatementBillAPIRequest) GetApiMethodName() string {
    return "tmall.txcs.finance.verify.statement.bill"
}

// IRequest interface 方法, 获取API参数
func (r TmallTxcsFinanceVerifyStatementBillAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuCode Setter
// 门店ID
func (r *TmallTxcsFinanceVerifyStatementBillAPIRequest) SetOuCode(_ouCode string) error {
    r._ouCode = _ouCode
    r.Set("ou_code", _ouCode)
    return nil
}

// OuCode Getter
func (r TmallTxcsFinanceVerifyStatementBillAPIRequest) GetOuCode() string {
    return r._ouCode
}
// VerificationBillDTO Setter
// 核销单内容
func (r *TmallTxcsFinanceVerifyStatementBillAPIRequest) SetVerificationBillDTO(_verificationBillDTO *VerificationBillDTO) error {
    r._verificationBillDTO = _verificationBillDTO
    r.Set("verification_bill_d_t_o", _verificationBillDTO)
    return nil
}

// VerificationBillDTO Getter
func (r TmallTxcsFinanceVerifyStatementBillAPIRequest) GetVerificationBillDTO() *VerificationBillDTO {
    return r._verificationBillDTO
}
