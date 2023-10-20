package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdReportGetAccountReportAPIResponse 账户报告 API返回值
// alibaba.scbp.ad.report.get.account.report
//
// 账户报告
type AlibabaScbpAdReportGetAccountReportAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdReportGetAccountReportAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdReportGetAccountReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdReportGetAccountReportAPIResponseModel).Reset()
}

// AlibabaScbpAdReportGetAccountReportAPIResponseModel is 账户报告 成功返回结果
type AlibabaScbpAdReportGetAccountReportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_report_get_account_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回参数
	Result *AccountReportDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdReportGetAccountReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaScbpAdReportGetAccountReportAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdReportGetAccountReportAPIResponse)
	},
}

// GetAlibabaScbpAdReportGetAccountReportAPIResponse 从 sync.Pool 获取 AlibabaScbpAdReportGetAccountReportAPIResponse
func GetAlibabaScbpAdReportGetAccountReportAPIResponse() *AlibabaScbpAdReportGetAccountReportAPIResponse {
	return poolAlibabaScbpAdReportGetAccountReportAPIResponse.Get().(*AlibabaScbpAdReportGetAccountReportAPIResponse)
}

// ReleaseAlibabaScbpAdReportGetAccountReportAPIResponse 将 AlibabaScbpAdReportGetAccountReportAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdReportGetAccountReportAPIResponse(v *AlibabaScbpAdReportGetAccountReportAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdReportGetAccountReportAPIResponse.Put(v)
}
