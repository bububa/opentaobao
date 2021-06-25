package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
账户级别小时报表获取 
taobao.simba.hour.report.account.get

获取账户小时实时报表数据
*/
func TaobaoSimbaHourReportAccountGet(clt *core.SDKClient, req *simba.TaobaoSimbaHourReportAccountGetRequest, session string) (*simba.TaobaoSimbaHourReportAccountGetResponse, error) {
    var resp simba.TaobaoSimbaHourReportAccountGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
