package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
组合商品更新接口 
alibaba.wdk.sku.combinesku.update

组合商品修改接口
*/
func AlibabaWdkSkuCombineskuUpdate(clt *core.SDKClient, req *wdk.AlibabaWdkSkuCombineskuUpdateRequest, session string) (*wdk.AlibabaWdkSkuCombineskuUpdateAPIResponse, error) {
    var resp wdk.AlibabaWdkSkuCombineskuUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
