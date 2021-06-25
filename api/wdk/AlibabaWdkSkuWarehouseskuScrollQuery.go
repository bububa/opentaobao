package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
仓商品遍历接口(游标) 
alibaba.wdk.sku.warehousesku.scroll.query

提供仓商品数据接口查询
*/
func AlibabaWdkSkuWarehouseskuScrollQuery(clt *core.SDKClient, req *wdk.AlibabaWdkSkuWarehouseskuScrollQueryRequest, session string) (*wdk.AlibabaWdkSkuWarehouseskuScrollQueryResponse, error) {
    var resp wdk.AlibabaWdkSkuWarehouseskuScrollQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
