package alihealth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth2"
)

/* 
解绑商品信息 
alibaba.alihealth.reserve.dental.unbinditem

绑定门店信息，商品信息
*/
func AlibabaAlihealthReserveDentalUnbinditem(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthReserveDentalUnbinditemRequest, session string) (*alihealth2.AlibabaAlihealthReserveDentalUnbinditemAPIResponse, error) {
    var resp alihealth2.AlibabaAlihealthReserveDentalUnbinditemAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
