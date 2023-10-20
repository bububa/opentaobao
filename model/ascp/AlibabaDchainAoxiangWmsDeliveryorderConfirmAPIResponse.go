package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangwmsdeliveryorderconfirmAPIResponse 回传发货单确认 API返回值
// alibaba.dchain.aoxiang.wms.deliveryorder.confirm
//
// 回传发货单确认
type AlibabadchainaoxiangwmsdeliveryorderconfirmAPIResponse struct {
	model.CommonResponse
	AlibabadchainaoxiangwmsdeliveryorderconfirmAPIResponseModel
}

// AlibabadchainaoxiangwmsdeliveryorderconfirmAPIResponseModel is 回传发货单确认 成功返回结果
type AlibabadchainaoxiangwmsdeliveryorderconfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_wms_deliveryorder_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 回传结果
	DeliveryOrderConfirmReportResponse *DeliveryOrderConfirmReportResponse `json:"delivery_order_confirm_report_response,omitempty" xml:"delivery_order_confirm_report_response,omitempty"`
}
