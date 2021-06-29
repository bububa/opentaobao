package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV查询门店审核状态 APIRequest
alibaba.alihealth.dental.store.audit.query

ISV查询门店审核状态
*/
type AlibabaAlihealthDentalStoreAuditQueryRequest struct {
    model.Params

    // 审核ID列表
    storeAuditIds   []int64 

}

func NewAlibabaAlihealthDentalStoreAuditQueryRequest() *AlibabaAlihealthDentalStoreAuditQueryRequest{
    return &AlibabaAlihealthDentalStoreAuditQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDentalStoreAuditQueryRequest) GetApiMethodName() string {
    return "alibaba.alihealth.dental.store.audit.query"
}

func (r AlibabaAlihealthDentalStoreAuditQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDentalStoreAuditQueryRequest) SetStoreAuditIds(storeAuditIds []int64) error {
    r.storeAuditIds = storeAuditIds
    r.Set("store_audit_ids", storeAuditIds)
    return nil
}

func (r AlibabaAlihealthDentalStoreAuditQueryRequest) GetStoreAuditIds() []int64 {
    return r.storeAuditIds
}

