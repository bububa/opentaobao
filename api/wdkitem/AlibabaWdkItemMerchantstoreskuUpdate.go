package wdkitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdkitem"
)

/* 
门店商品信息修改 
alibaba.wdk.item.merchantstoresku.update

门店商品信息修改
*/
func AlibabaWdkItemMerchantstoreskuUpdate(clt *core.SDKClient, req *wdkitem.AlibabaWdkItemMerchantstoreskuUpdateAPIRequest, session string) (*wdkitem.AlibabaWdkItemMerchantstoreskuUpdateAPIResponse, error) {
    var resp wdkitem.AlibabaWdkItemMerchantstoreskuUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
