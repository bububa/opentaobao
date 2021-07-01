package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询计划列表 API返回值 
taobao.feedflow.item.campaign.page

批量查询计划列表
*/
type TaobaoFeedflowItemCampaignPageAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemCampaignPageAPIResponseModel
}

// 批量查询计划列表 成功返回结果
type TaobaoFeedflowItemCampaignPageAPIResponseModel struct {
    XMLName xml.Name `xml:"feedflow_item_campaign_page_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Result   *TaobaoFeedflowItemCampaignPageResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
