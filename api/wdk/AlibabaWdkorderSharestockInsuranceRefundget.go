package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
共享库存投保业务售后逆向订单数据获取 
alibaba.wdkorder.sharestock.insurance.refundget

共享库存投保业务售后逆向订单数据获取
*/
func AlibabaWdkorderSharestockInsuranceRefundget(clt *core.SDKClient, req *wdk.AlibabaWdkorderSharestockInsuranceRefundgetRequest, session string) (*wdk.AlibabaWdkorderSharestockInsuranceRefundgetAPIResponse, error) {
    var resp wdk.AlibabaWdkorderSharestockInsuranceRefundgetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
