package jms

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jms"
)

/* 
删除ONS分组 
taobao.jushita.jms.group.delete

删除ONS分组
*/
func TaobaoJushitaJmsGroupDelete(clt *core.SDKClient, req *jms.TaobaoJushitaJmsGroupDeleteRequest, session string) (*jms.TaobaoJushitaJmsGroupDeleteAPIResponse, error) {
    var resp jms.TaobaoJushitaJmsGroupDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
