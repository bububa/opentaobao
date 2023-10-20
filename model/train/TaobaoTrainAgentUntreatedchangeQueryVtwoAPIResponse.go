package train

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentUntreatedchangeQueryVtwoAPIResponse 火车票代理商接口-查询待处理改签单列表-含鉴权校验 API返回值
// taobao.train.agent.untreatedchange.query.vtwo
//
// 火车票代理商接口-查询待处理改签单列表-含鉴权校验
type TaobaoTrainAgentUntreatedchangeQueryVtwoAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentUntreatedchangeQueryVtwoAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTrainAgentUntreatedchangeQueryVtwoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTrainAgentUntreatedchangeQueryVtwoAPIResponseModel).Reset()
}

// TaobaoTrainAgentUntreatedchangeQueryVtwoAPIResponseModel is 火车票代理商接口-查询待处理改签单列表-含鉴权校验 成功返回结果
type TaobaoTrainAgentUntreatedchangeQueryVtwoAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_untreatedchange_query_vtwo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 待处理改签申请单列表
	UntreatedChangeApplyIds string `json:"untreated_change_apply_ids,omitempty" xml:"untreated_change_apply_ids,omitempty"`
	// 待处理改签单数量
	UntreatedChangeOrderNum int64 `json:"untreated_change_order_num,omitempty" xml:"untreated_change_order_num,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTrainAgentUntreatedchangeQueryVtwoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.UntreatedChangeApplyIds = ""
	m.UntreatedChangeOrderNum = 0
}

var poolTaobaoTrainAgentUntreatedchangeQueryVtwoAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTrainAgentUntreatedchangeQueryVtwoAPIResponse)
	},
}

// GetTaobaoTrainAgentUntreatedchangeQueryVtwoAPIResponse 从 sync.Pool 获取 TaobaoTrainAgentUntreatedchangeQueryVtwoAPIResponse
func GetTaobaoTrainAgentUntreatedchangeQueryVtwoAPIResponse() *TaobaoTrainAgentUntreatedchangeQueryVtwoAPIResponse {
	return poolTaobaoTrainAgentUntreatedchangeQueryVtwoAPIResponse.Get().(*TaobaoTrainAgentUntreatedchangeQueryVtwoAPIResponse)
}

// ReleaseTaobaoTrainAgentUntreatedchangeQueryVtwoAPIResponse 将 TaobaoTrainAgentUntreatedchangeQueryVtwoAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTrainAgentUntreatedchangeQueryVtwoAPIResponse(v *TaobaoTrainAgentUntreatedchangeQueryVtwoAPIResponse) {
	v.Reset()
	poolTaobaoTrainAgentUntreatedchangeQueryVtwoAPIResponse.Put(v)
}
