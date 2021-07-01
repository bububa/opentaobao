package wdkitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdkitem"
)

/* 
商家商品修改 
alibaba.wdk.item.merchantsku.update

商家商品修改
*/
func AlibabaWdkItemMerchantskuUpdate(clt *core.SDKClient, req *wdkitem.AlibabaWdkItemMerchantskuUpdateAPIRequest, session string) (*wdkitem.AlibabaWdkItemMerchantskuUpdateAPIResponse, error) {
    var resp wdkitem.AlibabaWdkItemMerchantskuUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
