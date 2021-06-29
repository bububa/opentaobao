package uscesl

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/uscesl"
)

/* 
商品条码亮灯API 
taobao.uscesl.biz.item.light.up

亮灯API
*/
func TaobaoUsceslBizItemLightUp(clt *core.SDKClient, req *uscesl.TaobaoUsceslBizItemLightUpRequest, session string) (*uscesl.TaobaoUsceslBizItemLightUpAPIResponse, error) {
    var resp uscesl.TaobaoUsceslBizItemLightUpAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
