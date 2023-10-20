package ieagency

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentChangeQuerychangelistAPIResponse 卖家查询改签列表 API返回值
// taobao.alitrip.ie.agent.change.querychangelist
//
// 提供B2B卖家查看改签列表服务
type TaobaoAlitripIeAgentChangeQuerychangelistAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripIeAgentChangeQuerychangelistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripIeAgentChangeQuerychangelistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripIeAgentChangeQuerychangelistAPIResponseModel).Reset()
}

// TaobaoAlitripIeAgentChangeQuerychangelistAPIResponseModel is 卖家查询改签列表 成功返回结果
type TaobaoAlitripIeAgentChangeQuerychangelistAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ie_agent_change_querychangelist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *QueryChangeAgentListRs `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripIeAgentChangeQuerychangelistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripIeAgentChangeQuerychangelistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripIeAgentChangeQuerychangelistAPIResponse)
	},
}

// GetTaobaoAlitripIeAgentChangeQuerychangelistAPIResponse 从 sync.Pool 获取 TaobaoAlitripIeAgentChangeQuerychangelistAPIResponse
func GetTaobaoAlitripIeAgentChangeQuerychangelistAPIResponse() *TaobaoAlitripIeAgentChangeQuerychangelistAPIResponse {
	return poolTaobaoAlitripIeAgentChangeQuerychangelistAPIResponse.Get().(*TaobaoAlitripIeAgentChangeQuerychangelistAPIResponse)
}

// ReleaseTaobaoAlitripIeAgentChangeQuerychangelistAPIResponse 将 TaobaoAlitripIeAgentChangeQuerychangelistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripIeAgentChangeQuerychangelistAPIResponse(v *TaobaoAlitripIeAgentChangeQuerychangelistAPIResponse) {
	v.Reset()
	poolTaobaoAlitripIeAgentChangeQuerychangelistAPIResponse.Put(v)
}
