package user

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/user"
)

/* 
Open Account索引查询 
taobao.open.account.index.find

Open Account索引查询
*/
func TaobaoOpenAccountIndexFind(clt *core.SDKClient, req *user.TaobaoOpenAccountIndexFindRequest, session string) (*user.TaobaoOpenAccountIndexFindAPIResponse, error) {
    var resp user.TaobaoOpenAccountIndexFindAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
