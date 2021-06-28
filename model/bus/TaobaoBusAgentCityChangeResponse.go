package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
城市变更 APIResponse
taobao.bus.agent.city.change

代理商通知城市变更，比如可售变为不可售等
*/
type TaobaoBusAgentCityChangeAPIResponse struct {
    model.CommonResponse
    TaobaoBusAgentCityChangeResponse
}

type TaobaoBusAgentCityChangeResponse struct {
    XMLName xml.Name `xml:"bus_agent_city_change_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 系统异常错误码
    
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`

    
    // 异常描述
    
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`

    
}
