package jstinteractive

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jstinteractive"
)

/* 
互动积分发放接口 
taobao.jst.interactive.point.increase

向用户发放互动积分
*/
func TaobaoJstInteractivePointIncrease(clt *core.SDKClient, req *jstinteractive.TaobaoJstInteractivePointIncreaseRequest, session string) (*jstinteractive.TaobaoJstInteractivePointIncreaseAPIResponse, error) {
    var resp jstinteractive.TaobaoJstInteractivePointIncreaseAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
