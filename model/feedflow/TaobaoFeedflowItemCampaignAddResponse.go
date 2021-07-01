package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流增加推广计划 API返回值 
taobao.feedflow.item.campaign.add

信息流增加推广计划
*/
type TaobaoFeedflowItemCampaignAddAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemCampaignAddResponse
}

// 信息流增加推广计划 成功返回结果
type TaobaoFeedflowItemCampaignAddResponse struct {
    XMLName xml.Name `xml:"feedflow_item_campaign_add_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 异步获取历史数据接口返回结果
    Result   *ResultDTO `json:"result,omitempty" xml:"result,omitempty"`
}
