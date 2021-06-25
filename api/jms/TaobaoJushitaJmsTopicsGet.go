package jms

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jms"
)

/* 
根据用户nick获取开通的消息列表 
taobao.jushita.jms.topics.get

根据用户nick获取开通的消息列表
*/
func TaobaoJushitaJmsTopicsGet(clt *core.SDKClient, req *jms.TaobaoJushitaJmsTopicsGetRequest, session string) (*jms.TaobaoJushitaJmsTopicsGetResponse, error) {
    var resp jms.TaobaoJushitaJmsTopicsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
