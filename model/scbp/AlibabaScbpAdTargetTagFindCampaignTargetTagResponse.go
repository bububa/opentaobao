package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询标签数据 APIResponse
alibaba.scbp.ad.target.tag.find.campaign.target.tag

查询标签数据
*/
type AlibabaScbpAdTargetTagFindCampaignTargetTagAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpAdTargetTagFindCampaignTargetTagResponse `json:"alibaba_scbp_ad_target_tag_find_campaign_target_tag_response,omitempty"`
}

type AlibabaScbpAdTargetTagFindCampaignTargetTagResponse struct {

    // 返回实体
    ResultList   []AdsTargetingTagDto `json:"result_list,omitempty"`

}
