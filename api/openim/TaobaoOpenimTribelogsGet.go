package openim

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openim"
)

/* 
openim 群聊天记录导出接口 
taobao.openim.tribelogs.get

获取openim账号的群聊天记录
*/
func TaobaoOpenimTribelogsGet(clt *core.SDKClient, req *openim.TaobaoOpenimTribelogsGetAPIRequest, session string) (*openim.TaobaoOpenimTribelogsGetAPIResponse, error) {
    var resp openim.TaobaoOpenimTribelogsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
