package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
接收授信结果通知 APIRequest
tmall.aliauto.autofinance.credit.receive

天猫汽车的金融业务场景中，需要接收外部ISV对用户授信申请的通知结果.
*/
type TmallAliautoAutofinanceCreditReceiveRequest struct {
    model.Params

    // 授信通知描述
    creditReceiveDto   *CreditReceiveDto 

}

func NewTmallAliautoAutofinanceCreditReceiveRequest() *TmallAliautoAutofinanceCreditReceiveRequest{
    return &TmallAliautoAutofinanceCreditReceiveRequest{
        Params: model.NewParams(),
    }
}

func (r TmallAliautoAutofinanceCreditReceiveRequest) GetApiMethodName() string {
    return "tmall.aliauto.autofinance.credit.receive"
}

func (r TmallAliautoAutofinanceCreditReceiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallAliautoAutofinanceCreditReceiveRequest) SetCreditReceiveDto(creditReceiveDto *CreditReceiveDto) error {
    r.creditReceiveDto = creditReceiveDto
    r.Set("credit_receive_dto", creditReceiveDto)
    return nil
}

func (r TmallAliautoAutofinanceCreditReceiveRequest) GetCreditReceiveDto() *CreditReceiveDto {
    return r.creditReceiveDto
}

