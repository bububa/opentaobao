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
	RequestId     string         `json:"request_id,omitempty" xml:"bus_agent_city_change_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"