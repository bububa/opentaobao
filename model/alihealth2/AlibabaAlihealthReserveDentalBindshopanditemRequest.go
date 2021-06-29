package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
绑定门店信息，商品信息 APIRequest
alibaba.alihealth.reserve.dental.bindshopanditem

绑定门店信息，商品信息
*/
type AlibabaAlihealthReserveDentalBindshopanditemRequest struct {
    model.Params

    // bind_list
    bindList   []BindDTO 

}

func NewAlibabaAlihealthReserveDentalBindshopanditemRequest() *AlibabaAlihealthReserveDentalBindshopanditemRequest{
    return &AlibabaAlihealthReserveDentalBindshopanditemRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthReserveDentalBindshopanditemRequest) GetApiMethodName() string {
    return "alibaba.alihealth.reserve.dental.bindshopanditem"
}

func (r AlibabaAlihealthReserveDentalBindshopanditemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthReserveDentalBindshopanditemRequest) SetBindList(bindList []BindDTO) error {
    r.bindList = bindList
    r.Set("bind_list", bindList)
    return nil
}

func (r AlibabaAlihealthReserveDentalBindshopanditemRequest) GetBindList() []BindDTO {
    return r.bindList
}

