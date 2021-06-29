package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应链中台货品修改接口 APIResponse
alibaba.ascp.cnsku.update

供应链中台货品修改接口
*/
type AlibabaAscpCnskuUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaAscpCnskuUpdateResponse
}

type AlibabaAscpCnskuUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_ascp_cnsku_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 异常信息
    
    ErrorMessages   []string `json:"error_messages,omitempty" xml:"error_messages>string,omitempty"`
    
    
    // 货品id
    
    Data   string `json:"data,omitempty" xml:"data,omitempty"`

    
    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 是否系统异常
    
    IsSystemFailed   bool `json:"is_system_failed,omitempty" xml:"is_system_failed,omitempty"`

    
    // 异常信息Code
    
    SysErrorCode   string `json:"sys_error_code,omitempty" xml:"sys_error_code,omitempty"`

    
}
