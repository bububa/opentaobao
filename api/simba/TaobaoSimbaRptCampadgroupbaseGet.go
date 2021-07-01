package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
推广计划下的推广组报表基础数据查询(只有汇总数据，无分类类型) 
taobao.simba.rpt.campadgroupbase.get

推广计划下的推广组报表基础数据查询(只有汇总数据，无分类类型)
*/
func TaobaoSimbaRptCampadgroupbaseGet(clt *core.SDKClient, req *simba.TaobaoSimbaRptCampadgroupbaseGetAPIRequest, session string) (*simba.TaobaoSimbaRptCampadgroupbaseGetAPIResponse, error) {
    var resp simba.TaobaoSimbaRptCampadgroupbaseGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
