package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
库调单-回流单 
alibaba.wdk.ums.inventory.adjust.get

库调单-回流单
*/
func AlibabaWdkUmsInventoryAdjustGet(clt *core.SDKClient, req *wdk.AlibabaWdkUmsInventoryAdjustGetAPIRequest, session string) (*wdk.AlibabaWdkUmsInventoryAdjustGetAPIResponse, error) {
    var resp wdk.AlibabaWdkUmsInventoryAdjustGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
