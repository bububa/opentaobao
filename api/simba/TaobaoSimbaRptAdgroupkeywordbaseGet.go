package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
推广组下的词基础报表数据查询(明细数据不分类型查询) 
taobao.simba.rpt.adgroupkeywordbase.get

推广组下的词基础报表数据查询(明细数据不分类型查询)
*/
func TaobaoSimbaRptAdgroupkeywordbaseGet(clt *core.SDKClient, req *simba.TaobaoSimbaRptAdgroupkeywordbaseGetRequest, session string) (*simba.TaobaoSimbaRptAdgroupkeywordbaseGetAPIResponse, error) {
    var resp simba.TaobaoSimbaRptAdgroupkeywordbaseGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
