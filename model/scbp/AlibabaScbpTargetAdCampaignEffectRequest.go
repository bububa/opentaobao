package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-获取计划维度推广效果 APIRequest
alibaba.scbp.target.ad.campaign.effect

定向推广-获取计划维度推广效果
*/
type AlibabaScbpTargetAdCampaignEffectRequest struct {
    model.Params

    // 统计区间 只能为1 7 30
    interval   int64 

    // 结束时间 当inteval=7或30的时候 不需要填写，当inteval=1时需要填写（开始结束时间区间不允许大于180天）
    endDate   string 

    // 开始时间 当inteval=7或30的时候 不需要填写，当inteval=1时需要填写（开始结束时间区间不允许大于180天）
    beginDate   string 

    // 当填写时，展示指定id的数据，不填写，则展示全部计划总数据
    campaignId   int64 

}

func NewAlibabaScbpTargetAdCampaignEffectRequest() *AlibabaScbpTargetAdCampaignEffectRequest{
    return &AlibabaScbpTargetAdCampaignEffectRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpTargetAdCampaignEffectRequest) GetApiMethodName() string {
    return "alibaba.scbp.target.ad.campaign.effect"
}

func (r AlibabaScbpTargetAdCampaignEffectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpTargetAdCampaignEffectRequest) SetInterval(interval int64) error {
    r.interval = interval
    r.Set("interval", interval)
    return nil
}

func (r AlibabaScbpTargetAdCampaignEffectRequest) GetInterval() int64 {
    return r.interval
}

func (r *AlibabaScbpTargetAdCampaignEffectRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r AlibabaScbpTargetAdCampaignEffectRequest) GetEndDate() string {
    return r.endDate
}

func (r *AlibabaScbpTargetAdCampaignEffectRequest) SetBeginDate(beginDate string) error {
    r.beginDate = beginDate
    r.Set("begin_date", beginDate)
    return nil
}

func (r AlibabaScbpTargetAdCampaignEffectRequest) GetBeginDate() string {
    return r.beginDate
}

func (r *AlibabaScbpTargetAdCampaignEffectRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r AlibabaScbpTargetAdCampaignEffectRequest) GetCampaignId() int64 {
    return r.campaignId
}

