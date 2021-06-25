package jms

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jms"
)

/* 
查询某个用户是否同步消息 
taobao.jushita.jms.user.get

查询某个用户是否同步消息，只支持单个查询
*/
func TaobaoJushitaJmsUserGet(clt *core.SDKClient, req *jms.TaobaoJushitaJmsUserGetRequest, session string) (*jms.TaobaoJushitaJmsUserGetResponse, error) {
    var resp jms.TaobaoJushitaJmsUserGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
