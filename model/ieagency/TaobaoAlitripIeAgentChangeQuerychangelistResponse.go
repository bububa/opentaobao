package ieagency

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家查询改签列表 APIResponse
taobao.alitrip.ie.agent.change.querychangelist

提供B2B卖家查看改签列表服务
*/
type TaobaoAlitripIeAgentChangeQuerychangelistAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripIeAgentChangeQuerychangelistResponse
}

type TaobaoAlitripIeAgentChangeQuerychangelistResponse struct {
    XMLName xml.Name `xml:"alitrip_ie_agent_change_querychangelist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *QueryChangeAgentListRs `json:"result,omitempty" xml:"result,omitempty"`

    
}
