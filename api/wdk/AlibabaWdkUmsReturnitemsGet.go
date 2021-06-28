package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
退货库位商品查询（退货出库辅助）-回流单 
alibaba.wdk.ums.returnitems.get

退货库位商品查询（退货出库辅助）-回流单
*/
func AlibabaWdkUmsReturnitemsGet(clt *core.SDKClient, req *wdk.AlibabaWdkUmsReturnitemsGetRequest, session string) (*wdk.AlibabaWdkUmsReturnitemsGetAPIResponse, error) {
    var resp wdk.AlibabaWdkUmsReturnitemsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
