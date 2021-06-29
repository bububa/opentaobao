package fundplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
储值卡激活 APIRequest
alibaba.fundplatform.cardorder.card.active

储值卡接货接口，可以通过外部订单号或者卡号进行批量激活。如果储值卡已经被激活过仍然幂等返回成功。资金平台保证批量激活时一定全部成功或全部失败。
如果批量激活储值卡时，如果部分储值卡处于已激活，部分储值卡处于未激活，则会返回失败
*/
type AlibabaFundplatformCardorderCardActiveRequest struct {
    model.Params

    // 入参对象
    paramCardActiveRequest   *CardActiveRequest 

}

func NewAlibabaFundplatformCardorderCardActiveRequest() *AlibabaFundplatformCardorderCardActiveRequest{
    return &AlibabaFundplatformCardorderCardActiveRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaFundplatformCardorderCardActiveRequest) GetApiMethodName() string {
    return "alibaba.fundplatform.cardorder.card.active"
}

func (r AlibabaFundplatformCardorderCardActiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaFundplatformCardorderCardActiveRequest) SetParamCardActiveRequest(paramCardActiveRequest *CardActiveRequest) error {
    r.paramCardActiveRequest = paramCardActiveRequest
    r.Set("param_card_active_request", paramCardActiveRequest)
    return nil
}

func (r AlibabaFundplatformCardorderCardActiveRequest) GetParamCardActiveRequest() *CardActiveRequest {
    return r.paramCardActiveRequest
}

