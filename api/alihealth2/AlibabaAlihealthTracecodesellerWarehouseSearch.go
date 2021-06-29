package alihealth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth2"
)

/* 
查询仓库api 
alibaba.alihealth.tracecodeseller.warehouse.search

查询仓库api
*/
func AlibabaAlihealthTracecodesellerWarehouseSearch(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthTracecodesellerWarehouseSearchRequest, session string) (*alihealth2.AlibabaAlihealthTracecodesellerWarehouseSearchAPIResponse, error) {
    var resp alihealth2.AlibabaAlihealthTracecodesellerWarehouseSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
