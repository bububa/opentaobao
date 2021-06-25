package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
新增商品 
alibaba.wdk.sku.add

创建RT门店商品或DC商品
*/
func AlibabaWdkSkuAdd(clt *core.SDKClient, req *wdk.AlibabaWdkSkuAddRequest, session string) (*wdk.AlibabaWdkSkuAddResponse, error) {
    var resp wdk.AlibabaWdkSkuAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
