package tanx

import (
    "github.com/bububa/opentaobao/model"
)

/* 
dsp托管创意新增接口 APIResponse
taobao.tanx.audit.depositcreative.add

dsp托管创意新增接口
*/
type TaobaoTanxAuditDepositcreativeAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTanxAuditDepositcreativeAddResponse `json:"tanx_audit_depositcreative_add_response,omitempty"` 
    TaobaoTanxAuditDepositcreativeAddResponse
}

/* model for simplify = false
type TaobaoTanxAuditDepositcreativeAddResponse struct {

    // 调用的成功信息或失败信息
    
    Message   string `json:"message,omitempty"`
    

    // 调用返回码
    
    TanxErrorCode   int64 `json:"tanx_error_code,omitempty"`
    

    // 是否成功
    
    IsOk   bool `json:"is_ok,omitempty"`
    

}
*/

type TaobaoTanxAuditDepositcreativeAddResponse struct {

    // 调用的成功信息或失败信息
    Message   string `json:"message,omitempty"`

    // 调用返回码
    TanxErrorCode   int64 `json:"tanx_error_code,omitempty"`

    // 是否成功
    IsOk   bool `json:"is_ok,omitempty"`

}
