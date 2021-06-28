package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
用户发起售中取消 
alibaba.wdk.channel.order.usercancel

用户发起售中取消
*/
func AlibabaWdkChannelOrderUsercancel(clt *core.SDKClient, req *wdk.AlibabaWdkChannelOrderUsercancelRequest, session string) (*wdk.AlibabaWdkChannelOrderUsercancelAPIResponse, error) {
    var resp wdk.AlibabaWdkChannelOrderUsercancelAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
