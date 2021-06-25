package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推送订单 APIRequest
alibaba.ele.fengniao.order.push

推送淘宝订单至蜂鸟开放平台配送
*/
type AlibabaEleFengniaoOrderPushRequest struct {
    model.Params

    // 参数param
    param   *Param 

}

func NewAlibabaEleFengniaoOrderPushRequest() *AlibabaEleFengniaoOrderPushRequest{
    return &AlibabaEleFengniaoOrderPushRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleFengniaoOrderPushRequest) GetApiMethodName() string {
    return "alibaba.ele.fengniao.order.push"
}

func (r AlibabaEleFengniaoOrderPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleFengniaoOrderPushRequest) SetParam(param *Param) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaEleFengniaoOrderPushRequest) GetParam() *Param {
    return r.param
}

