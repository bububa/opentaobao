package user

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/user"
)

/* 
open account数据搜索 
taobao.open.account.search

open account数据搜索
*/
func TaobaoOpenAccountSearch(clt *core.SDKClient, req *user.TaobaoOpenAccountSearchAPIRequest, session string) (*user.TaobaoOpenAccountSearchAPIResponse, error) {
    var resp user.TaobaoOpenAccountSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
