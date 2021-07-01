package lstvending

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/lstvending"
)

/* 
售货机出货回传接口 
alibaba.lst.vending.shipping.callback

零售通自动售货机商品出货回传接口，同步商品出库最新状态。
*/
func AlibabaLstVendingShippingCallback(clt *core.SDKClient, req *lstvending.AlibabaLstVendingShippingCallbackAPIRequest, session string) (*lstvending.AlibabaLstVendingShippingCallbackAPIResponse, error) {
    var resp lstvending.AlibabaLstVendingShippingCallbackAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
