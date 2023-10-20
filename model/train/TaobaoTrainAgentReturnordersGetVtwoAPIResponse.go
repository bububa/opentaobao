package train

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentReturnordersGetVtwoAPIResponse 获取待退票的订单v2--增加鉴权校验 API返回值
// taobao.train.agent.returnorders.get.vtwo
//
// 代理商用来获取待退票的订单列表及数量，防止代理商掉单。
type TaobaoTrainAgentReturnordersGetVtwoAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentReturnordersGetVtwoAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTrainAgentReturnordersGetVtwoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTrainAgentReturnordersGetVtwoAPIResponseModel).Reset()
}

// TaobaoTrainAgentReturnordersGetVtwoAPIResponseModel is 获取待退票的订单v2--增加鉴权校验 成功返回结果
type TaobaoTrainAgentReturnordersGetVtwoAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_returnorders_get_vtwo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 子订单号字符串，每个订单以逗号分隔
	OrderIds string `json:"order_ids,omitempty" xml:"order_ids,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 主订单id列表
	MainOrderIds string `json:"main_order_ids,omitempty" xml:"main_order_ids,omitempty"`
	// 申请时间，每个时间以逗号分隔
	RefundApplyTimes string `json:"refund_apply_times,omitempty" xml:"refund_apply_times,omitempty"`
	// 待退票的订单数
	OrderCount int64 `json:"order_count,omitempty" xml:"order_count,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTrainAgentReturnordersGetVtwoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OrderIds = ""
	m.ErrorMsg = ""
	m.MainOrderIds = ""
	m.RefundApplyTimes = ""
	m.OrderCount = 0
}

var poolTaobaoTrainAgentReturnordersGetVtwoAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTrainAgentReturnordersGetVtwoAPIResponse)
	},
}

// GetTaobaoTrainAgentReturnordersGetVtwoAPIResponse 从 sync.Pool 获取 TaobaoTrainAgentReturnordersGetVtwoAPIResponse
func GetTaobaoTrainAgentReturnordersGetVtwoAPIResponse() *TaobaoTrainAgentReturnordersGetVtwoAPIResponse {
	return poolTaobaoTrainAgentReturnordersGetVtwoAPIResponse.Get().(*TaobaoTrainAgentReturnordersGetVtwoAPIResponse)
}

// ReleaseTaobaoTrainAgentReturnordersGetVtwoAPIResponse 将 TaobaoTrainAgentReturnordersGetVtwoAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTrainAgentReturnordersGetVtwoAPIResponse(v *TaobaoTrainAgentReturnordersGetVtwoAPIResponse) {
	v.Reset()
	poolTaobaoTrainAgentReturnordersGetVtwoAPIResponse.Put(v)
}
