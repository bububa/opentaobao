package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取优惠券模版列表 APIResponse
alibaba.alsc.crm.voucher.template.list

获取优惠券模版列表
*/
type AlibabaAlscCrmVoucherTemplateListAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alsc_crm_voucher_template_list_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 分页返回模型
    
    Result   *CommonPageResult `json:"result,omitempty" xml:"