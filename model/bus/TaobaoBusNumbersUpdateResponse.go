package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
汽车票车次更新服务 APIResponse
taobao.bus.numbers.update

用于汽车票车次信息的新增、更新和逻辑删除
*/
type TaobaoBusNumbersUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBusNumbersUpdateResponse `json:"bus_numbers_update_response,omitempty"` 
    TaobaoBusNumbersUpdateResponse
}

/* model for simplify = false
type TaobaoBusNumbersUpdateResponse struct {

    // 错误码
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // 错误描述
    
    ResultMsg   string `json:"result_msg,omitempty"`
    

    // 成功数量
    
    SuccCount   int64 `json:"succ_count,omitempty"`
    

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoBusNumbersUpdateResponse struct {

    // 错误码
    ResultCode   string `json:"result_code,omitempty"`

    // 错误描述
    ResultMsg   string `json:"result_msg,omitempty"`

    // 成功数量
    SuccCount   int64 `json:"succ_count,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
