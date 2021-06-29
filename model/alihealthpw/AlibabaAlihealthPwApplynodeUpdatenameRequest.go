package alihealthpw

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
回调变更患者姓名 API请求
alibaba.alihealth.pw.applynode.updatename

回调变更患者姓名
*/
type AlibabaAlihealthPwApplynodeUpdatenameRequest struct {
    model.Params
    // 回调入参
    body   *ModifyNameRo
}

// 初始化AlibabaAlihealthPwApplynodeUpdatenameRequest对象
func NewAlibabaAlihealthPwApplynodeUpdatenameRequest() *AlibabaAlihealthPwApplynodeUpdatenameRequest{
    return &AlibabaAlihealthPwApplynodeUpdatenameRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPwApplynodeUpdatenameRequest) GetApiMethodName() string {
    return "alibaba.alihealth.pw.applynode.updatename"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPwApplynodeUpdatenameRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Body Setter
// 回调入参
func (r *AlibabaAlihealthPwApplynodeUpdatenameRequest) SetBody(body *ModifyNameRo) error {
    r.body = body
    r.Set("body", body)
    return nil
}

// Body Getter
func (r AlibabaAlihealthPwApplynodeUpdatenameRequest) GetBody() *ModifyNameRo {
    return r.body
}
