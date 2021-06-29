package nrt

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/nrt"
)

/* 
商品库存查询接口 
alibaba.ascp.inventory.query

新零售联盟商家库存查询接口，用于商家商品库存数量感知查询
*/
func AlibabaAscpInventoryQuery(clt *core.SDKClient, req *nrt.AlibabaAscpInventoryQueryRequest, session string) (*nrt.AlibabaAscpInventoryQueryAPIResponse, error) {
    var resp nrt.AlibabaAscpInventoryQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
