package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
盘点结果单-回流单 
alibaba.wdk.ums.inventory.check.get

盘点结果单-回流单
*/
func AlibabaWdkUmsInventoryCheckGet(clt *core.SDKClient, req *wdk.AlibabaWdkUmsInventoryCheckGetAPIRequest, session string) (*wdk.AlibabaWdkUmsInventoryCheckGetAPIResponse, error) {
    var resp wdk.AlibabaWdkUmsInventoryCheckGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
