package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
仓封箱回告 
alibaba.wdk.fulfill.warehouse.work.order.sealbox

仓封箱回告箱与包裹的关系
*/
func AlibabaWdkFulfillWarehouseWorkOrderSealbox(clt *core.SDKClient, req *wdk.AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIRequest, session string) (*wdk.AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponse, error) {
    var resp wdk.AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
