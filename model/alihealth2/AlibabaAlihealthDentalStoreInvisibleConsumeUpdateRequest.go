package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店无隐形消费签约 API请求
alibaba.alihealth.dental.store.invisible.consume.update

门店无隐形消费签约
*/
type AlibabaAlihealthDentalStoreInvisibleConsumeUpdateRequest struct {
    model.Params
    // 入参
    _store   *DentalOuterStoreNicRequest
}

// 初始化AlibabaAlihealthDentalStoreInvisibleConsumeUpdateRequest对象
func NewAlibabaAlihealthDentalStoreInvisibleConsumeUpdateRequest() *AlibabaAlihealthDentalStoreInvisibleConsumeUpdateRequest{
    return &AlibabaAlihealthDentalStoreInvisibleConsumeUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDentalStoreInvisibleConsumeUpdateRequest) GetApiMethodName() string {
    return "alibaba.alihealth.dental.store.invisible.consume.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDentalStoreInvisibleConsumeUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Store Setter
// 入参
func (r *AlibabaAlihealthDentalStoreInvisibleConsumeUpdateRequest) SetStore(_store *DentalOuterStoreNicRequest) error {
    r._store = _store
    r.Set("store", _store)
    return nil
}

// Store Getter
func (r AlibabaAlihealthDentalStoreInvisibleConsumeUpdateRequest) GetStore() *DentalOuterStoreNicRequest {
    return r._store
}
