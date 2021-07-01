package openim

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openim"
)

/* 
OPENIM群主动加入 
taobao.openim.tribe.join

OPENIM群主动加入
*/
func TaobaoOpenimTribeJoin(clt *core.SDKClient, req *openim.TaobaoOpenimTribeJoinAPIRequest, session string) (*openim.TaobaoOpenimTribeJoinAPIResponse, error) {
    var resp openim.TaobaoOpenimTribeJoinAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
