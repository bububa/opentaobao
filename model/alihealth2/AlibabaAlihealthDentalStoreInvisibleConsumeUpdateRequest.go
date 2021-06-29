package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店无隐形消费签约 APIRequest
alibaba.alihealth.dental.store.invisible.consume.update

门店无隐形消费签约
*/
type AlibabaAlihealthDentalStoreInvisibleConsumeUpdateRequest struct {
    model.Params

    // 入参
    store   *DentalOuterStoreNicRequest 

}

func NewAlibabaAlihealthDentalStoreInvisibleConsumeUpdateRequest() *AlibabaAlihealthDentalStoreInvisibleConsumeUpdateRequest{
    return &AlibabaAlihealthDentalStoreInvisibleConsumeUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDentalStoreInvisibleConsumeUpdateRequest) GetApiMethodName() string {
    return "alibaba.alihealth.dental.store.invisible.consume.update"
}

func (r AlibabaAlihealthDentalStoreInvisibleConsumeUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDentalStoreInvisibleConsumeUpdateRequest) SetStore(store *DentalOuterStoreNicRequest) error {
    r.store = store
    r.Set("store", store)
    return nil
}

func (r AlibabaAlihealthDentalStoreInvisibleConsumeUpdateRequest) GetStore() *DentalOuterStoreNicRequest {
    return r.store
}

