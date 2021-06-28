package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
推广组下的创意报表效果数据查询(汇总数据，不分类型) 
taobao.simba.rpt.adgroupcreativeeffect.get

推广组下的创意报表效果数据查询(汇总数据，不分类型)
*/
func TaobaoSimbaRptAdgroupcreativeeffectGet(clt *core.SDKClient, req *simba.TaobaoSimbaRptAdgroupcreativeeffectGetRequest, session string) (*simba.TaobaoSimbaRptAdgroupcreativeeffectGetAPIResponse, error) {
    var resp simba.TaobaoSimbaRptAdgroupcreativeeffectGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
