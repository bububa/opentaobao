package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV查询绑定审核状态 APIRequest
alibaba.alihealth.dental.bind.audit.query

ISV查询绑定审核状态
*/
type AlibabaAlihealthDentalBindAuditQueryRequest struct {
    model.Params

    // 绑定ID列表
    bindIds   []int64 

}

func NewAlibabaAlihealthDentalBindAuditQueryRequest() *AlibabaAlihealthDentalBindAuditQueryRequest{
    return &AlibabaAlihealthDentalBindAuditQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDentalBindAuditQueryRequest) GetApiMethodName() string {
    return "alibaba.alihealth.dental.bind.audit.query"
}

func (r AlibabaAlihealthDentalBindAuditQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDentalBindAuditQueryRequest) SetBindIds(bindIds []int64) error {
    r.bindIds = bindIds
    r.Set("bind_ids", bindIds)
    return nil
}

func (r AlibabaAlihealthDentalBindAuditQueryRequest) GetBindIds() []int64 {
    return r.bindIds
}

