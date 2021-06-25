package elife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌惠卡券核销 APIRequest
taobao.elife.lifecard.consume

用户线上购买生活汇品牌惠虚拟消费卡，线下购物时，商家码枪核销，涉及用户虚拟卡余额扣减操作
*/
type TaobaoElifeLifecardConsumeRequest struct {
    model.Params

    // 交易请求参数
    consumeRequest   *ConsumeRequest 

}

func NewTaobaoElifeLifecardConsumeRequest() *TaobaoElifeLifecardConsumeRequest{
    return &TaobaoElifeLifecardConsumeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoElifeLifecardConsumeRequest) GetApiMethodName() string {
    return "taobao.elife.lifecard.consume"
}

func (r TaobaoElifeLifecardConsumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoElifeLifecardConsumeRequest) SetConsumeRequest(consumeRequest *ConsumeRequest) error {
    r.consumeRequest = consumeRequest
    r.Set("consume_request", consumeRequest)
    return nil
}

func (r TaobaoElifeLifecardConsumeRequest) GetConsumeRequest() *ConsumeRequest {
    return r.consumeRequest
}

