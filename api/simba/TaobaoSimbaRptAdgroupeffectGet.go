package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
推广组效果报表数据对象 
taobao.simba.rpt.adgroupeffect.get

推广组效果报表数据对象
*/
func TaobaoSimbaRptAdgroupeffectGet(clt *core.SDKClient, req *simba.TaobaoSimbaRptAdgroupeffectGetRequest, session string) (*simba.TaobaoSimbaRptAdgroupeffectGetResponse, error) {
    var resp simba.TaobaoSimbaRptAdgroupeffectGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
