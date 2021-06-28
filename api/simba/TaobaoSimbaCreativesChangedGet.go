package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
分页获取修改过的广告创意ID和修改时间 
taobao.simba.creatives.changed.get

分页获取修改过的广告创意ID和修改时间
*/
func TaobaoSimbaCreativesChangedGet(clt *core.SDKClient, req *simba.TaobaoSimbaCreativesChangedGetRequest, session string) (*simba.TaobaoSimbaCreativesChangedGetAPIResponse, error) {
    var resp simba.TaobaoSimbaCreativesChangedGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
