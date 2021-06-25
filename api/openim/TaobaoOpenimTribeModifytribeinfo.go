package openim

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openim"
)

/* 
OPENIM群信息修改 
taobao.openim.tribe.modifytribeinfo

OPENIM群信息修改
*/
func TaobaoOpenimTribeModifytribeinfo(clt *core.SDKClient, req *openim.TaobaoOpenimTribeModifytribeinfoRequest, session string) (*openim.TaobaoOpenimTribeModifytribeinfoResponse, error) {
    var resp openim.TaobaoOpenimTribeModifytribeinfoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
