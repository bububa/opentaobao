package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
接收放款结果通知 APIRequest
tmall.aliauto.autofinance.loan.receive

天猫汽车的金融业务场景中，需要接收外部ISV对用户支用放款的通知结果
*/
type TmallAliautoAutofinanceLoanReceiveRequest struct {
    model.Params

    // 接收外部金融结构的放款结果通知参数
    loanReceiveDto   *LoanReceiveDto 

}

func NewTmallAliautoAutofinanceLoanReceiveRequest() *TmallAliautoAutofinanceLoanReceiveRequest{
    return &TmallAliautoAutofinanceLoanReceiveRequest{
        Params: model.NewParams(),
    }
}

func (r TmallAliautoAutofinanceLoanReceiveRequest) GetApiMethodName() string {
    return "tmall.aliauto.autofinance.loan.receive"
}

func (r TmallAliautoAutofinanceLoanReceiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallAliautoAutofinanceLoanReceiveRequest) SetLoanReceiveDto(loanReceiveDto *LoanReceiveDto) error {
    r.loanReceiveDto = loanReceiveDto
    r.Set("loan_receive_dto", loanReceiveDto)
    return nil
}

func (r TmallAliautoAutofinanceLoanReceiveRequest) GetLoanReceiveDto() *LoanReceiveDto {
    return r.loanReceiveDto
}

