package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询计划 API请求
alibaba.scbp.ad.campaign.find.campaign.page

分页查询计划
*/
type AlibabaScbpAdCampaignFindCampaignPageRequest struct {
    model.Params
    // 请求实体类
    _campaignQuery   *CampaignQueryDto
    // 用户信息
    _topContext   *TopContextDto
}

// 初始化AlibabaScbpAdCampaignFindCampaignPageRequest对象
func NewAlibabaScbpAdCampaignFindCampaignPageRequest() *AlibabaScbpAdCampaignFindCampaignPageRequest{
    return &AlibabaScbpAdCampaignFindCampaignPageRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdCampaignFindCampaignPageRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.campaign.find.campaign.page"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdCampaignFindCampaignPageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignQuery Setter
// 请求实体类
func (r *AlibabaScbpAdCampaignFindCampaignPageRequest) SetCampaignQuery(_campaignQuery *CampaignQueryDto) error {
    r._campaignQuery = _campaignQuery
    r.Set("campaign_query", _campaignQuery)
    return nil
}

// CampaignQuery Getter
func (r AlibabaScbpAdCampaignFindCampaignPageRequest) GetCampaignQuery() *CampaignQueryDto {
    return r._campaignQuery
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdCampaignFindCampaignPageRequest) SetTopContext(_topContext *TopContextDto) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdCampaignFindCampaignPageRequest) GetTopContext() *TopContextDto {
    return r._topContext
}
