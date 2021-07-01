package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
北京小蜜蜂配作业回传 
alibaba.wdk.fulfill.dms.ebeecake.work.order.callback

北京小蜜蜂配作业回传。
*/
func AlibabaWdkFulfillDmsEbeecakeWorkOrderCallback(clt *core.SDKClient, req *wdk.AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest, session string) (*wdk.AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIResponse, error) {
    var resp wdk.AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
