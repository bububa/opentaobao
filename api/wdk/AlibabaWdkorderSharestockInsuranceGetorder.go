package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
共享库存订单投保消息获取 
alibaba.wdkorder.sharestock.insurance.getorder

共享库存订单投保消息获取
*/
func AlibabaWdkorderSharestockInsuranceGetorder(clt *core.SDKClient, req *wdk.AlibabaWdkorderSharestockInsuranceGetorderAPIRequest, session string) (*wdk.AlibabaWdkorderSharestockInsuranceGetorderAPIResponse, error) {
    var resp wdk.AlibabaWdkorderSharestockInsuranceGetorderAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
