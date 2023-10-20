package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangWmsOrderprocessReportAPIResponse 回传发货单流水通知 API返回值
// alibaba.dchain.aoxiang.wms.orderprocess.report
//
// 回传发货单流水通知
type AlibabaDchainAoxiangWmsOrderprocessReportAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangWmsOrderprocessReportAPIResponseModel
}

// AlibabaDchainAoxiangWmsOrderprocessReportAPIResponseModel is 回传发货单流水通知 成功返回结果
type AlibabaDchainAoxiangWmsOrderprocessReportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_wms_orderprocess_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 回传结果
	OrderProcessReportResponse *OrderProcessReportReponse `json:"order_process_report_response,omitempty" xml:"order_process_report_response,omitempty"`
}
