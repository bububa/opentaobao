package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询屏蔽品 APIRequest
alibaba.scbp.ad.group.find.forbidden.product

查询屏蔽品
*/
type AlibabaScbpAdGroupFindForbiddenProductRequest struct {
    model.Params

    // 计划id
    campaignId   int64 

    // 用户信息
    topContext   *TopContextDto 

}

func NewAlibabaScbpAdGroupFindForbiddenProductRequest() *AlibabaScbpAdGroupFindForbiddenProductRequest{
    return &AlibabaScbpAdGroupFindForbiddenProductRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdGroupFindForbiddenProductRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.group.find.forbidden.product"
}

func (r AlibabaScbpAdGroupFindForbiddenProductRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdGroupFindForbiddenProductRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r AlibabaScbpAdGroupFindForbiddenProductRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *AlibabaScbpAdGroupFindForbiddenProductRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

func (r AlibabaScbpAdGroupFindForbiddenProductRequest) GetTopContext() *TopContextDto {
    return r.topContext
}

