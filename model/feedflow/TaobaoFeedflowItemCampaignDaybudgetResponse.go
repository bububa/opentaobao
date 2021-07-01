package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取当日投放日预算总额 API返回值 
taobao.feedflow.item.campaign.daybudget

获取当日投放日预算总额
*/
type TaobaoFeedflowItemCampaignDaybudgetAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemCampaignDaybudgetResponse
}

// 获取当日投放日预算总额 成功返回结果
type TaobaoFeedflowItemCampaignDaybudgetResponse struct {
    XMLName xml.Name `xml:"feedflow_item_campaign_daybudget_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Result   *ResultDTO `json:"result,omitempty" xml:"result,omitempty"`
}
