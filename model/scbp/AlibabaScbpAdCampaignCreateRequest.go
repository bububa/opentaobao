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
    _campaignOperation   *CampaignOperationDTO
    // 用户信息
    _topContext   *TopContextDTO
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
func (r *AlibabaScbpAdCampaignCreateRequest) SetCampaignOperation(_campaignOperation *CampaignOperationDTO) error {
    r._campaignOperation = _campaignOperation
    r.Set("campaign_operation", _campaignOperation)
    return nil
}

// CampaignOperation Getter
func (r AlibabaScbpAdCampaignCreateRequest) GetCampaignOperation() *CampaignOperationDTO {
    return r._campaignOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdCampaignCreateRequest) SetTopContext(_topContext *TopContextDTO) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdCampaignCreateRequest) GetTopContext() *TopContextDTO {
    return r._topContext
}
