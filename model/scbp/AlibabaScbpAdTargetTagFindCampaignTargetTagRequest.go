package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询标签数据 APIRequest
alibaba.scbp.ad.target.tag.find.campaign.target.tag

查询标签数据
*/
type AlibabaScbpAdTargetTagFindCampaignTargetTagRequest struct {
    model.Params

    // 计划id
    campaignId   int64 

    // 请求参数
    targetTagOperation   *TargetTagOperationDto 

    // 用户信息
    topContext   *TopContextDto 

}

func NewAlibabaScbpAdTargetTagFindCampaignTargetTagRequest() *AlibabaScbpAdTargetTagFindCampaignTargetTagRequest{
    return &AlibabaScbpAdTargetTagFindCampaignTargetTagRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdTargetTagFindCampaignTargetTagRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.target.tag.find.campaign.target.tag"
}

func (r AlibabaScbpAdTargetTagFindCampaignTargetTagRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdTargetTagFindCampaignTargetTagRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r AlibabaScbpAdTargetTagFindCampaignTargetTagRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *AlibabaScbpAdTargetTagFindCampaignTargetTagRequest) SetTargetTagOperation(targetTagOperation *TargetTagOperationDto) error {
    r.targetTagOperation = targetTagOperation
    r.Set("target_tag_operation", targetTagOperation)
    return nil
}

func (r AlibabaScbpAdTargetTagFindCampaignTargetTagRequest) GetTargetTagOperation() *TargetTagOperationDto {
    return r.targetTagOperation
}

func (r *AlibabaScbpAdTargetTagFindCampaignTargetTagRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

func (r AlibabaScbpAdTargetTagFindCampaignTargetTagRequest) GetTopContext() *TopContextDto {
    return r.topContext
}

