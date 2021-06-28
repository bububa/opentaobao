package tanx

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创意修改接口 APIResponse
taobao.tanx.audit.creative.modify

创意修改接口
*/
type TaobaoTanxAuditCreativeModifyAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTanxAuditCreativeModifyResponse `json:"tanx_audit_creative_modify_response,omitempty"` 
    TaobaoTanxAuditCreativeModifyResponse
}

/* model for simplify = false
type TaobaoTanxAuditCreativeModifyResponse struct {

    // 调用的成功信息或失败信息
    
    Message   string `json:"message,omitempty"`
    

    // 调用返回码
    
    TanxErrorCode   int64 `json:"tanx_error_code,omitempty"`
    

    // 是否成功
    
    IsOk   bool `json:"is_ok,omitempty"`
    

}
*/

type TaobaoTanxAuditCreativeModifyResponse struct {

    // 调用的成功信息或失败信息
    Message   string `json:"message,omitempty"`

    // 调用返回码
    TanxErrorCode   int64 `json:"tanx_error_code,omitempty"`

    // 是否成功
    IsOk   bool `json:"is_ok,omitempty"`

}
