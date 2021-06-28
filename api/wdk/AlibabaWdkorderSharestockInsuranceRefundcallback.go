package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
共享库存逆向订单理赔单回传 
alibaba.wdkorder.sharestock.insurance.refundcallback

共享库存逆向订单理赔单回传
*/
func AlibabaWdkorderSharestockInsuranceRefundcallback(clt *core.SDKClient, req *wdk.AlibabaWdkorderSharestockInsuranceRefundcallbackRequest, session string) (*wdk.AlibabaWdkorderSharestockInsuranceRefundcallbackAPIResponse, error) {
    var resp wdk.AlibabaWdkorderSharestockInsuranceRefundcallbackAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
