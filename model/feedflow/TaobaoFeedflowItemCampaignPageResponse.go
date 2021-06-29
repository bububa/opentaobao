package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询计划列表 APIResponse
taobao.feedflow.item.campaign.page

批量查询计划列表
*/
type TaobaoFeedflowItemCampaignPageAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemCampaignPageResponse
}

type TaobaoFeedflowItemCampaignPageResponse struct {
    XMLName xml.Name `xml:"feedflow_item_campaign_page_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *TaobaoFeedflowItemCampaignPageResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
