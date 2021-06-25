package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
业务重新触发发码短信 APIRequest
taobao.vmarket.eticket.flow.resend

业务重新触发发码短信
*/
type TaobaoVmarketEticketFlowResendRequest struct {
    model.Params

    // 业务单号
    outerId   string 

    // 业务类型值，可联系淘宝业务运营取得具体值
    bizType   int64 

}

func NewTaobaoVmarketEticketFlowResendRequest() *TaobaoVmarketEticketFlowResendRequest{
    return &TaobaoVmarketEticketFlowResendRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoVmarketEticketFlowResendRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.flow.resend"
}

func (r TaobaoVmarketEticketFlowResendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoVmarketEticketFlowResendRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TaobaoVmarketEticketFlowResendRequest) GetOuterId() string {
    return r.outerId
}

func (r *TaobaoVmarketEticketFlowResendRequest) SetBizType(bizType int64) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

func (r TaobaoVmarketEticketFlowResendRequest) GetBizType() int64 {
    return r.bizType
}

