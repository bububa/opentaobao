package jst

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jst"
)

/* 
小程序活动创建 
taobao.jst.miniapp.crowd.create

小程序活动创建
*/
func TaobaoJstMiniappCrowdCreate(clt *core.SDKClient, req *jst.TaobaoJstMiniappCrowdCreateRequest, session string) (*jst.TaobaoJstMiniappCrowdCreateResponse, error) {
    var resp jst.TaobaoJstMiniappCrowdCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
