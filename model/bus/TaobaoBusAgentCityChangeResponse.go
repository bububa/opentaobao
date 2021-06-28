package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
城市变更 APIResponse
taobao.bus.agent.city.change

代理商通知城市变更，比如可售变为不可售等
*/
type TaobaoBusAgentCityChangeAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBusAgentCityChangeResponse `json:"bus_agent_city_change_response,omitempty"` 
    TaobaoBusAgentCityChangeResponse
}

/* model for simplify = false
type TaobaoBusAgentCityChangeResponse struct {

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 系统异常错误码
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 异常描述
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

}
*/

type TaobaoBusAgentCityChangeResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 系统异常错误码
    ErrCode   string `json:"err_code,omitempty"`

    // 异常描述
    ErrMsg   string `json:"err_msg,omitempty"`

}
