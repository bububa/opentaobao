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
    campaignQuery   *CampaignQueryDto
    // 用户信息
    topContext   *TopContextDto
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
func (r *AlibabaScbpAdCampaignFindRealCostRequest) SetCampaignQuery(campaignQuery *CampaignQueryDto) error {
    r.campaignQuery = campaignQuery
    r.Set("campaign_query", campaignQuery)
    return nil
}

// CampaignQuery Getter
func (r AlibabaScbpAdCampaignFindRealCostRequest) GetCampaignQuery() *CampaignQueryDto {
    return r.campaignQuery
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdCampaignFindRealCostRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdCampaignFindRealCostRequest) GetTopContext() *TopContextDto {
    return r.topContext
}
