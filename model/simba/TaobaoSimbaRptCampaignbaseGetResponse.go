package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
推广计划报表基础数据对象 APIResponse
taobao.simba.rpt.campaignbase.get

推广计划报表基础数据对象
*/
type TaobaoSimbaRptCampaignbaseGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaRptCampaignbaseGetResponse `json:"taobao_simba_rpt_campaignbase_get_response,omitempty"`
}

type TaobaoSimbaRptCampaignbaseGetResponse struct {

    // 推广计划查询结果
    RptCampaignBaseList   string `json:"rpt_campaign_base_list,omitempty"`

}
