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
type AlibabaScbpAdGroupFindAdGroupRequest struct {
    model.Params
    // 计划id
    _campaignId   int64
    // 入参
    _adGroupQuery   *AdGroupQueryDTO
    // 用户信息
    _topContext   *TopContextDTO
}

// 初始化AlibabaScbpAdGroupFindAdGroupRequest对象
func NewAlibabaScbpAdGroupFindAdGroupRequest() *AlibabaScbpAdGroupFindAdGroupRequest{
    return &AlibabaScbpAdGroupFindAdGroupRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdGroupFindAdGroupRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.group.find.ad.group"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdGroupFindAdGroupRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdGroupFindAdGroupRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdGroupFindAdGroupRequest) GetCampaignId() int64 {
    return r._campaignId
}
// AdGroupQuery Setter
// 入参
func (r *AlibabaScbpAdGroupFindAdGroupRequest) SetAdGroupQuery(_adGroupQuery *AdGroupQueryDTO) error {
    r._adGroupQuery = _adGroupQuery
    r.Set("ad_group_query", _adGroupQuery)
    return nil
}

// AdGroupQuery Getter
func (r AlibabaScbpAdGroupFindAdGroupRequest) GetAdGroupQuery() *AdGroupQueryDTO {
    return r._adGroupQuery
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdGroupFindAdGroupRequest) SetTopContext(_topContext *TopContextDTO) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdGroupFindAdGroupRequest) GetTopContext() *TopContextDTO {
    return r._topContext
}
