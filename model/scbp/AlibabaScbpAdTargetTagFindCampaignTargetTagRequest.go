package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询标签数据 API请求
alibaba.scbp.ad.target.tag.find.campaign.target.tag

查询标签数据
*/
type AlibabaScbpAdTargetTagFindCampaignTargetTagRequest struct {
    model.Params
    // 计划id
    _campaignId   int64
    // 请求参数
    _targetTagOperation   *TargetTagOperationDTO
    // 用户信息
    _topContext   *TopContextDTO
}

// 初始化AlibabaScbpAdTargetTagFindCampaignTargetTagRequest对象
func NewAlibabaScbpAdTargetTagFindCampaignTargetTagRequest() *AlibabaScbpAdTargetTagFindCampaignTargetTagRequest{
    return &AlibabaScbpAdTargetTagFindCampaignTargetTagRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdTargetTagFindCampaignTargetTagRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.target.tag.find.campaign.target.tag"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdTargetTagFindCampaignTargetTagRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdTargetTagFindCampaignTargetTagRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdTargetTagFindCampaignTargetTagRequest) GetCampaignId() int64 {
    return r._campaignId
}
// TargetTagOperation Setter
// 请求参数
func (r *AlibabaScbpAdTargetTagFindCampaignTargetTagRequest) SetTargetTagOperation(_targetTagOperation *TargetTagOperationDTO) error {
    r._targetTagOperation = _targetTagOperation
    r.Set("target_tag_operation", _targetTagOperation)
    return nil
}

// TargetTagOperation Getter
func (r AlibabaScbpAdTargetTagFindCampaignTargetTagRequest) GetTargetTagOperation() *TargetTagOperationDTO {
    return r._targetTagOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdTargetTagFindCampaignTargetTagRequest) SetTopContext(_topContext *TopContextDTO) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdTargetTagFindCampaignTargetTagRequest) GetTopContext() *TopContextDTO {
    return r._topContext
}
