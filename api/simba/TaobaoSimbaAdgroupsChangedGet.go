package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
分页获取修改的推广组ID和修改时间 
taobao.simba.adgroups.changed.get

分页获取修改的推广组ID和修改时间
*/
func TaobaoSimbaAdgroupsChangedGet(clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupsChangedGetRequest, session string) (*simba.TaobaoSimbaAdgroupsChangedGetResponse, error) {
    var resp simba.TaobaoSimbaAdgroupsChangedGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
