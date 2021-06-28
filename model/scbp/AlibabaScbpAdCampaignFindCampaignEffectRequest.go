package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询计划效果数据 APIRequest
alibaba.scbp.ad.campaign.find.campaign.effect

批量查询计划效果数据
*/
type AlibabaScbpAdCampaignFindCampaignEffectRequest struct {
    model.Params

    // 计划id集合
    campaignIdList   []int64 

    // 开始时间
    beginDate   string 

    // 结束时间
    endDate   string 

    // 用户信息
    topContext   *TopContextDto 

}

func NewAlibabaScbpAdCampaignFindCampaignEffectRequest() *AlibabaScbpAdCampaignFindCampaignEffectRequest{
    return &AlibabaScbpAdCampaignFindCampaignEffectRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdCampaignFindCampaignEffectRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.campaign.find.campaign.effect"
}

func (r AlibabaScbpAdCampaignFindCampaignEffectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdCampaignFindCampaignEffectRequest) SetCampaignIdList(campaignIdList []int64) error {
    r.campaignIdList = campaignIdList
    r.Set("campaign_id_list", campaignIdList)
    return nil
}

func (r AlibabaScbpAdCampaignFindCampaignEffectRequest) GetCampaignIdList() []int64 {
    return r.campaignIdList
}

func (r *AlibabaScbpAdCampaignFindCampaignEffectRequest) SetBeginDate(beginDate string) error {
    r.beginDate = beginDate
    r.Set("begin_date", beginDate)
    return nil
}

func (r AlibabaScbpAdCampaignFindCampaignEffectRequest) GetBeginDate() string {
    return r.beginDate
}

func (r *AlibabaScbpAdCampaignFindCampaignEffectRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r AlibabaScbpAdCampaignFindCampaignEffectRequest) GetEndDate() string {
    return r.endDate
}

func (r *AlibabaScbpAdCampaignFindCampaignEffectRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

func (r AlibabaScbpAdCampaignFindCampaignEffectRequest) GetTopContext() *TopContextDto {
    return r.topContext
}

