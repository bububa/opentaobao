package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
解绑商品信息 APIRequest
alibaba.alihealth.reserve.dental.unbinditem

绑定门店信息，商品信息
*/
type AlibabaAlihealthReserveDentalUnbinditemRequest struct {
    model.Params

    // 服务商门店id
    spStoreId   string 

    // 服务商商品id
    spItemId   string 

}

func NewAlibabaAlihealthReserveDentalUnbinditemRequest() *AlibabaAlihealthReserveDentalUnbinditemRequest{
    return &AlibabaAlihealthReserveDentalUnbinditemRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthReserveDentalUnbinditemRequest) GetApiMethodName() string {
    return "alibaba.alihealth.reserve.dental.unbinditem"
}

func (r AlibabaAlihealthReserveDentalUnbinditemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthReserveDentalUnbinditemRequest) SetSpStoreId(spStoreId string) error {
    r.spStoreId = spStoreId
    r.Set("sp_store_id", spStoreId)
    return nil
}

func (r AlibabaAlihealthReserveDentalUnbinditemRequest) GetSpStoreId() string {
    return r.spStoreId
}

func (r *AlibabaAlihealthReserveDentalUnbinditemRequest) SetSpItemId(spItemId string) error {
    r.spItemId = spItemId
    r.Set("sp_item_id", spItemId)
    return nil
}

func (r AlibabaAlihealthReserveDentalUnbinditemRequest) GetSpItemId() string {
    return r.spItemId
}

