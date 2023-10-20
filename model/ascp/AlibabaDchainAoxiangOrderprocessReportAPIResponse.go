package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangorderprocessreportAPIResponse 回传仓内作业节点 API返回值
// alibaba.dchain.aoxiang.orderprocess.report
//
// 回传仓内作业节点
type AlibabadchainaoxiangorderprocessreportAPIResponse struct {
	model.CommonResponse
	AlibabadchainaoxiangorderprocessreportAPIResponseModel
}

// AlibabadchainaoxiangorderprocessreportAPIResponseModel is 回传仓内作业节点 成功返回结果
type AlibabadchainaoxiangorderprocessreportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_orderprocess_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 回传结果
	OrderprocessReportResponse *OrderProcessReportResponse `json:"orderprocess_report_response,omitempty" xml:"orderprocess_report_response,omitempty"`
}
