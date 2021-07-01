package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
获取城市维度报表 
taobao.simba.report.city.get

获取城市维度报表
*/
func TaobaoSimbaReportCityGet(clt *core.SDKClient, req *simba.TaobaoSimbaReportCityGetAPIRequest, session string) (*simba.TaobaoSimbaReportCityGetAPIResponse, error) {
    var resp simba.TaobaoSimbaReportCityGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
