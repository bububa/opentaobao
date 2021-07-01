package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
获取定向效果报表数据 
taobao.simba.rpt.targetingtageffect.get

获取定向效果报表数据
*/
func TaobaoSimbaRptTargetingtageffectGet(clt *core.SDKClient, req *simba.TaobaoSimbaRptTargetingtageffectGetAPIRequest, session string) (*simba.TaobaoSimbaRptTargetingtageffectGetAPIResponse, error) {
    var resp simba.TaobaoSimbaRptTargetingtageffectGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
