package train

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentChangeorderdetailQueryVtwoAPIResponse 火车票代理商接口-查询跑腿改签订单详情-含鉴权校验 API返回值
// taobao.train.agent.changeorderdetail.query.vtwo
//
// 火车票代理商接口-查询跑腿改签订单详情-含鉴权校验
type TaobaoTrainAgentChangeorderdetailQueryVtwoAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentChangeorderdetailQueryVtwoAPIResponseModel
}

// TaobaoTrainAgentChangeorderdetailQueryVtwoAPIResponseModel is 火车票代理商接口-查询跑腿改签订单详情-含鉴权校验 成功返回结果
type TaobaoTrainAgentChangeorderdetailQueryVtwoAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_changeorderdetail_query_vtwo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 改签子单列表
	ChangeApplySubOrders []ChangeApplySubOrderDto `json:"change_apply_sub_orders,omitempty" xml:"change_apply_sub_orders>change_apply_sub_order_dto,omitempty"`
	// 改签申请单主单
	ChangeApplyOrder *ChangeApplyOrderDto `json:"change_apply_order,omitempty" xml:"change_apply_order,omitempty"`
}
