package openim

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openim"
)

/* 
OPENIM群成员获取 
taobao.openim.tribe.getmembers

OPENIM群成员获取
*/
func TaobaoOpenimTribeGetmembers(clt *core.SDKClient, req *openim.TaobaoOpenimTribeGetmembersRequest, session string) (*openim.TaobaoOpenimTribeGetmembersResponse, error) {
    var resp openim.TaobaoOpenimTribeGetmembersAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
