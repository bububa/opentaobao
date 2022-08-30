package alscmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscMerchantExtTicketCancelAPIResponse 凭证作废 API返回值
// alibaba.alsc.merchant.ext.ticket.cancel
//
// isv调用银行接口超时导致凭证信息同步失败,触发关单处理,调用作废接口
type AlibabaAlscMerchantExtTicketCancelAPIResponse struct {
	model.CommonResponse
	AlibabaAlscMerchantExtTicketCancelAPIResponseModel
}

// AlibabaAlscMerchantExtTicketCancelAPIResponseModel is 凭证作废 成功返回结果
type AlibabaAlscMerchantExtTicketCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_merchant_ext_ticket_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 凭证任务单Id
	TicketTaskId string `json:"ticket_task_id,omitempty" xml:"ticket_task_id,omitempty"`
	// 错误编码/success/fail
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 是否需要重试(返回true时重试)
	CanRetry bool `json:"can_retry,omitempty" xml:"can_retry,omitempty"`
}
