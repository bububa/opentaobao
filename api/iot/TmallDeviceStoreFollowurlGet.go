package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
获取店铺关注链接 
tmall.device.store.followurl.get

获取智能硬件上的关注店铺的URL
*/
func TmallDeviceStoreFollowurlGet(clt *core.SDKClient, req *iot.TmallDeviceStoreFollowurlGetAPIRequest, session string) (*iot.TmallDeviceStoreFollowurlGetAPIResponse, error) {
    var resp iot.TmallDeviceStoreFollowurlGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
