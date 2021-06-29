package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流修改计划 APIResponse
taobao.feedflow.item.campaign.modify

信息流修改计划
*/
type TaobaoFeedflowItemCampaignModifyAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemCampaignModifyResponse
}

type TaobaoFeedflowItemCampaignModifyResponse struct {
    XMLName xml.Name `xml:"feedflow_item_campaign_modify_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *TaobaoFeedflowItemCampaignModifyResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
