package alihealthpw

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
申请节点变更回调 API请求
alibaba.alihealth.pw.applynode.update

基金会回调阿里健康更新药品援助申请单的状态
*/
type AlibabaAlihealthPwApplynodeUpdateRequest struct {
    model.Params
    // 回调入参
    body   *AuditRollbackRo
}

// 初始化AlibabaAlihealthPwApplynodeUpdateRequest对象
func NewAlibabaAlihealthPwApplynodeUpdateRequest() *AlibabaAlihealthPwApplynodeUpdateRequest{
    return &AlibabaAlihealthPwApplynodeUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPwApplynodeUpdateRequest) GetApiMethodName() string {
    return "alibaba.alihealth.pw.applynode.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPwApplynodeUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Body Setter
// 回调入参
func (r *AlibabaAlihealthPwApplynodeUpdateRequest) SetBody(body *AuditRollbackRo) error {
    r.body = body
    r.Set("body", body)
    return nil
}

// Body Getter
func (r AlibabaAlihealthPwApplynodeUpdateRequest) GetBody() *AuditRollbackRo {
    return r.body
}
