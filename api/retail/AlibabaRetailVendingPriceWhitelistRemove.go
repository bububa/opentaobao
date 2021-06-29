package retail

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/retail"
)

/* 
价格管控白名单去除 
alibaba.retail.vending.price.whitelist.remove

商家价格管控白名单去除
*/
func AlibabaRetailVendingPriceWhitelistRemove(clt *core.SDKClient, req *retail.AlibabaRetailVendingPriceWhitelistRemoveRequest, session string) (*retail.AlibabaRetailVendingPriceWhitelistRemoveAPIResponse, error) {
    var resp retail.AlibabaRetailVendingPriceWhitelistRemoveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
