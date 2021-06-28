package tanx

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创意预审新增接口 APIResponse
taobao.tanx.creative.add

创意预审新增接口
*/
type TaobaoTanxCreativeAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTanxCreativeAddResponse `json:"tanx_creative_add_response,omitempty"` 
    TaobaoTanxCreativeAddResponse
}

/* model for simplify = false
type TaobaoTanxCreativeAddResponse struct {

    // 调用的成功信息或失败信息
    
    Message   string `json:"message,omitempty"`
    

    // 调用返回码
    
    Code   int64 `json:"code,omitempty"`
    

    // 是否成功
    
    IsOk   bool `json:"is_ok,omitempty"`
    

}
*/

type TaobaoTanxCreativeAddResponse struct {

    // 调用的成功信息或失败信息
    Message   string `json:"message,omitempty"`

    // 调用返回码
    Code   int64 `json:"code,omitempty"`

    // 是否成功
    IsOk   bool `json:"is_ok,omitempty"`

}
