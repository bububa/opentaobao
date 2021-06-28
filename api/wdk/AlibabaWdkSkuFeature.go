package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
商品标记接口 
alibaba.wdk.sku.feature

给淘鲜达商品属性之外的打标通用能力，满足商品一些特殊的需求，比如是否参加营销。
*/
func AlibabaWdkSkuFeature(clt *core.SDKClient, req *wdk.AlibabaWdkSkuFeatureRequest, session string) (*wdk.AlibabaWdkSkuFeatureAPIResponse, error) {
    var resp wdk.AlibabaWdkSkuFeatureAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
