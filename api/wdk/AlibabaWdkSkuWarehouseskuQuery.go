package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
仓商品查询接口(指定商品编码) 
alibaba.wdk.sku.warehousesku.query

提供指定仓商品编码查询
*/
func AlibabaWdkSkuWarehouseskuQuery(clt *core.SDKClient, req *wdk.AlibabaWdkSkuWarehouseskuQueryRequest, session string) (*wdk.AlibabaWdkSkuWarehouseskuQueryAPIResponse, error) {
    var resp wdk.AlibabaWdkSkuWarehouseskuQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
