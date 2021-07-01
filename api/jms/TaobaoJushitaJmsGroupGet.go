package jms

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jms"
)

/* 
查询ONS分组 
taobao.jushita.jms.group.get

查询当前appkey在ONS中已有的分组
*/
func TaobaoJushitaJmsGroupGet(clt *core.SDKClient, req *jms.TaobaoJushitaJmsGroupGetAPIRequest, session string) (*jms.TaobaoJushitaJmsGroupGetAPIResponse, error) {
    var resp jms.TaobaoJushitaJmsGroupGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
