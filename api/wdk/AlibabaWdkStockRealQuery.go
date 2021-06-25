package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
仓内实时库存查询 
alibaba.wdk.stock.real.query

查询仓内实时库存信息
*/
func AlibabaWdkStockRealQuery(clt *core.SDKClient, req *wdk.AlibabaWdkStockRealQueryRequest, session string) (*wdk.AlibabaWdkStockRealQueryResponse, error) {
    var resp wdk.AlibabaWdkStockRealQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
