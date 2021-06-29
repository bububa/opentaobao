package alihealthpw

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
回调变更患者姓名 APIResponse
alibaba.alihealth.pw.applynode.updatename

回调变更患者姓名
*/
type AlibabaAlihealthPwApplynodeUpdatenameAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthPwApplynodeUpdatenameResponse
}

type AlibabaAlihealthPwApplynodeUpdatenameResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_pw_applynode_updatename_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // pap项目状态码
    
    PapCode   string `json:"pap_code,omitempty" xml:"pap_code,omitempty"`

    
    // pap项目状态描述
    
    PapMessage   string `json:"pap_message,omitempty" xml:"pap_message,omitempty"`

    
}
