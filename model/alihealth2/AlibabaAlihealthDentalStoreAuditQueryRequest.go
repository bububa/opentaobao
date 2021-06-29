package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV查询门店审核状态 API请求
alibaba.alihealth.dental.store.audit.query

ISV查询门店审核状态
*/
type AlibabaAlihealthDentalStoreAuditQueryRequest struct {
    model.Params
    // 审核ID列表
    storeAuditIds   []int64
}

// 初始化AlibabaAlihealthDentalStoreAuditQueryRequest对象
func NewAlibabaAlihealthDentalStoreAuditQueryRequest() *AlibabaAlihealthDentalStoreAuditQueryRequest{
    return &AlibabaAlihealthDentalStoreAuditQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDentalStoreAuditQueryRequest) GetApiMethodName() string {
    return "alibaba.alihealth.dental.store.audit.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDentalStoreAuditQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreAuditIds Setter
// 审核ID列表
func (r *AlibabaAlihealthDentalStoreAuditQueryRequest) SetStoreAuditIds(storeAuditIds []int64) error {
    r.storeAuditIds = storeAuditIds
    r.Set("store_audit_ids", storeAuditIds)
    return nil
}

// StoreAuditIds Getter
func (r AlibabaAlihealthDentalStoreAuditQueryRequest) GetStoreAuditIds() []int64 {
    return r.storeAuditIds
}
