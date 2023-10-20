package charity

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaValueCoinIssueAPIResponse 爱豆发放 API返回值
// alibaba.value.coin.issue
//
// 爱豆发放
type AlibabaValueCoinIssueAPIResponse struct {
	model.CommonResponse
	AlibabaValueCoinIssueAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaValueCoinIssueAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaValueCoinIssueAPIResponseModel).Reset()
}

// AlibabaValueCoinIssueAPIResponseModel is 爱豆发放 成功返回结果
type AlibabaValueCoinIssueAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_value_coin_issue_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	RespMsg string `json:"resp_msg,omitempty" xml:"resp_msg,omitempty"`
	// 错误码
	RespCode int64 `json:"resp_code,omitempty" xml:"resp_code,omitempty"`
	// null值
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	SuccessStatus bool `json:"success_status,omitempty" xml:"success_status,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaValueCoinIssueAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RespMsg = ""
	m.RespCode = 0
	m.Data = false
	m.SuccessStatus = false
}

var poolAlibabaValueCoinIssueAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaValueCoinIssueAPIResponse)
	},
}

// GetAlibabaValueCoinIssueAPIResponse 从 sync.Pool 获取 AlibabaValueCoinIssueAPIResponse
func GetAlibabaValueCoinIssueAPIResponse() *AlibabaValueCoinIssueAPIResponse {
	return poolAlibabaValueCoinIssueAPIResponse.Get().(*AlibabaValueCoinIssueAPIResponse)
}

// ReleaseAlibabaValueCoinIssueAPIResponse 将 AlibabaValueCoinIssueAPIResponse 保存到 sync.Pool
func ReleaseAlibabaValueCoinIssueAPIResponse(v *AlibabaValueCoinIssueAPIResponse) {
	v.Reset()
	poolAlibabaValueCoinIssueAPIResponse.Put(v)
}
