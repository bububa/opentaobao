package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询计划列表 APIRequest
taobao.feedflow.item.campaign.page

批量查询计划列表
*/
type TaobaoFeedflowItemCampaignPageRequest struct {
    model.Params

    // 入参
    campaignQuery   *CampaignQueryDto 

}

func NewTaobaoFeedflowItemCampaignPageRequest() *TaobaoFeedflowItemCampaignPageRequest{
    return &TaobaoFeedflowItemCampaignPageRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemCampaignPageRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.campaign.page"
}

func (r TaobaoFeedflowItemCampaignPageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemCampaignPageRequest) SetCampaignQuery(campaignQuery *CampaignQueryDto) error {
    r.campaignQuery = campaignQuery
    r.Set("campaign_query", campaignQuery)
    return nil
}

func (r TaobaoFeedflowItemCampaignPageRequest) GetCampaignQuery() *CampaignQueryDto {
    return r.campaignQuery
}

