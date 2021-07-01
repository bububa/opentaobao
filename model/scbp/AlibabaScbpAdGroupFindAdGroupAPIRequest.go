package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询推广组 API请求
alibaba.scbp.ad.group.find.ad.group

查询推广组
*/
type AlibabaScbpAdGroupFindAdGroupAPIRequest struct {
    model.Params
    // 计划id
    _campaignId   int64
    // 入参
    _adGroupQuery   *AdGroupQueryDto
    // 用户信息
    _topContext   *TopContextDto
}

// 初始化AlibabaScbpAdGroupFindAdGroupAPIRequest对象
func NewAlibabaScbpAdGroupFindAdGroupRequest() *AlibabaScbpAdGroupFindAdGroupAPIRequest{
    return &AlibabaScbpAdGroupFindAdGroupAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdGroupFindAdGroupAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.group.find.ad.group"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdGroupFindAdGroupAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdGroupFindAdGroupAPIRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdGroupFindAdGroupAPIRequest) GetCampaignId() int64 {
    return r._campaignId
}
// AdGroupQuery Setter
// 入参
func (r *AlibabaScbpAdGroupFindAdGroupAPIRequest) SetAdGroupQuery(_adGroupQuery *AdGroupQueryDto) error {
    r._adGroupQuery = _adGroupQuery
    r.Set("ad_group_query", _adGroupQuery)
    return nil
}

// AdGroupQuery Getter
func (r AlibabaScbpAdGroupFindAdGroupAPIRequest) GetAdGroupQuery() *AdGroupQueryDto {
    return r._adGroupQuery
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdGroupFindAdGroupAPIRequest) SetTopContext(_topContext *TopContextDto) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdGroupFindAdGroupAPIRequest) GetTopContext() *TopContextDto {
    return r._topContext
}
