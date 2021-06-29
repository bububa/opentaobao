package jstinteractive

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jstinteractive"
)

/* 
互动积分扣减接口 
taobao.jst.interactive.point.decrease

扣减用户的互动积分
*/
func TaobaoJstInteractivePointDecrease(clt *core.SDKClient, req *jstinteractive.TaobaoJstInteractivePointDecreaseRequest, session string) (*jstinteractive.TaobaoJstInteractivePointDecreaseAPIResponse, error) {
    var resp jstinteractive.TaobaoJstInteractivePointDecreaseAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
