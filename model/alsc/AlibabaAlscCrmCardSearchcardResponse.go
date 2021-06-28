package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
搜索卡实例列表(支持号段查询) APIResponse
alibaba.alsc.crm.card.searchcard

搜索卡实例列表(支持号段查询)
*/
type AlibabaAlscCrmCardSearchcardAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmCardSearchcardResponse `json:"alibaba_alsc_crm_card_searchcard_response,omitempty"` 
    AlibabaAlscCrmCardSearchcardResponse
}

/* model for simplify = false
type AlibabaAlscCrmCardSearchcardResponse struct {

    // 分页返回模型
    
    Result  *struct {
        CommonPageResult  *CommonPageResult `json:"common_page_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmCardSearchcardResponse struct {

    // 分页返回模型
    Result   *CommonPageResult `json:"result,omitempty"`

}
