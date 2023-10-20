package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainIsvWmsOrderprocessReportAPIResponse 仓作业信息同步 API返回值
// alibaba.dchain.isv.wms.orderprocess.report
//
// 仓作业信息同步
type AlibabaDchainIsvWmsOrderprocessReportAPIResponse struct {
	model.CommonResponse
	AlibabaDchainIsvWmsOrderprocessReportAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainIsvWmsOrderprocessReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainIsvWmsOrderprocessReportAPIResponseModel).Reset()
}

// AlibabaDchainIsvWmsOrderprocessReportAPIResponseModel is 仓作业信息同步 成功返回结果
type AlibabaDchainIsvWmsOrderprocessReportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_isv_wms_orderprocess_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求出参
	OrderProcessReportResponse *WmsOrderProcessReportResponse `json:"order_process_report_response,omitempty" xml:"order_process_report_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainIsvWmsOrderprocessReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OrderProcessReportResponse = nil
}

var poolAlibabaDchainIsvWmsOrderprocessReportAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainIsvWmsOrderprocessReportAPIResponse)
	},
}

// GetAlibabaDchainIsvWmsOrderprocessReportAPIResponse 从 sync.Pool 获取 AlibabaDchainIsvWmsOrderprocessReportAPIResponse
func GetAlibabaDchainIsvWmsOrderprocessReportAPIResponse() *AlibabaDchainIsvWmsOrderprocessReportAPIResponse {
	return poolAlibabaDchainIsvWmsOrderprocessReportAPIResponse.Get().(*AlibabaDchainIsvWmsOrderprocessReportAPIResponse)
}

// ReleaseAlibabaDchainIsvWmsOrderprocessReportAPIResponse 将 AlibabaDchainIsvWmsOrderprocessReportAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainIsvWmsOrderprocessReportAPIResponse(v *AlibabaDchainIsvWmsOrderprocessReportAPIResponse) {
	v.Reset()
	poolAlibabaDchainIsvWmsOrderprocessReportAPIResponse.Put(v)
}
