package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取优惠券模版列表 APIResponse
alibaba.alsc.crm.voucher.template.list

获取优惠券模版列表
*/
type AlibabaAlscCrmVoucherTemplateListAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmVoucherTemplateListResponse `json:"alibaba_alsc_crm_voucher_template_list_response,omitempty"` 
    AlibabaAlscCrmVoucherTemplateListResponse
}

/* model for simplify = false
type AlibabaAlscCrmVoucherTemplateListResponse struct {

    // 分页返回模型
    
    Result  *struct {
        CommonPageResult  *CommonPageResult `json:"common_page_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmVoucherTemplateListResponse struct {

    // 分页返回模型
    Result   *CommonPageResult `json:"result,omitempty"`

}
