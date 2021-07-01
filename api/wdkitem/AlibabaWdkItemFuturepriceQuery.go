package wdkitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdkitem"
)

/* 
单个商品未来价查询接口 
alibaba.wdk.item.futureprice.query

查询单个商品未来价，融合了未来基础售价+未来促销价
*/
func AlibabaWdkItemFuturepriceQuery(clt *core.SDKClient, req *wdkitem.AlibabaWdkItemFuturepriceQueryAPIRequest, session string) (*wdkitem.AlibabaWdkItemFuturepriceQueryAPIResponse, error) {
    var resp wdkitem.AlibabaWdkItemFuturepriceQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
