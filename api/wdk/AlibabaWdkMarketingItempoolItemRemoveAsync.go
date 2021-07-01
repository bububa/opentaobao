package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
商品池删除商品 
alibaba.wdk.marketing.itempool.item.remove.async

新模型下删除商品
*/
func AlibabaWdkMarketingItempoolItemRemoveAsync(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest, session string) (*wdk.AlibabaWdkMarketingItempoolItemRemoveAsyncAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingItempoolItemRemoveAsyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
