package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询分享营销客户领券信息 APIResponse
alibaba.alsc.crm.marketing.share.customer.info

查询分享营销活动的客户领券信息
*/
type AlibabaAlscCrmMarketingShareCustomerInfoAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscCrmMarketingShareCustomerInfoResponse `json:"alibaba_alsc_crm_marketing_share_customer_info_response,omitempty"`
}

type AlibabaAlscCrmMarketingShareCustomerInfoResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
