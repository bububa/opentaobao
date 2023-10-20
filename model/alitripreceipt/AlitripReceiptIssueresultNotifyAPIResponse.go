package alitripreceipt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripReceiptIssueresultNotifyAPIResponse 飞猪发票开票结果通知 API返回值
// alitrip.receipt.issueresult.notify
//
// 飞猪发票开票结果通知
type AlitripReceiptIssueresultNotifyAPIResponse struct {
	model.CommonResponse
	AlitripReceiptIssueresultNotifyAPIResponseModel
}

// AlitripReceiptIssueresultNotifyAPIResponseModel is 飞猪发票开票结果通知 成功返回结果
type AlitripReceiptIssueresultNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_receipt_issueresult_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否开票成功
	IsSuccess string `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 错误code
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误原因
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
}
