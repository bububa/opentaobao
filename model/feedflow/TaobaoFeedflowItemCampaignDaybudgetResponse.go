package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取当日投放日预算总额 APIResponse
taobao.feedflow.item.campaign.daybudget

获取当日投放日预算总额
*/
type TaobaoFeedflowItemCampaignDaybudgetAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemCampaignDaybudgetResponse
}

type TaobaoFeedflowItemCampaignDaybudgetResponse struct {
    XMLName xml.Name `xml:"feedflow_item_campaign_daybudget_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *TaobaoFeedflowItemCampaignDaybudgetResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
