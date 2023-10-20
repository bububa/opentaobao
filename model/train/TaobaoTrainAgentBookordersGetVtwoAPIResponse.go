package train

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentBookordersGetVtwoAPIResponse 代理商获取待出票订单列表v2--增加鉴权校验 API返回值
// taobao.train.agent.bookorders.get.vtwo
//
// 代理商获取待出票订单列表，只返回订单号
type TaobaoTrainAgentBookordersGetVtwoAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentBookordersGetVtwoAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTrainAgentBookordersGetVtwoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTrainAgentBookordersGetVtwoAPIResponseModel).Reset()
}

// TaobaoTrainAgentBookordersGetVtwoAPIResponseModel is 代理商获取待出票订单列表v2--增加鉴权校验 成功返回结果
type TaobaoTrainAgentBookordersGetVtwoAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_bookorders_get_vtwo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单号集合，用半角逗号(,)连接，只会返回固定数量
	OrderIds string `json:"order_ids,omitempty" xml:"order_ids,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 待处理订单总数
	OrderCount int64 `json:"order_count,omitempty" xml:"order_count,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTrainAgentBookordersGetVtwoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OrderIds = ""
	m.ErrorMsg = ""
	m.OrderCount = 0
}

var poolTaobaoTrainAgentBookordersGetVtwoAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTrainAgentBookordersGetVtwoAPIResponse)
	},
}

// GetTaobaoTrainAgentBookordersGetVtwoAPIResponse 从 sync.Pool 获取 TaobaoTrainAgentBookordersGetVtwoAPIResponse
func GetTaobaoTrainAgentBookordersGetVtwoAPIResponse() *TaobaoTrainAgentBookordersGetVtwoAPIResponse {
	return poolTaobaoTrainAgentBookordersGetVtwoAPIResponse.Get().(*TaobaoTrainAgentBookordersGetVtwoAPIResponse)
}

// ReleaseTaobaoTrainAgentBookordersGetVtwoAPIResponse 将 TaobaoTrainAgentBookordersGetVtwoAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTrainAgentBookordersGetVtwoAPIResponse(v *TaobaoTrainAgentBookordersGetVtwoAPIResponse) {
	v.Reset()
	poolTaobaoTrainAgentBookordersGetVtwoAPIResponse.Put(v)
}
