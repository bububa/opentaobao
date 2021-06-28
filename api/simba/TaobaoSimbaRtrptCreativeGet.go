package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
获取创意实时报表数据 
taobao.simba.rtrpt.creative.get

获取创意实时报表数据
*/
func TaobaoSimbaRtrptCreativeGet(clt *core.SDKClient, req *simba.TaobaoSimbaRtrptCreativeGetRequest, session string) (*simba.TaobaoSimbaRtrptCreativeGetAPIResponse, error) {
    var resp simba.TaobaoSimbaRtrptCreativeGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
