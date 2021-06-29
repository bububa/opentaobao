package retail

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/retail"
)

/* 
贩卖机价格修改白名单 
alibaba.retail.vending.price.whitelist.add

贩卖机价格修改白名单
*/
func AlibabaRetailVendingPriceWhitelistAdd(clt *core.SDKClient, req *retail.AlibabaRetailVendingPriceWhitelistAddRequest, session string) (*retail.AlibabaRetailVendingPriceWhitelistAddAPIResponse, error) {
    var resp retail.AlibabaRetailVendingPriceWhitelistAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
