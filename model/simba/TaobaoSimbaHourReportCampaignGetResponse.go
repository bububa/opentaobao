package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
计划维度小时报表获取 APIResponse
taobao.simba.hour.report.campaign.get

计划维度小时报表获取
*/
type TaobaoSimbaHourReportCampaignGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_hour_report_campaign_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 11
    
    Results   *RtRptResultEntityDTO `json:"results,omitempty" xml:"