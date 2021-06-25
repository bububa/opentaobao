package openim

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openim"
)

/* 
获取用户群列表 
taobao.openim.tribe.getalltribes

OPENIM群服务获取用户群列表
*/
func TaobaoOpenimTribeGetalltribes(clt *core.SDKClient, req *openim.TaobaoOpenimTribeGetalltribesRequest, session string) (*openim.TaobaoOpenimTribeGetalltribesResponse, error) {
    var resp openim.TaobaoOpenimTribeGetalltribesAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
