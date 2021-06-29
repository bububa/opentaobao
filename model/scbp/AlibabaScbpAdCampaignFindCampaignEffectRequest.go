package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询计划效果数据 API请求
alibaba.scbp.ad.campaign.find.campaign.effect

批量查询计划效果数据
*/
type AlibabaScbpAdCampaignFindCampaignEffectRequest struct {
    model.Params
    // 计划id集合
    _campaignIdList   []int64
    // 开始时间
    _beginDate   string
    // 结束时间
    _endDate   string
    // 用户信息
    _topContext   *TopContextDto
}

// 初始化AlibabaScbpAdCampaignFindCampaignEffectRequest对象
func NewAlibabaScbpAdCampaignFindCampaignEffectRequest() *AlibabaScbpAdCampaignFindCampaignEffectRequest{
    return &AlibabaScbpAdCampaignFindCampaignEffectRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdCampaignFindCampaignEffectRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.campaign.find.campaign.effect"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdCampaignFindCampaignEffectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignIdList Setter
// 计划id集合
func (r *AlibabaScbpAdCampaignFindCampaignEffectRequest) SetCampaignIdList(_campaignIdList []int64) error {
    r._campaignIdList = _campaignIdList
    r.Set("campaign_id_list", _campaignIdList)
    return nil
}

// CampaignIdList Getter
func (r AlibabaScbpAdCampaignFindCampaignEffectRequest) GetCampaignIdList() []int64 {
    return r._campaignIdList
}
// BeginDate Setter
// 开始时间
func (r *AlibabaScbpAdCampaignFindCampaignEffectRequest) SetBeginDate(_beginDate string) error {
    r._beginDate = _beginDate
    r.Set("begin_date", _beginDate)
    return nil
}

// BeginDate Getter
func (r AlibabaScbpAdCampaignFindCampaignEffectRequest) GetBeginDate() string {
    return r._beginDate
}
// EndDate Setter
// 结束时间
func (r *AlibabaScbpAdCampaignFindCampaignEffectRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r AlibabaScbpAdCampaignFindCampaignEffectRequest) GetEndDate() string {
    return r._endDate
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdCampaignFindCampaignEffectRequest) SetTopContext(_topContext *TopContextDto) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdCampaignFindCampaignEffectRequest) GetTopContext() *TopContextDto {
    return r._topContext
}
