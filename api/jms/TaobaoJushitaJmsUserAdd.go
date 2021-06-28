package jms

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jms"
)

/* 
添加ONS消息同步用户 
taobao.jushita.jms.user.add

添加ONS消息同步用户
*/
func TaobaoJushitaJmsUserAdd(clt *core.SDKClient, req *jms.TaobaoJushitaJmsUserAddRequest, session string) (*jms.TaobaoJushitaJmsUserAddAPIResponse, error) {
    var resp jms.TaobaoJushitaJmsUserAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
