package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发票规则设置 APIResponse
alitrip.btrip.invoice.setting.rule

发票规则设置
*/
type AlitripBtripInvoiceSettingRuleAPIResponse struct {
    model.CommonResponse
    AlitripBtripInvoiceSettingRuleResponse
}

type AlitripBtripInvoiceSettingRuleResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_invoice_setting_rule_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回值
    
    Result   *OpenInvoiceRuleRs `json:"result,omitempty" xml:"result,omitempty"`

    
    // 结果码
    
    ResultCode   int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 结果描述
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
}
