package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
获取智能硬件旗舰店入会码 
tmall.device.brand.memberurl.get

获取旗舰店在智能硬件上的入会码
*/
func TmallDeviceBrandMemberurlGet(clt *core.SDKClient, req *iot.TmallDeviceBrandMemberurlGetAPIRequest, session string) (*iot.TmallDeviceBrandMemberurlGetAPIResponse, error) {
    var resp iot.TmallDeviceBrandMemberurlGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
