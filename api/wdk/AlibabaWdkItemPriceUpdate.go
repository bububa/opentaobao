package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
五道口商品中心价格修改 
alibaba.wdk.item.price.update

修改门店商品售价和会员价
*/
func AlibabaWdkItemPriceUpdate(clt *core.SDKClient, req *wdk.AlibabaWdkItemPriceUpdateAPIRequest, session string) (*wdk.AlibabaWdkItemPriceUpdateAPIResponse, error) {
    var resp wdk.AlibabaWdkItemPriceUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
