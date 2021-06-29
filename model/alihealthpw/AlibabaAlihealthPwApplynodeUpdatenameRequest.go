package alihealthpw

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
回调变更患者姓名 APIRequest
alibaba.alihealth.pw.applynode.updatename

回调变更患者姓名
*/
type AlibabaAlihealthPwApplynodeUpdatenameRequest struct {
    model.Params

    // 回调入参
    body   *ModifyNameRo 

}

func NewAlibabaAlihealthPwApplynodeUpdatenameRequest() *AlibabaAlihealthPwApplynodeUpdatenameRequest{
    return &AlibabaAlihealthPwApplynodeUpdatenameRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthPwApplynodeUpdatenameRequest) GetApiMethodName() string {
    return "alibaba.alihealth.pw.applynode.updatename"
}

func (r AlibabaAlihealthPwApplynodeUpdatenameRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthPwApplynodeUpdatenameRequest) SetBody(body *ModifyNameRo) error {
    r.body = body
    r.Set("body", body)
    return nil
}

func (r AlibabaAlihealthPwApplynodeUpdatenameRequest) GetBody() *ModifyNameRo {
    return r.body
}

