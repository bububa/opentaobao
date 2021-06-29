package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
通过计划id查询计划 APIResponse
taobao.feedflow.item.campaign.get

通过计划id查询计划
*/
type TaobaoFeedflowItemCampaignGetAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemCampaignGetResponse
}

type TaobaoFeedflowItemCampaignGetResponse struct {
    XMLName xml.Name `xml:"feedflow_item_campaign_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *TaobaoFeedflowItemCampaignGetResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
