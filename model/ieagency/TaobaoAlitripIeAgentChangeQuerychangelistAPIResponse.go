package ieagency

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripieagentchangequerychangelistAPIResponse 卖家查询改签列表 API返回值
// taobao.alitrip.ie.agent.change.querychangelist
//
// 提供B2B卖家查看改签列表服务
type TaobaoalitripieagentchangequerychangelistAPIResponse struct {
	model.CommonResponse
	TaobaoalitripieagentchangequerychangelistAPIResponseModel
}

// TaobaoalitripieagentchangequerychangelistAPIResponseModel is 卖家查询改签列表 成功返回结果
type TaobaoalitripieagentchangequerychangelistAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ie_agent_change_querychangelist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *QueryChangeAgentListRs `json:"result,omitempty" xml:"result,omitempty"`
}
