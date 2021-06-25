package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
商户订单履约数据获取 
alibaba.wdkorder.sharestock.fulfill.get

商户订单履约数据获取
*/
func AlibabaWdkorderSharestockFulfillGet(clt *core.SDKClient, req *wdk.AlibabaWdkorderSharestockFulfillGetRequest, session string) (*wdk.AlibabaWdkorderSharestockFulfillGetResponse, error) {
    var resp wdk.AlibabaWdkorderSharestockFulfillGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
