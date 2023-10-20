package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicesettlementfbbilldetailqueryAPIResponse 服务商工单结算对账查询-流水查询 API返回值
// tmall.service.settlement.fb.bill.detail.query
//
// 服务商工单结算对账查询-流水查询，用于查询服务工单费用流水，含服务费、退款、分成、提现等。
type TmallservicesettlementfbbilldetailqueryAPIResponse struct {
	model.CommonResponse
	TmallservicesettlementfbbilldetailqueryAPIResponseModel
}

// TmallservicesettlementfbbilldetailqueryAPIResponseModel is 服务商工单结算对账查询-流水查询 成功返回结果
type TmallservicesettlementfbbilldetailqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_service_settlement_fb_bill_detail_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误提示
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 错误码
	RetCode string `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 分页数据
	Obj *PagedResult `json:"obj,omitempty" xml:"obj,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
