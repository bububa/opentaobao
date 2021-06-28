package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取促销规则列表 APIResponse
alibaba.alsc.crm.promotion.list

获取品牌的促销规则列表
*/
type AlibabaAlscCrmPromotionListAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmPromotionListResponse `json:"alibaba_alsc_crm_promotion_list_response,omitempty"` 
    AlibabaAlscCrmPromotionListResponse
}

/* model for simplify = false
type AlibabaAlscCrmPromotionListResponse struct {

    // 分页返回模型
    
    Result  *struct {
        CommonPageResult  *CommonPageResult `json:"common_page_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmPromotionListResponse struct {

    // 分页返回模型
    Result   *CommonPageResult `json:"result,omitempty"`

}
