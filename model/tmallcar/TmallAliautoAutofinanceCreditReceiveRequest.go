package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
接收授信结果通知 API请求
tmall.aliauto.autofinance.credit.receive

天猫汽车的金融业务场景中，需要接收外部ISV对用户授信申请的通知结果.
*/
type TmallAliautoAutofinanceCreditReceiveAPIRequest struct {
    model.Params
    // 授信通知描述
    _creditReceiveDto   *CreditReceiveDTO
}

// 初始化TmallAliautoAutofinanceCreditReceiveAPIRequest对象
func NewTmallAliautoAutofinanceCreditReceiveRequest() *TmallAliautoAutofinanceCreditReceiveAPIRequest{
    return &TmallAliautoAutofinanceCreditReceiveAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallAliautoAutofinanceCreditReceiveAPIRequest) GetApiMethodName() string {
    return "tmall.aliauto.autofinance.credit.receive"
}

// IRequest interface 方法, 获取API参数
func (r TmallAliautoAutofinanceCreditReceiveAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CreditReceiveDto Setter
// 授信通知描述
func (r *TmallAliautoAutofinanceCreditReceiveAPIRequest) SetCreditReceiveDto(_creditReceiveDto *CreditReceiveDTO) error {
    r._creditReceiveDto = _creditReceiveDto
    r.Set("credit_receive_dto", _creditReceiveDto)
    return nil
}

// CreditReceiveDto Getter
func (r TmallAliautoAutofinanceCreditReceiveAPIRequest) GetCreditReceiveDto() *CreditReceiveDTO {
    return r._creditReceiveDto
}
