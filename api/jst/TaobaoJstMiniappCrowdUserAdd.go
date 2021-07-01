package jst

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jst"
)

/* 
小程序添加用户到指定的活动 
taobao.jst.miniapp.crowd.user.add

小程序添加用户到指定的活动
*/
func TaobaoJstMiniappCrowdUserAdd(clt *core.SDKClient, req *jst.TaobaoJstMiniappCrowdUserAddAPIRequest, session string) (*jst.TaobaoJstMiniappCrowdUserAddAPIResponse, error) {
    var resp jst.TaobaoJstMiniappCrowdUserAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
