package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
新增渠道商品 
alibaba.wdk.sku.channelsku.add

盒马帮1期新增渠道商品
*/
func AlibabaWdkSkuChannelskuAdd(clt *core.SDKClient, req *wdk.AlibabaWdkSkuChannelskuAddRequest, session string) (*wdk.AlibabaWdkSkuChannelskuAddResponse, error) {
    var resp wdk.AlibabaWdkSkuChannelskuAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
