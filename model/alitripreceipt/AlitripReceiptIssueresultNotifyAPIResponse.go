package alitripreceipt

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlitripReceiptIssueresultNotifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripReceiptIssueresultNotifyAPIResponseModel).Reset()
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

// Reset 清空结构体
func (m *AlitripReceiptIssueresultNotifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = ""
	m.ResultCode = ""
	m.ResultMsg = ""
}

var poolAlitripReceiptIssueresultNotifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripReceiptIssueresultNotifyAPIResponse)
	},
}

// GetAlitripReceiptIssueresultNotifyAPIResponse 从 sync.Pool 获取 AlitripReceiptIssueresultNotifyAPIResponse
func GetAlitripReceiptIssueresultNotifyAPIResponse() *AlitripReceiptIssueresultNotifyAPIResponse {
	return poolAlitripReceiptIssueresultNotifyAPIResponse.Get().(*AlitripReceiptIssueresultNotifyAPIResponse)
}

// ReleaseAlitripReceiptIssueresultNotifyAPIResponse 将 AlitripReceiptIssueresultNotifyAPIResponse 保存到 sync.Pool
func ReleaseAlitripReceiptIssueresultNotifyAPIResponse(v *AlitripReceiptIssueresultNotifyAPIResponse) {
	v.Reset()
	poolAlitripReceiptIssueresultNotifyAPIResponse.Put(v)
}
