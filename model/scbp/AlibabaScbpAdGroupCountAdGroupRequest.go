package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
统计adgroup数量 API请求
alibaba.scbp.ad.group.count.ad.group

统计adgroup数量
*/
type AlibabaScbpAdGroupCountAdGroupRequest struct {
    model.Params
    // 计划id
    _campaignId   int64
    // 查询条件
    _adGroupQuery   *AdGroupQueryDTO
    // 用户信息
    _topContext   *TopContextDTO
}

// 初始化AlibabaScbpAdGroupCountAdGroupRequest对象
func NewAlibabaScbpAdGroupCountAdGroupRequest() *AlibabaScbpAdGroupCountAdGroupRequest{
    return &AlibabaScbpAdGroupCountAdGroupRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdGroupCountAdGroupRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.group.count.ad.group"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdGroupCountAdGroupRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdGroupCountAdGroupRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdGroupCountAdGroupRequest) GetCampaignId() int64 {
    return r._campaignId
}
// AdGroupQuery Setter
// 查询条件
func (r *AlibabaScbpAdGroupCountAdGroupRequest) SetAdGroupQuery(_adGroupQuery *AdGroupQueryDTO) error {
    r._adGroupQuery = _adGroupQuery
    r.Set("ad_group_query", _adGroupQuery)
    return nil
}

// AdGroupQuery Getter
func (r AlibabaScbpAdGroupCountAdGroupRequest) GetAdGroupQuery() *AdGroupQueryDTO {
    return r._adGroupQuery
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdGroupCountAdGroupRequest) SetTopContext(_topContext *TopContextDTO) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdGroupCountAdGroupRequest) GetTopContext() *TopContextDTO {
    return r._topContext
}
