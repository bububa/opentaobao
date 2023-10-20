package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainIsvWmsOrderprocessBatchReportAPIResponse 仓作业信息批量同步 API返回值
// alibaba.dchain.isv.wms.orderprocess.batch.report
//
// 仓作业信息批量同步
type AlibabaDchainIsvWmsOrderprocessBatchReportAPIResponse struct {
	model.CommonResponse
	AlibabaDchainIsvWmsOrderprocessBatchReportAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainIsvWmsOrderprocessBatchReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainIsvWmsOrderprocessBatchReportAPIResponseModel).Reset()
}

// AlibabaDchainIsvWmsOrderprocessBatchReportAPIResponseModel is 仓作业信息批量同步 成功返回结果
type AlibabaDchainIsvWmsOrderprocessBatchReportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_isv_wms_orderprocess_batch_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求出参
	OrderProcessBatchReportResponse *WmsOrderProcessBatchReportResponse `json:"order_process_batch_report_response,omitempty" xml:"order_process_batch_report_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainIsvWmsOrderprocessBatchReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OrderProcessBatchReportResponse = nil
}

var poolAlibabaDchainIsvWmsOrderprocessBatchReportAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainIsvWmsOrderprocessBatchReportAPIResponse)
	},
}

// GetAlibabaDchainIsvWmsOrderprocessBatchReportAPIResponse 从 sync.Pool 获取 AlibabaDchainIsvWmsOrderprocessBatchReportAPIResponse
func GetAlibabaDchainIsvWmsOrderprocessBatchReportAPIResponse() *AlibabaDchainIsvWmsOrderprocessBatchReportAPIResponse {
	return poolAlibabaDchainIsvWmsOrderprocessBatchReportAPIResponse.Get().(*AlibabaDchainIsvWmsOrderprocessBatchReportAPIResponse)
}

// ReleaseAlibabaDchainIsvWmsOrderprocessBatchReportAPIResponse 将 AlibabaDchainIsvWmsOrderprocessBatchReportAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainIsvWmsOrderprocessBatchReportAPIResponse(v *AlibabaDchainIsvWmsOrderprocessBatchReportAPIResponse) {
	v.Reset()
	poolAlibabaDchainIsvWmsOrderprocessBatchReportAPIResponse.Put(v)
}
