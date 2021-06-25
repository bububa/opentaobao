package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
商家注册更新路由信息 
alibaba.wdk.merchant.routing.register

商家注册更新路由信息
*/
func AlibabaWdkMerchantRoutingRegister(clt *core.SDKClient, req *wdk.AlibabaWdkMerchantRoutingRegisterRequest, session string) (*wdk.AlibabaWdkMerchantRoutingRegisterResponse, error) {
    var resp wdk.AlibabaWdkMerchantRoutingRegisterAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
