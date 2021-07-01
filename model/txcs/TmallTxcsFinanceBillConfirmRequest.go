package txcs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商账单确认 API请求
tmall.txcs.finance.bill.confirm

提供天猫超市外部合作商家：财务账单对账
*/
type TmallTxcsFinanceBillConfirmAPIRequest struct {
    model.Params
    // 门店ID
    _ouCode   string
    // 系统自动生成
    _statementBillConfirmDTO   *StatementBillConfirmDTO
}

// 初始化TmallTxcsFinanceBillConfirmAPIRequest对象
func NewTmallTxcsFinanceBillConfirmRequest() *TmallTxcsFinanceBillConfirmAPIRequest{
    return &TmallTxcsFinanceBillConfirmAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallTxcsFinanceBillConfirmAPIRequest) GetApiMethodName() string {
    return "tmall.txcs.finance.bill.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TmallTxcsFinanceBillConfirmAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuCode Setter
// 门店ID
func (r *TmallTxcsFinanceBillConfirmAPIRequest) SetOuCode(_ouCode string) error {
    r._ouCode = _ouCode
    r.Set("ou_code", _ouCode)
    return nil
}

// OuCode Getter
func (r TmallTxcsFinanceBillConfirmAPIRequest) GetOuCode() string {
    return r._ouCode
}
// StatementBillConfirmDTO Setter
// 系统自动生成
func (r *TmallTxcsFinanceBillConfirmAPIRequest) SetStatementBillConfirmDTO(_statementBillConfirmDTO *StatementBillConfirmDTO) error {
    r._statementBillConfirmDTO = _statementBillConfirmDTO
    r.Set("statement_bill_confirm_d_t_o", _statementBillConfirmDTO)
    return nil
}

// StatementBillConfirmDTO Getter
func (r TmallTxcsFinanceBillConfirmAPIRequest) GetStatementBillConfirmDTO() *StatementBillConfirmDTO {
    return r._statementBillConfirmDTO
}
