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
type AlibabaScbpAdCampaignUpdateAPIRequest struct {
    model.Params
    // 修改数据
    _campaignOperation   *CampaignOperationDto
    // 用户信息
    _topContext   *TopContextDto
}

// 初始化AlibabaScbpAdCampaignUpdateAPIRequest对象
func NewAlibabaScbpAdCampaignUpdateRequest() *AlibabaScbpAdCampaignUpdateAPIRequest{
    return &AlibabaScbpAdCampaignUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdCampaignUpdateAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.campaign.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdCampaignUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignOperation Setter
// 修改数据
func (r *AlibabaScbpAdCampaignUpdateAPIRequest) SetCampaignOperation(_campaignOperation *CampaignOperationDto) error {
    r._campaignOperation = _campaignOperation
    r.Set("campaign_operation", _campaignOperation)
    return nil
}

// CampaignOperation Getter
func (r AlibabaScbpAdCampaignUpdateAPIRequest) GetCampaignOperation() *CampaignOperationDto {
    return r._campaignOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdCampaignUpdateAPIRequest) SetTopContext(_topContext *TopContextDto) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdCampaignUpdateAPIRequest) GetTopContext() *TopContextDto {
    return r._topContext
}
