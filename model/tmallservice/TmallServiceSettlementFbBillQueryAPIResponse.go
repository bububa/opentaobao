package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallServiceSettlementFbBillQueryAPIResponse 服务商工单结算对账查询 API返回值
// tmall.service.settlement.fb.bill.query
//
// 服务商工单结算对账查询，用于查询服务工单对应的结算费用情况，含工单对应的服务费、退款、增加费用、分成费用、提现流水
type TmallServiceSettlementFbBillQueryAPIResponse struct {
	model.CommonResponse
	TmallServiceSettlementFbBillQueryAPIResponseModel
}

// TmallServiceSettlementFbBillQueryAPIResponseModel is 服务商工单结算对账查询 成功返回结果
type TmallServiceSettlementFbBillQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_service_settlement_fb_bill_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 错误码
	RetCode string `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 分页数据
	Obj *PagedResult `json:"obj,omitempty" xml:"obj,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
