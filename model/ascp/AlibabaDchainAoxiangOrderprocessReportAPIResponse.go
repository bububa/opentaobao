package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangOrderprocessReportAPIResponse 回传仓内作业节点 API返回值
// alibaba.dchain.aoxiang.orderprocess.report
//
// 回传仓内作业节点
type AlibabaDchainAoxiangOrderprocessReportAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangOrderprocessReportAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangOrderprocessReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangOrderprocessReportAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangOrderprocessReportAPIResponseModel is 回传仓内作业节点 成功返回结果
type AlibabaDchainAoxiangOrderprocessReportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_orderprocess_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 回传结果
	OrderprocessReportResponse *OrderProcessReportResponse `json:"orderprocess_report_response,omitempty" xml:"orderprocess_report_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangOrderprocessReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OrderprocessReportResponse = nil
}

var poolAlibabaDchainAoxiangOrderprocessReportAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangOrderprocessReportAPIResponse)
	},
}

// GetAlibabaDchainAoxiangOrderprocessReportAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangOrderprocessReportAPIResponse
func GetAlibabaDchainAoxiangOrderprocessReportAPIResponse() *AlibabaDchainAoxiangOrderprocessReportAPIResponse {
	return poolAlibabaDchainAoxiangOrderprocessReportAPIResponse.Get().(*AlibabaDchainAoxiangOrderprocessReportAPIResponse)
}

// ReleaseAlibabaDchainAoxiangOrderprocessReportAPIResponse 将 AlibabaDchainAoxiangOrderprocessReportAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangOrderprocessReportAPIResponse(v *AlibabaDchainAoxiangOrderprocessReportAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangOrderprocessReportAPIResponse.Put(v)
}
