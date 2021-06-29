package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改计划 API请求
alibaba.scbp.ad.campaign.update

修改计划
*/
type AlibabaScbpAdCampaignUpdateRequest struct {
    model.Params
    // 修改数据
    campaignOperation   *CampaignOperationDto
    // 用户信息
    topContext   *TopContextDto
}

// 初始化AlibabaScbpAdCampaignUpdateRequest对象
func NewAlibabaScbpAdCampaignUpdateRequest() *AlibabaScbpAdCampaignUpdateRequest{
    return &AlibabaScbpAdCampaignUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdCampaignUpdateRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.campaign.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdCampaignUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignOperation Setter
// 修改数据
func (r *AlibabaScbpAdCampaignUpdateRequest) SetCampaignOperation(campaignOperation *CampaignOperationDto) error {
    r.campaignOperation = campaignOperation
    r.Set("campaign_operation", campaignOperation)
    return nil
}

// CampaignOperation Getter
func (r AlibabaScbpAdCampaignUpdateRequest) GetCampaignOperation() *CampaignOperationDto {
    return r.campaignOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdCampaignUpdateRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdCampaignUpdateRequest) GetTopContext() *TopContextDto {
    return r.topContext
}
