package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
系列品商品变更-添加商品 
alibaba.wdk.series.sku.add

系列品商品变更-添加商品
*/
func AlibabaWdkSeriesSkuAdd(clt *core.SDKClient, req *wdk.AlibabaWdkSeriesSkuAddAPIRequest, session string) (*wdk.AlibabaWdkSeriesSkuAddAPIResponse, error) {
    var resp wdk.AlibabaWdkSeriesSkuAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
