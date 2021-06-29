package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV新增/修改口腔门店 APIRequest
alibaba.alihealth.dental.store.insertorupdate

ISV新增/修改口腔门店
*/
type AlibabaAlihealthDentalStoreInsertorupdateRequest struct {
    model.Params

    // 门店
    store   *DentalOuterStoreRequest 

}

func NewAlibabaAlihealthDentalStoreInsertorupdateRequest() *AlibabaAlihealthDentalStoreInsertorupdateRequest{
    return &AlibabaAlihealthDentalStoreInsertorupdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDentalStoreInsertorupdateRequest) GetApiMethodName() string {
    return "alibaba.alihealth.dental.store.insertorupdate"
}

func (r AlibabaAlihealthDentalStoreInsertorupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDentalStoreInsertorupdateRequest) SetStore(store *DentalOuterStoreRequest) error {
    r.store = store
    r.Set("store", store)
    return nil
}

func (r AlibabaAlihealthDentalStoreInsertorupdateRequest) GetStore() *DentalOuterStoreRequest {
    return r.store
}

