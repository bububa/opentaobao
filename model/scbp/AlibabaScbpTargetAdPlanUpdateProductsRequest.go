package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广 按照id操作推广计划的产品，包括新增，删除和更新 APIRequest
alibaba.scbp.target.ad.plan.update.products

定向推广 按照id操作推广计划的产品，包括新增，删除和更新
*/
type AlibabaScbpTargetAdPlanUpdateProductsRequest struct {
    model.Params

    // 系统生成
    paramTopP4pModifyQuickCampaignProductDTO   *TopP4pModifyQuickCampaignProductDto 

}

func NewAlibabaScbpTargetAdPlanUpdateProductsRequest() *AlibabaScbpTargetAdPlanUpdateProductsRequest{
    return &AlibabaScbpTargetAdPlanUpdateProductsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpTargetAdPlanUpdateProductsRequest) GetApiMethodName() string {
    return "alibaba.scbp.target.ad.plan.update.products"
}

func (r AlibabaScbpTargetAdPlanUpdateProductsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpTargetAdPlanUpdateProductsRequest) SetParamTopP4pModifyQuickCampaignProductDTO(paramTopP4pModifyQuickCampaignProductDTO *TopP4pModifyQuickCampaignProductDto) error {
    r.paramTopP4pModifyQuickCampaignProductDTO = paramTopP4pModifyQuickCampaignProductDTO
    r.Set("param_top_p4p_modify_quick_campaign_product_d_t_o", paramTopP4pModifyQuickCampaignProductDTO)
    return nil
}

func (r AlibabaScbpTargetAdPlanUpdateProductsRequest) GetParamTopP4pModifyQuickCampaignProductDTO() *TopP4pModifyQuickCampaignProductDto {
    return r.paramTopP4pModifyQuickCampaignProductDTO
}

