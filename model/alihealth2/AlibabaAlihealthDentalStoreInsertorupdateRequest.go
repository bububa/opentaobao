package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV新增/修改口腔门店 API请求
alibaba.alihealth.dental.store.insertorupdate

ISV新增/修改口腔门店
*/
type AlibabaAlihealthDentalStoreInsertorupdateRequest struct {
    model.Params
    // 门店
    _store   *DentalOuterStoreRequest
}

// 初始化AlibabaAlihealthDentalStoreInsertorupdateRequest对象
func NewAlibabaAlihealthDentalStoreInsertorupdateRequest() *AlibabaAlihealthDentalStoreInsertorupdateRequest{
    return &AlibabaAlihealthDentalStoreInsertorupdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDentalStoreInsertorupdateRequest) GetApiMethodName() string {
    return "alibaba.alihealth.dental.store.insertorupdate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDentalStoreInsertorupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Store Setter
// 门店
func (r *AlibabaAlihealthDentalStoreInsertorupdateRequest) SetStore(_store *DentalOuterStoreRequest) error {
    r._store = _store
    r.Set("store", _store)
    return nil
}

// Store Getter
func (r AlibabaAlihealthDentalStoreInsertorupdateRequest) GetStore() *DentalOuterStoreRequest {
    return r._store
}
