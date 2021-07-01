package wdkitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdkitem"
)

/* 
品牌信息查询 
alibaba.wdk.item.brand.query

品牌信息查询
*/
func AlibabaWdkItemBrandQuery(clt *core.SDKClient, req *wdkitem.AlibabaWdkItemBrandQueryAPIRequest, session string) (*wdkitem.AlibabaWdkItemBrandQueryAPIResponse, error) {
    var resp wdkitem.AlibabaWdkItemBrandQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
