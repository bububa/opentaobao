package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-计划开启/暂停/删除 APIRequest
alibaba.scbp.target.ad.plan.operation

定向推广-计划开启/暂停/删除
*/
type AlibabaScbpTargetAdPlanOperationRequest struct {
    model.Params

    // TopP4pModifyQuickCampaignDTO
    topP4pModifyQuickCampaignDTO   *TopP4pModifyQuickCampaignDto 

}

func NewAlibabaScbpTargetAdPlanOperationRequest() *AlibabaScbpTargetAdPlanOperationRequest{
    return &AlibabaScbpTargetAdPlanOperationRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpTargetAdPlanOperationRequest) GetApiMethodName() string {
    return "alibaba.scbp.target.ad.plan.operation"
}

func (r AlibabaScbpTargetAdPlanOperationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpTargetAdPlanOperationRequest) SetTopP4pModifyQuickCampaignDTO(topP4pModifyQuickCampaignDTO *TopP4pModifyQuickCampaignDto) error {
    r.topP4pModifyQuickCampaignDTO = topP4pModifyQuickCampaignDTO
    r.Set("top_p4p_modify_quick_campaign_d_t_o", topP4pModifyQuickCampaignDTO)
    return nil
}

func (r AlibabaScbpTargetAdPlanOperationRequest) GetTopP4pModifyQuickCampaignDTO() *TopP4pModifyQuickCampaignDto {
    return r.topP4pModifyQuickCampaignDTO
}

