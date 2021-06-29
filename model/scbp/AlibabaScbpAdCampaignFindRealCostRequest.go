package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询计划消耗数据 API请求
alibaba.scbp.ad.campaign.find.real.cost

批量查询计划消耗数据
*/
type AlibabaScbpAdCampaignFindRealCostRequest struct {
    model.Params
    // 系统自动生成
    _campaignQuery   *CampaignQueryDto
    // 用户信息
    _topContext   *TopContextDto
}

// 初始化AlibabaScbpAdCampaignFindRealCostRequest对象
func NewAlibabaScbpAdCampaignFindRealCostRequest() *AlibabaScbpAdCampaignFindRealCostRequest{
    return &AlibabaScbpAdCampaignFindRealCostRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdCampaignFindRealCostRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.campaign.find.real.cost"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdCampaignFindRealCostRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignQuery Setter
// 系统自动生成
func (r *AlibabaScbpAdCampaignFindRealCostRequest) SetCampaignQuery(_campaignQuery *CampaignQueryDto) error {
    r._campaignQuery = _campaignQuery
    r.Set("campaign_query", _campaignQuery)
    return nil
}

// CampaignQuery Getter
func (r AlibabaScbpAdCampaignFindRealCostRequest) GetCampaignQuery() *CampaignQueryDto {
    return r._campaignQuery
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdCampaignFindRealCostRequest) SetTopContext(_topContext *TopContextDto) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdCampaignFindRealCostRequest) GetTopContext() *TopContextDto {
    return r._topContext
}
