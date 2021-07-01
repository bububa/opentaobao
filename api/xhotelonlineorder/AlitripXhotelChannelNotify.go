package xhotelonlineorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelonlineorder"
)

/* 
分销渠道各类通知接口 
alitrip.xhotel.channel.notify

分销渠道支付通知
*/
func AlitripXhotelChannelNotify(clt *core.SDKClient, req *xhotelonlineorder.AlitripXhotelChannelNotifyAPIRequest, session string) (*xhotelonlineorder.AlitripXhotelChannelNotifyAPIResponse, error) {
    var resp xhotelonlineorder.AlitripXhotelChannelNotifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
