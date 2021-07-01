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
type AlibabaScbpAdCampaignCreateAPIRequest struct {
    model.Params
    // 返回数据
    _campaignOperation   *CampaignOperationDto
    // 用户信息
    _topContext   *TopContextDto
}

// 初始化AlibabaScbpAdCampaignCreateAPIRequest对象
func NewAlibabaScbpAdCampaignCreateRequest() *AlibabaScbpAdCampaignCreateAPIRequest{
    return &AlibabaScbpAdCampaignCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdCampaignCreateAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.campaign.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdCampaignCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignOperation Setter
// 返回数据
func (r *AlibabaScbpAdCampaignCreateAPIRequest) SetCampaignOperation(_campaignOperation *CampaignOperationDto) error {
    r._campaignOperation = _campaignOperation
    r.Set("campaign_operation", _campaignOperation)
    return nil
}

// CampaignOperation Getter
func (r AlibabaScbpAdCampaignCreateAPIRequest) GetCampaignOperation() *CampaignOperationDto {
    return r._campaignOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdCampaignCreateAPIRequest) SetTopContext(_topContext *TopContextDto) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdCampaignCreateAPIRequest) GetTopContext() *TopContextDto {
    return r._topContext
}
