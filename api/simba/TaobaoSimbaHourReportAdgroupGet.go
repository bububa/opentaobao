package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
推广单元小时级别实时报表查询 
taobao.simba.hour.report.adgroup.get

推广单元小时级别实时报表查询
*/
func TaobaoSimbaHourReportAdgroupGet(clt *core.SDKClient, req *simba.TaobaoSimbaHourReportAdgroupGetRequest, session string) (*simba.TaobaoSimbaHourReportAdgroupGetAPIResponse, error) {
    var resp simba.TaobaoSimbaHourReportAdgroupGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
