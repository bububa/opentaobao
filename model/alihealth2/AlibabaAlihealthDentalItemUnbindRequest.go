package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV解绑商品 APIRequest
alibaba.alihealth.dental.item.unbind

ISV解绑商品
*/
type AlibabaAlihealthDentalItemUnbindRequest struct {
    model.Params

    // ISV门店ID
    spStoreId   string 

    // ISV商品ID
    spItemId   string 

}

func NewAlibabaAlihealthDentalItemUnbindRequest() *AlibabaAlihealthDentalItemUnbindRequest{
    return &AlibabaAlihealthDentalItemUnbindRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDentalItemUnbindRequest) GetApiMethodName() string {
    return "alibaba.alihealth.dental.item.unbind"
}

func (r AlibabaAlihealthDentalItemUnbindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDentalItemUnbindRequest) SetSpStoreId(spStoreId string) error {
    r.spStoreId = spStoreId
    r.Set("sp_store_id", spStoreId)
    return nil
}

func (r AlibabaAlihealthDentalItemUnbindRequest) GetSpStoreId() string {
    return r.spStoreId
}

func (r *AlibabaAlihealthDentalItemUnbindRequest) SetSpItemId(spItemId string) error {
    r.spItemId = spItemId
    r.Set("sp_item_id", spItemId)
    return nil
}

func (r AlibabaAlihealthDentalItemUnbindRequest) GetSpItemId() string {
    return r.spItemId
}

