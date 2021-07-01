package alihealth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth2"
)

/* 
查询商户门店，商品列表 
alibaba.alihealth.reserve.dental.storesanditems

查询商户门店，商品列表
*/
func AlibabaAlihealthReserveDentalStoresanditems(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthReserveDentalStoresanditemsAPIRequest, session string) (*alihealth2.AlibabaAlihealthReserveDentalStoresanditemsAPIResponse, error) {
    var resp alihealth2.AlibabaAlihealthReserveDentalStoresanditemsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
