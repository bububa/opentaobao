package tmc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmc"
)

/* 
获取用户开通的topic列表 
taobao.tmc.user.topics.get

获取用户开通的topic列表，授权消息使用
*/
func TaobaoTmcUserTopicsGet(clt *core.SDKClient, req *tmc.TaobaoTmcUserTopicsGetRequest, session string) (*tmc.TaobaoTmcUserTopicsGetResponse, error) {
    var resp tmc.TaobaoTmcUserTopicsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
