package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV查询绑定审核状态 API请求
alibaba.alihealth.dental.bind.audit.query

ISV查询绑定审核状态
*/
type AlibabaAlihealthDentalBindAuditQueryRequest struct {
    model.Params
    // 绑定ID列表
    _bindIds   []int64
}

// 初始化AlibabaAlihealthDentalBindAuditQueryRequest对象
func NewAlibabaAlihealthDentalBindAuditQueryRequest() *AlibabaAlihealthDentalBindAuditQueryRequest{
    return &AlibabaAlihealthDentalBindAuditQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDentalBindAuditQueryRequest) GetApiMethodName() string {
    return "alibaba.alihealth.dental.bind.audit.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDentalBindAuditQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BindIds Setter
// 绑定ID列表
func (r *AlibabaAlihealthDentalBindAuditQueryRequest) SetBindIds(_bindIds []int64) error {
    r._bindIds = _bindIds
    r.Set("bind_ids", _bindIds)
    return nil
}

// BindIds Getter
func (r AlibabaAlihealthDentalBindAuditQueryRequest) GetBindIds() []int64 {
    return r._bindIds
}
