package aliexpress

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV更新INS私信发送的结果 APIResponse
aliexpress.social.ins.directresult.update

ISV更新INS私信发送的结果
*/
type AliexpressSocialInsDirectresultUpdateAPIResponse struct {
    model.CommonResponse
    AliexpressSocialInsDirectresultUpdateResponse
}

type AliexpressSocialInsDirectresultUpdateResponse struct {
    XMLName xml.Name `xml:"aliexpress_social_ins_directresult_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 更新是否成功
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
    // 此次调用是否成功
    
    Successs   bool `json:"successs,omitempty" xml:"successs,omitempty"`

    
    // 错误码
    
    ErrorCodee   string `json:"error_codee,omitempty" xml:"error_codee,omitempty"`

    
    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`

    
}
