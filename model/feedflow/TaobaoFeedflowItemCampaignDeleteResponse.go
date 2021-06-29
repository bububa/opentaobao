package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除计划 API返回值 
taobao.feedflow.item.campaign.delete

删除计划
*/
type TaobaoFeedflowItemCampaignDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemCampaignDeleteResponse
}

// 删除计划 成功返回结果
type TaobaoFeedflowItemCampaignDeleteResponse struct {
    XMLName xml.Name `xml:"feedflow_item_campaign_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Result   *TaobaoFeedflowItemCampaignDeleteResultDTO `json:"result,omitempty" xml:"result,omitempty"`
}
