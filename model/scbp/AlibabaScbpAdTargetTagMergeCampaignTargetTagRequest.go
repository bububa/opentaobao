package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
标签增删改 APIRequest
alibaba.scbp.ad.target.tag.merge.campaign.target.tag

标签增删改
*/
type AlibabaScbpAdTargetTagMergeCampaignTargetTagRequest struct {
    model.Params

    // 计划id
    campaignId   int64 

    // 标签数据，json格式。 最外层key：人群标签crowd/地域标签region、priceMode，第二层key: 增add、删del、改mod，第三层key：optionValue、bidRate、tagId  eg: 删除：{"crowd":{"del":[{"tagId":3595769030}]}}   修改：{"crowd":{"mod":[{"optionValue":"high_potential_order_user","bidRate":"151"}]}} 增加：{"crowd":{"add":[{"optionValue":"user_area_CA","bidRate":"133"}]}}
    data   string 

    // 用户信息
    topContext   *TopContextDto 

}

func NewAlibabaScbpAdTargetTagMergeCampaignTargetTagRequest() *AlibabaScbpAdTargetTagMergeCampaignTargetTagRequest{
    return &AlibabaScbpAdTargetTagMergeCampaignTargetTagRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdTargetTagMergeCampaignTargetTagRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.target.tag.merge.campaign.target.tag"
}

func (r AlibabaScbpAdTargetTagMergeCampaignTargetTagRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdTargetTagMergeCampaignTargetTagRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r AlibabaScbpAdTargetTagMergeCampaignTargetTagRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *AlibabaScbpAdTargetTagMergeCampaignTargetTagRequest) SetData(data string) error {
    r.data = data
    r.Set("data", data)
    return nil
}

func (r AlibabaScbpAdTargetTagMergeCampaignTargetTagRequest) GetData() string {
    return r.data
}

func (r *AlibabaScbpAdTargetTagMergeCampaignTargetTagRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

func (r AlibabaScbpAdTargetTagMergeCampaignTargetTagRequest) GetTopContext() *TopContextDto {
    return r.topContext
}

