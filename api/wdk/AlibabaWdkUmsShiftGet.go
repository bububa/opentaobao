package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
移库单获取 
alibaba.wdk.ums.shift.get

移库单获取
*/
func AlibabaWdkUmsShiftGet(clt *core.SDKClient, req *wdk.AlibabaWdkUmsShiftGetAPIRequest, session string) (*wdk.AlibabaWdkUmsShiftGetAPIResponse, error) {
    var resp wdk.AlibabaWdkUmsShiftGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
