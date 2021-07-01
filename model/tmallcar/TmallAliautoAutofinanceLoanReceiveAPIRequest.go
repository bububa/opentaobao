package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
接收放款结果通知 API请求
tmall.aliauto.autofinance.loan.receive

天猫汽车的金融业务场景中，需要接收外部ISV对用户支用放款的通知结果
*/
type TmallAliautoAutofinanceLoanReceiveAPIRequest struct {
    model.Params
    // 接收外部金融结构的放款结果通知参数
    _loanReceiveDto   *LoanReceiveDto
}

// 初始化TmallAliautoAutofinanceLoanReceiveAPIRequest对象
func NewTmallAliautoAutofinanceLoanReceiveRequest() *TmallAliautoAutofinanceLoanReceiveAPIRequest{
    return &TmallAliautoAutofinanceLoanReceiveAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallAliautoAutofinanceLoanReceiveAPIRequest) GetApiMethodName() string {
    return "tmall.aliauto.autofinance.loan.receive"
}

// IRequest interface 方法, 获取API参数
func (r TmallAliautoAutofinanceLoanReceiveAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LoanReceiveDto Setter
// 接收外部金融结构的放款结果通知参数
func (r *TmallAliautoAutofinanceLoanReceiveAPIRequest) SetLoanReceiveDto(_loanReceiveDto *LoanReceiveDto) error {
    r._loanReceiveDto = _loanReceiveDto
    r.Set("loan_receive_dto", _loanReceiveDto)
    return nil
}

// LoanReceiveDto Getter
func (r TmallAliautoAutofinanceLoanReceiveAPIRequest) GetLoanReceiveDto() *LoanReceiveDto {
    return r._loanReceiveDto
}
