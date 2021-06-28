package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
搜索人群实时报表 
taobao.simba.rtrpt.targetingtag.get

获取搜搜人群实时报表
*/
func TaobaoSimbaRtrptTargetingtagGet(clt *core.SDKClient, req *simba.TaobaoSimbaRtrptTargetingtagGetRequest, session string) (*simba.TaobaoSimbaRtrptTargetingtagGetAPIResponse, error) {
    var resp simba.TaobaoSimbaRtrptTargetingtagGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
