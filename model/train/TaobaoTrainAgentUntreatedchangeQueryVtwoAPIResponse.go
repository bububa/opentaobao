package train

import (
	"encoding/xml"

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
