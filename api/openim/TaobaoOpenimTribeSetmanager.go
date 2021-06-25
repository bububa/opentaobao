package openim

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openim"
)

/* 
OPENIM群设置管理员 
taobao.openim.tribe.setmanager

OPENIM群设置管理员
*/
func TaobaoOpenimTribeSetmanager(clt *core.SDKClient, req *openim.TaobaoOpenimTribeSetmanagerRequest, session string) (*openim.TaobaoOpenimTribeSetmanagerResponse, error) {
    var resp openim.TaobaoOpenimTribeSetmanagerAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
