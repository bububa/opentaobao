package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
流量钱包流量发放-面向支付宝用户 
alibaba.aliqin.flow.alipay.publish

用户淘宝流量钱包商家给支付宝用户发放流量
*/
func AlibabaAliqinFlowAlipayPublish(clt *core.SDKClient, req *alicom.AlibabaAliqinFlowAlipayPublishRequest, session string) (*alicom.AlibabaAliqinFlowAlipayPublishResponse, error) {
    var resp alicom.AlibabaAliqinFlowAlipayPublishAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
