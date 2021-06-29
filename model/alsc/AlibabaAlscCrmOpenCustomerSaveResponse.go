package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
保存和更新顾客 APIResponse
alibaba.alsc.crm.open.customer.save

用来保存顾客，如果已经存在的话，则更新顾客
*/
type AlibabaAlscCrmOpenCustomerSaveAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmOpenCustomerSaveResponse
}

type AlibabaAlscCrmOpenCustomerSaveResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_open_customer_save_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
