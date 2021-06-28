package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
获取推广组实时报表数据 
taobao.simba.rtrpt.adgroup.get

获取推广组实时报表数据
*/
func TaobaoSimbaRtrptAdgroupGet(clt *core.SDKClient, req *simba.TaobaoSimbaRtrptAdgroupGetRequest, session string) (*simba.TaobaoSimbaRtrptAdgroupGetAPIResponse, error) {
    var resp simba.TaobaoSimbaRtrptAdgroupGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
