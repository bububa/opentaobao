package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
直调审批 APIResponse
alibaba.alihealth.drug.kyt.destbill.check

为药企提供直调单据审批操作
*/
type AlibabaAlihealthDrugKytDestbillCheckAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytDestbillCheckResponse
}

type AlibabaAlihealthDrugKytDestbillCheckResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_destbill_check_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功
    
    ResponseSuccess   bool `json:"response_success,omitempty" xml:"response_success,omitempty"`

    
    // 执行结果
    
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`

    
    // 返回结果描述
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
    // 返回结果标识
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
}
