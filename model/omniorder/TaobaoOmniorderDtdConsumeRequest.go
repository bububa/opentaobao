package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店自送对码进行核销 APIRequest
taobao.omniorder.dtd.consume

该接口根据传入的码及订单信息，如果码与订单一致，则对门店自送服务进行核销。
*/
type TaobaoOmniorderDtdConsumeRequest struct {
    model.Params

    // 核销信息
    paramDoor2doorConsumeRequest   *Door2doorConsumeRequest 

}

func NewTaobaoOmniorderDtdConsumeRequest() *TaobaoOmniorderDtdConsumeRequest{
    return &TaobaoOmniorderDtdConsumeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniorderDtdConsumeRequest) GetApiMethodName() string {
    return "taobao.omniorder.dtd.consume"
}

func (r TaobaoOmniorderDtdConsumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniorderDtdConsumeRequest) SetParamDoor2doorConsumeRequest(paramDoor2doorConsumeRequest *Door2doorConsumeRequest) error {
    r.paramDoor2doorConsumeRequest = paramDoor2doorConsumeRequest
    r.Set("param_door2door_consume_request", paramDoor2doorConsumeRequest)
    return nil
}

func (r TaobaoOmniorderDtdConsumeRequest) GetParamDoor2doorConsumeRequest() *Door2doorConsumeRequest {
    return r.paramDoor2doorConsumeRequest
}

