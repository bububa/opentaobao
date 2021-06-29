package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建计划 API请求
alibaba.scbp.ad.campaign.create

创建计划
*/
type AlibabaScbpAdCampaignCreateRequest struct {
    model.Params
    // 返回数据
    campaignOperation   *CampaignOperationDto
    // 用户信息
    topContext   *TopContextDto
}

// 初始化AlibabaScbpAdCampaignCreateRequest对象
func NewAlibabaScbpAdCampaignCreateRequest() *AlibabaScbpAdCampaignCreateRequest{
    return &AlibabaScbpAdCampaignCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdCampaignCreateRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.campaign.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdCampaignCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignOperation Setter
// 返回数据
func (r *AlibabaScbpAdCampaignCreateRequest) SetCampaignOperation(campaignOperation *CampaignOperationDto) error {
    r.campaignOperation = campaignOperation
    r.Set("campaign_operation", campaignOperation)
    return nil
}

// CampaignOperation Getter
func (r AlibabaScbpAdCampaignCreateRequest) GetCampaignOperation() *CampaignOperationDto {
    return r.campaignOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdCampaignCreateRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdCampaignCreateRequest) GetTopContext() *TopContextDto {
    return r.topContext
}
