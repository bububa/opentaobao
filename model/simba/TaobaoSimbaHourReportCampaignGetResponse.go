package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
计划维度小时报表获取 APIResponse
taobao.simba.hour.report.campaign.get

计划维度小时报表获取
*/
type TaobaoSimbaHourReportCampaignGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaHourReportCampaignGetResponse `json:"taobao_simba_hour_report_campaign_get_response,omitempty"`
}

type TaobaoSimbaHourReportCampaignGetResponse struct {

    // 11
    Results   *RtRptResultEntityDTO `json:"results,omitempty"`

}
