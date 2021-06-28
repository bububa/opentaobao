package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
推送订单 
alibaba.ele.fengniao.order.push

推送淘宝订单至蜂鸟开放平台配送
*/
func AlibabaEleFengniaoOrderPush(clt *core.SDKClient, req *logistic.AlibabaEleFengniaoOrderPushRequest, session string) (*logistic.AlibabaEleFengniaoOrderPushAPIResponse, error) {
    var resp logistic.AlibabaEleFengniaoOrderPushAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
