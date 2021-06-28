package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
更新渠道商品 
alibaba.wdk.sku.channelsku.update

批量更新渠道商品，商家通过Top接入
*/
func AlibabaWdkSkuChannelskuUpdate(clt *core.SDKClient, req *wdk.AlibabaWdkSkuChannelskuUpdateRequest, session string) (*wdk.AlibabaWdkSkuChannelskuUpdateAPIResponse, error) {
    var resp wdk.AlibabaWdkSkuChannelskuUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
