package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广 按照id操作推广计划的定向溢价标签，包括新增，删除和更新 APIRequest
alibaba.scbp.target.ad.plan.update.tags

定向推广 按照id操作推广计划的定向溢价标签，包括新增，删除和更新
*/
type AlibabaScbpTargetAdPlanUpdateTagsRequest struct {
    model.Params

    // 系统生成
    paramTopP4pModifyQuickCampaignTagDTO   *TopP4pModifyQuickCampaignTagDto 

}

func NewAlibabaScbpTargetAdPlanUpdateTagsRequest() *AlibabaScbpTargetAdPlanUpdateTagsRequest{
    return &AlibabaScbpTargetAdPlanUpdateTagsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpTargetAdPlanUpdateTagsRequest) GetApiMethodName() string {
    return "alibaba.scbp.target.ad.plan.update.tags"
}

func (r AlibabaScbpTargetAdPlanUpdateTagsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpTargetAdPlanUpdateTagsRequest) SetParamTopP4pModifyQuickCampaignTagDTO(paramTopP4pModifyQuickCampaignTagDTO *TopP4pModifyQuickCampaignTagDto) error {
    r.paramTopP4pModifyQuickCampaignTagDTO = paramTopP4pModifyQuickCampaignTagDTO
    r.Set("param_top_p4p_modify_quick_campaign_tag_d_t_o", paramTopP4pModifyQuickCampaignTagDTO)
    return nil
}

func (r AlibabaScbpTargetAdPlanUpdateTagsRequest) GetParamTopP4pModifyQuickCampaignTagDTO() *TopP4pModifyQuickCampaignTagDto {
    return r.paramTopP4pModifyQuickCampaignTagDTO
}

