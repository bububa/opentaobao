package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainisvwmsorderprocessreportAPIResponse 仓作业信息同步 API返回值
// alibaba.dchain.isv.wms.orderprocess.report
//
// 仓作业信息同步
type AlibabadchainisvwmsorderprocessreportAPIResponse struct {
	model.CommonResponse
	AlibabadchainisvwmsorderprocessreportAPIResponseModel
}

// AlibabadchainisvwmsorderprocessreportAPIResponseModel is 仓作业信息同步 成功返回结果
type AlibabadchainisvwmsorderprocessreportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_isv_wms_orderprocess_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求出参
	OrderProcessReportResponse *WmsOrderProcessReportResponse `json:"order_process_report_response,omitempty" xml:"order_process_report_response,omitempty"`
}
