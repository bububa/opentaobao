package uscesl

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/uscesl"
)

/* 
价签LED等点亮 
taobao.uscesl.biz.light.up

价签LED等点亮
*/
func TaobaoUsceslBizLightUp(clt *core.SDKClient, req *uscesl.TaobaoUsceslBizLightUpAPIRequest, session string) (*uscesl.TaobaoUsceslBizLightUpAPIResponse, error) {
    var resp uscesl.TaobaoUsceslBizLightUpAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
