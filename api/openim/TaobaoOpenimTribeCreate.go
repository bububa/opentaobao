package openim

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openim"
)

/* 
创建群 
taobao.openim.tribe.create

创建一个openim的群
*/
func TaobaoOpenimTribeCreate(clt *core.SDKClient, req *openim.TaobaoOpenimTribeCreateAPIRequest, session string) (*openim.TaobaoOpenimTribeCreateAPIResponse, error) {
    var resp openim.TaobaoOpenimTribeCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
