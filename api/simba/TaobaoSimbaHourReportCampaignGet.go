package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
计划维度小时报表获取 
taobao.simba.hour.report.campaign.get

计划维度小时报表获取
*/
func TaobaoSimbaHourReportCampaignGet(clt *core.SDKClient, req *simba.TaobaoSimbaHourReportCampaignGetRequest, session string) (*simba.TaobaoSimbaHourReportCampaignGetAPIResponse, error) {
    var resp simba.TaobaoSimbaHourReportCampaignGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
