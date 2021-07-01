package openim

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openim"
)

/* 
openim应用聊天记录查询 
taobao.openim.app.chatlogs.get

查询openim应用的聊天记录
*/
func TaobaoOpenimAppChatlogsGet(clt *core.SDKClient, req *openim.TaobaoOpenimAppChatlogsGetAPIRequest, session string) (*openim.TaobaoOpenimAppChatlogsGetAPIResponse, error) {
    var resp openim.TaobaoOpenimAppChatlogsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
