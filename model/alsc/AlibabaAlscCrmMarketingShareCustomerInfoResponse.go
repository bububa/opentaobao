package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询分享营销客户领券信息 APIResponse
alibaba.alsc.crm.marketing.share.customer.info

查询分享营销活动的客户领券信息
*/
type AlibabaAlscCrmMarketingShareCustomerInfoAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmMarketingShareCustomerInfoResponse
}

type AlibabaAlscCrmMarketingShareCustomerInfoResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_marketing_share_customer_info_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
