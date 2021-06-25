package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
品牌营销涉及到的商品的库存同步接口 
alibaba.wdk.bm.stock.publish

用于操作sku的库存
*/
func AlibabaWdkBmStockPublish(clt *core.SDKClient, req *wdk.AlibabaWdkBmStockPublishRequest, session string) (*wdk.AlibabaWdkBmStockPublishResponse, error) {
    var resp wdk.AlibabaWdkBmStockPublishAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
