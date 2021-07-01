package uscesl

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/uscesl"
)

/* 
激活AP价签通讯模块 
taobao.uscesl.biz.ap.activate

激活AP价签通讯模块
*/
func TaobaoUsceslBizApActivate(clt *core.SDKClient, req *uscesl.TaobaoUsceslBizApActivateAPIRequest, session string) (*uscesl.TaobaoUsceslBizApActivateAPIResponse, error) {
    var resp uscesl.TaobaoUsceslBizApActivateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
