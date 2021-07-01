package wdkitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdkitem"
)

/* 
商家商品信息查询 
alibaba.wdk.item.merchantsku.query

商家商品信息查询
*/
func AlibabaWdkItemMerchantskuQuery(clt *core.SDKClient, req *wdkitem.AlibabaWdkItemMerchantskuQueryAPIRequest, session string) (*wdkitem.AlibabaWdkItemMerchantskuQueryAPIResponse, error) {
    var resp wdkitem.AlibabaWdkItemMerchantskuQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
