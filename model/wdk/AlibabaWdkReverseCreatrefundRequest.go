package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
逆向提交 APIRequest
alibaba.wdk.reverse.creatrefund

逆向申请提交
*/
type AlibabaWdkReverseCreatrefundRequest struct {
    model.Params

    // CreateReverseReq
    paramCreateReverseReq   *CreateReverseReq 

}

func NewAlibabaWdkReverseCreatrefundRequest() *AlibabaWdkReverseCreatrefundRequest{
    return &AlibabaWdkReverseCreatrefundRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkReverseCreatrefundRequest) GetApiMethodName() string {
    return "alibaba.wdk.reverse.creatrefund"
}

func (r AlibabaWdkReverseCreatrefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkReverseCreatrefundRequest) SetParamCreateReverseReq(paramCreateReverseReq *CreateReverseReq) error {
    r.paramCreateReverseReq = paramCreateReverseReq
    r.Set("param_create_reverse_req", paramCreateReverseReq)
    return nil
}

func (r AlibabaWdkReverseCreatrefundRequest) GetParamCreateReverseReq() *CreateReverseReq {
    return r.paramCreateReverseReq
}

