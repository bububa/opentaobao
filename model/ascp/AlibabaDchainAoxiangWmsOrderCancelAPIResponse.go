package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangWmsOrderCancelAPIResponse 回传发货单取消通知 API返回值
// alibaba.dchain.aoxiang.wms.order.cancel
//
// 回传发货单取消通知
type AlibabaDchainAoxiangWmsOrderCancelAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangWmsOrderCancelAPIResponseModel
}

// AlibabaDchainAoxiangWmsOrderCancelAPIResponseModel is 回传发货单取消通知 成功返回结果
type AlibabaDchainAoxiangWmsOrderCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_wms_order_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 回传结果
	OrderCancelReportResponse *OrderCancelReportResponse `json:"order_cancel_report_response,omitempty" xml:"order_cancel_report_response,omitempty"`
}
