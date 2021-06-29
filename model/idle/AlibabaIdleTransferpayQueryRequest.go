package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼转账结果查询 APIRequest
alibaba.idle.transferpay.query

商家业务 转账支付的结果查询
*/
type AlibabaIdleTransferpayQueryRequest struct {
    model.Params

    // 入参
    param   *PayQueryRequest 

}

func NewAlibabaIdleTransferpayQueryRequest() *AlibabaIdleTransferpayQueryRequest{
    return &AlibabaIdleTransferpayQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleTransferpayQueryRequest) GetApiMethodName() string {
    return "alibaba.idle.transferpay.query"
}

func (r AlibabaIdleTransferpayQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleTransferpayQueryRequest) SetParam(param *PayQueryRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaIdleTransferpayQueryRequest) GetParam() *PayQueryRequest {
    return r.param
}

