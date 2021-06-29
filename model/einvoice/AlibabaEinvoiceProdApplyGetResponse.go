package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询发票申请 APIResponse
alibaba.einvoice.prod.apply.get

查询申请的详细信息，包含申请所关联的发票摘要信息+板式文件+预览图；
场景使用：1、业务前台收到申请状态变更消息后，调用此接口查询申请详情；2、主动补偿查询：当指定了自动开票，且发票申请长时间未收到状态变更通知时，可能存在丢消息的情况，此时可主动查询该申请，然后更新本地工单状态。
*/
type AlibabaEinvoiceProdApplyGetAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceProdApplyGetResponse
}

type AlibabaEinvoiceProdApplyGetResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_prod_apply_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
