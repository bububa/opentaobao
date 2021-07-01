package alihealth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth2"
)

/* 
ISV绑定外部门店id和外部商品id 
alibaba.alihealth.dental.item.bind

ISV绑定外部门店id和外部商品id
*/
func AlibabaAlihealthDentalItemBind(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthDentalItemBindAPIRequest, session string) (*alihealth2.AlibabaAlihealthDentalItemBindAPIResponse, error) {
    var resp alihealth2.AlibabaAlihealthDentalItemBindAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
