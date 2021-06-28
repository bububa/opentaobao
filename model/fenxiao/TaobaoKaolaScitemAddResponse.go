package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
考拉货品新增接口 APIResponse
taobao.kaola.scitem.add

考拉货品新增接口
*/
type TaobaoKaolaScitemAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoKaolaScitemAddResponse `json:"kaola_scitem_add_response,omitempty"` 
    TaobaoKaolaScitemAddResponse
}

/* model for simplify = false
type TaobaoKaolaScitemAddResponse struct {

    // 异常信息
    
    ErrorMessages  struct {
        String  []string `json:"string,omitempty"`
    } `json:"error_messages,omitempty"`
    

    // 货品id
    
    Data   int64 `json:"data,omitempty"`
    

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 是否系统异常
    
    IsSystemFailed   bool `json:"is_system_failed,omitempty"`
    

    // 异常信息Code
    
    SysErrorCode   string `json:"sys_error_code,omitempty"`
    

}
*/

type TaobaoKaolaScitemAddResponse struct {

    // 异常信息
    ErrorMessages   []string `json:"error_messages,omitempty"`

    // 货品id
    Data   int64 `json:"data,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 是否系统异常
    IsSystemFailed   bool `json:"is_system_failed,omitempty"`

    // 异常信息Code
    SysErrorCode   string `json:"sys_error_code,omitempty"`

}
