package tanx

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创意预审新增接口 APIResponse
taobao.tanx.audit.creative.add

创意预审新增接口
*/
type TaobaoTanxAuditCreativeAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTanxAuditCreativeAddResponse `json:"tanx_audit_creative_add_response,omitempty"` 
    TaobaoTanxAuditCreativeAddResponse
}

/* model for simplify = false
type TaobaoTanxAuditCreativeAddResponse struct {

    // 调用的成功信息或失败信息
    
    Message   string `json:"message,omitempty"`
    

    // 调用返回码
    
    TanxErrorCode   int64 `json:"tanx_error_code,omitempty"`
    

    // 是否成功
    
    IsOk   bool `json:"is_ok,omitempty"`
    

}
*/

type TaobaoTanxAuditCreativeAddResponse struct {

    // 调用的成功信息或失败信息
    Message   string `json:"message,omitempty"`

    // 调用返回码
    TanxErrorCode   int64 `json:"tanx_error_code,omitempty"`

    // 是否成功
    IsOk   bool `json:"is_ok,omitempty"`

}
