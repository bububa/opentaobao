package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推广计划分日数据查询 APIResponse
taobao.feedflow.item.campaign.rptdailylist

推广计划分日数据查询
*/
type TaobaoFeedflowItemCampaignRptdailylistAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemCampaignRptdailylistResponse
}

type TaobaoFeedflowItemCampaignRptdailylistResponse struct {
    XMLName xml.Name `xml:"feedflow_item_campaign_rptdailylist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TaobaoFeedflowItemCampaignRptdailylistResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
