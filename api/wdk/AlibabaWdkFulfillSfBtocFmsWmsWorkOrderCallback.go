package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
顺丰仓作业回传 
alibaba.wdk.fulfill.sf.btoc.fms.wms.work.order.callback

顺丰仓作业单回传接口
*/
func AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallback(clt *core.SDKClient, req *wdk.AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackRequest, session string) (*wdk.AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIResponse, error) {
    var resp wdk.AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
