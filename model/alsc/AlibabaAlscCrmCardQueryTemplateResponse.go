package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询卡模板详情 APIResponse
alibaba.alsc.crm.card.query.template

查询卡模板详情
*/
type AlibabaAlscCrmCardQueryTemplateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmCardQueryTemplateResponse `json:"alibaba_alsc_crm_card_query_template_response,omitempty"` 
    AlibabaAlscCrmCardQueryTemplateResponse
}

/* model for simplify = false
type AlibabaAlscCrmCardQueryTemplateResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmCardQueryTemplateResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
