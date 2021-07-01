package openim

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openim"
)

/* 
获取openim账号的聊天关系 
taobao.openim.relations.get

获取openim账号的聊天关系
*/
func TaobaoOpenimRelationsGet(clt *core.SDKClient, req *openim.TaobaoOpenimRelationsGetAPIRequest, session string) (*openim.TaobaoOpenimRelationsGetAPIResponse, error) {
    var resp openim.TaobaoOpenimRelationsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
