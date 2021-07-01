package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
阿里通信芝麻订单通知 
alibaba.telecom.zhima.ordernotify.callback

商家通知阿里通信，芝麻订单状态，阿里通信侧进行代扣支付、发货并确认收货
*/
func AlibabaTelecomZhimaOrdernotifyCallback(clt *core.SDKClient, req *alicom.AlibabaTelecomZhimaOrdernotifyCallbackAPIRequest, session string) (*alicom.AlibabaTelecomZhimaOrdernotifyCallbackAPIResponse, error) {
    var resp alicom.AlibabaTelecomZhimaOrdernotifyCallbackAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
