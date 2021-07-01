package damai

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/damai"
)

/* 
大麦换验平台-第三方对外开放-票品接口pushItem 
alibaba.damai.mev.open.pushitem

开放接口 推送票品
*/
func AlibabaDamaiMevOpenPushitem(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenPushitemAPIRequest, session string) (*damai.AlibabaDamaiMevOpenPushitemAPIResponse, error) {
    var resp damai.AlibabaDamaiMevOpenPushitemAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
