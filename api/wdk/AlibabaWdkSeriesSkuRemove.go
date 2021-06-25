package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
系列品商品变更-移除商品 
alibaba.wdk.series.sku.remove

系列品商品变更-移除商品
*/
func AlibabaWdkSeriesSkuRemove(clt *core.SDKClient, req *wdk.AlibabaWdkSeriesSkuRemoveRequest, session string) (*wdk.AlibabaWdkSeriesSkuRemoveResponse, error) {
    var resp wdk.AlibabaWdkSeriesSkuRemoveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
