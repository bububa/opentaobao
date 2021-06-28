package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
查询渠道商品 
alibaba.wdk.sku.channelsku.query

查询渠道商品
*/
func AlibabaWdkSkuChannelskuQuery(clt *core.SDKClient, req *wdk.AlibabaWdkSkuChannelskuQueryRequest, session string) (*wdk.AlibabaWdkSkuChannelskuQueryAPIResponse, error) {
    var resp wdk.AlibabaWdkSkuChannelskuQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
