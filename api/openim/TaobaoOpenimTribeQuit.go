package openim

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openim"
)

/* 
OPENIM群成员退出 
taobao.openim.tribe.quit

OPENIM群成员退出
*/
func TaobaoOpenimTribeQuit(clt *core.SDKClient, req *openim.TaobaoOpenimTribeQuitRequest, session string) (*openim.TaobaoOpenimTribeQuitAPIResponse, error) {
    var resp openim.TaobaoOpenimTribeQuitAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
