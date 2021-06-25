package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
智能硬件会员判断 
tmall.device.member.identity.get

用来识别该用户是否是商家会员·
*/
func TmallDeviceMemberIdentityGet(clt *core.SDKClient, req *iot.TmallDeviceMemberIdentityGetRequest, session string) (*iot.TmallDeviceMemberIdentityGetResponse, error) {
    var resp iot.TmallDeviceMemberIdentityGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
