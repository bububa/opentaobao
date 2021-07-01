package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
溯源url透出 
alibaba.wdk.trace.url.get

根据shopId和skuCode返回商品溯源url
*/
func AlibabaWdkTraceUrlGet(clt *core.SDKClient, req *wdk.AlibabaWdkTraceUrlGetAPIRequest, session string) (*wdk.AlibabaWdkTraceUrlGetAPIResponse, error) {
    var resp wdk.AlibabaWdkTraceUrlGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
