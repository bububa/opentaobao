package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
仓内实时库位库存查询 
alibaba.wdk.stock.cabinet.query

查询仓内实时库位库存信息
*/
func AlibabaWdkStockCabinetQuery(clt *core.SDKClient, req *wdk.AlibabaWdkStockCabinetQueryRequest, session string) (*wdk.AlibabaWdkStockCabinetQueryResponse, error) {
    var resp wdk.AlibabaWdkStockCabinetQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
