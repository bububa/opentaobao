package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
根据shopId和skuCode返回商品静态溯源url 
alibaba.wdk.item.trace.url.get

根据shopId和skuCode返回商品静态溯源url
*/
func AlibabaWdkItemTraceUrlGet(clt *core.SDKClient, req *wdk.AlibabaWdkItemTraceUrlGetAPIRequest, session string) (*wdk.AlibabaWdkItemTraceUrlGetAPIResponse, error) {
    var resp wdk.AlibabaWdkItemTraceUrlGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
