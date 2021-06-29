package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推广计划报表基础数据对象 APIResponse
taobao.simba.rpt.campaignbase.get

推广计划报表基础数据对象
*/
type TaobaoSimbaRptCampaignbaseGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaRptCampaignbaseGetResponse
}

type TaobaoSimbaRptCampaignbaseGetResponse struct {
    XMLName xml.Name `xml:"simba_rpt_campaignbase_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 推广计划查询结果
    
    RptCampaignBaseList   string `json:"rpt_campaign_base_list,omitempty" xml:"rpt_campaign_base_list,omitempty"`

    
}
