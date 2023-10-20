package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangWmsDeliveryorderCreateAPIResponse 回传仓库接单通知 API返回值
// alibaba.dchain.aoxiang.wms.deliveryorder.create
//
// WMS上报仓库接单节点状态信息，代表接单环节。
type AlibabaDchainAoxiangWmsDeliveryorderCreateAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangWmsDeliveryorderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangWmsDeliveryorderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangWmsDeliveryorderCreateAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangWmsDeliveryorderCreateAPIResponseModel is 回传仓库接单通知 成功返回结果
type AlibabaDchainAoxiangWmsDeliveryorderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_wms_deliveryorder_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 回传结果
	DeliveryOrderReportResponse *DeliveryOrderReportResponse `json:"delivery_order_report_response,omitempty" xml:"delivery_order_report_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangWmsDeliveryorderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DeliveryOrderReportResponse = nil
}

var poolAlibabaDchainAoxiangWmsDeliveryorderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangWmsDeliveryorderCreateAPIResponse)
	},
}

// GetAlibabaDchainAoxiangWmsDeliveryorderCreateAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangWmsDeliveryorderCreateAPIResponse
func GetAlibabaDchainAoxiangWmsDeliveryorderCreateAPIResponse() *AlibabaDchainAoxiangWmsDeliveryorderCreateAPIResponse {
	return poolAlibabaDchainAoxiangWmsDeliveryorderCreateAPIResponse.Get().(*AlibabaDchainAoxiangWmsDeliveryorderCreateAPIResponse)
}

// ReleaseAlibabaDchainAoxiangWmsDeliveryorderCreateAPIResponse 将 AlibabaDchainAoxiangWmsDeliveryorderCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangWmsDeliveryorderCreateAPIResponse(v *AlibabaDchainAoxiangWmsDeliveryorderCreateAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangWmsDeliveryorderCreateAPIResponse.Put(v)
}
