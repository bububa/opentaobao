package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
设置开发环境证书 APIResponse
taobao.openim.ioscert.sandbox.set

设置开发环境证书
*/
type TaobaoOpenimIoscertSandboxSetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOpenimIoscertSandboxSetResponse `json:"openim_ioscert_sandbox_set_response,omitempty"` 
    TaobaoOpenimIoscertSandboxSetResponse
}

/* model for simplify = false
type TaobaoOpenimIoscertSandboxSetResponse struct {

    // 操作成功
    
    Code   string `json:"code,omitempty"`
    

}
*/

type TaobaoOpenimIoscertSandboxSetResponse struct {

    // 操作成功
    Code   string `json:"code,omitempty"`

}
