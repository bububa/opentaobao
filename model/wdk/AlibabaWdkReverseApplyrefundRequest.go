package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
逆向申请接口 APIRequest
alibaba.wdk.reverse.applyrefund

逆向渲染
*/
type AlibabaWdkReverseApplyrefundRequest struct {
    model.Params

    // 入参
    paramApplyReverseReq   *ApplyReverseReq 

}

func NewAlibabaWdkReverseApplyrefundRequest() *AlibabaWdkReverseApplyrefundRequest{
    return &AlibabaWdkReverseApplyrefundRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkReverseApplyrefundRequest) GetApiMethodName() string {
    return "alibaba.wdk.reverse.applyrefund"
}

func (r AlibabaWdkReverseApplyrefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkReverseApplyrefundRequest) SetParamApplyReverseReq(paramApplyReverseReq *ApplyReverseReq) error {
    r.paramApplyReverseReq = paramApplyReverseReq
    r.Set("param_apply_reverse_req", paramApplyReverseReq)
    return nil
}

func (r AlibabaWdkReverseApplyrefundRequest) GetParamApplyReverseReq() *ApplyReverseReq {
    return r.paramApplyReverseReq
}

