package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIResponse 回传发货单确认 API返回值
// alibaba.dchain.aoxiang.wms.deliveryorder.confirm
//
// 回传发货单确认
type AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIResponseModel is 回传发货单确认 成功返回结果
type AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_wms_deliveryorder_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 回传结果
	DeliveryOrderConfirmReportResponse *DeliveryOrderConfirmReportResponse `json:"delivery_order_confirm_report_response,omitempty" xml:"delivery_order_confirm_report_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DeliveryOrderConfirmReportResponse = nil
}

var poolAlibabaDchainAoxiangWmsDeliveryorderConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIResponse)
	},
}

// GetAlibabaDchainAoxiangWmsDeliveryorderConfirmAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIResponse
func GetAlibabaDchainAoxiangWmsDeliveryorderConfirmAPIResponse() *AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIResponse {
	return poolAlibabaDchainAoxiangWmsDeliveryorderConfirmAPIResponse.Get().(*AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIResponse)
}

// ReleaseAlibabaDchainAoxiangWmsDeliveryorderConfirmAPIResponse 将 AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangWmsDeliveryorderConfirmAPIResponse(v *AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangWmsDeliveryorderConfirmAPIResponse.Put(v)
}
