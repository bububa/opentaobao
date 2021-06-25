package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
蜂鸟询用户T APIRequest
alibaba.ele.fengniao.user.time.query

蜂鸟询用户T
*/
type AlibabaEleFengniaoUserTimeQueryRequest struct {
    model.Params

    // 询T入参
    param   *PredictDeliveryTimeParam 

}

func NewAlibabaEleFengniaoUserTimeQueryRequest() *AlibabaEleFengniaoUserTimeQueryRequest{
    return &AlibabaEleFengniaoUserTimeQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleFengniaoUserTimeQueryRequest) GetApiMethodName() string {
    return "alibaba.ele.fengniao.user.time.query"
}

func (r AlibabaEleFengniaoUserTimeQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleFengniaoUserTimeQueryRequest) SetParam(param *PredictDeliveryTimeParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaEleFengniaoUserTimeQueryRequest) GetParam() *PredictDeliveryTimeParam {
    return r.param
}

