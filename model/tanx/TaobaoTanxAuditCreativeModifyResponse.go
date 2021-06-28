package tanx

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创意修改接口 APIResponse
taobao.tanx.audit.creative.modify

创意修改接口
*/
type TaobaoTanxAuditCreativeModifyAPIResponse struct {
    model.CommonResponse
    TaobaoTanxAuditCreativeModifyResponse
}

type TaobaoTanxAuditCreativeModifyResponse struct {
    XMLName xml.Name `xml:"tanx_audit_creative_modify_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用的成功信息或失败信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 调用返回码
    
    TanxErrorCode   int64 `json:"tanx_error_code,omitempty" xml:"tanx_error_code,omitempty"`

    
    // 是否成功
    
    IsOk   bool `json:"is_ok,omitempty" xml:"is_ok,omitempty"`

    
}
