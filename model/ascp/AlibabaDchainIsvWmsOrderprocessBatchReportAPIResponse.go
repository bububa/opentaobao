package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainisvwmsorderprocessbatchreportAPIResponse 仓作业信息批量同步 API返回值
// alibaba.dchain.isv.wms.orderprocess.batch.report
//
// 仓作业信息批量同步
type AlibabadchainisvwmsorderprocessbatchreportAPIResponse struct {
	model.CommonResponse
	AlibabadchainisvwmsorderprocessbatchreportAPIResponseModel
}

// AlibabadchainisvwmsorderprocessbatchreportAPIResponseModel is 仓作业信息批量同步 成功返回结果
type AlibabadchainisvwmsorderprocessbatchreportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_isv_wms_orderprocess_batch_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求出参
	OrderProcessBatchReportResponse *WmsOrderProcessBatchReportResponse `json:"order_process_batch_report_response,omitempty" xml:"order_process_batch_report_response,omitempty"`
}
