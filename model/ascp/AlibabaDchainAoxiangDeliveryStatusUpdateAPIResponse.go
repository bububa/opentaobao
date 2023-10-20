package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangDeliveryStatusUpdateAPIResponse 启用/停用配资源 API返回值
// alibaba.dchain.aoxiang.delivery.status.update
//
// 启用/停用配资源
type AlibabaDchainAoxiangDeliveryStatusUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangDeliveryStatusUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangDeliveryStatusUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangDeliveryStatusUpdateAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangDeliveryStatusUpdateAPIResponseModel is 启用/停用配资源 成功返回结果
type AlibabaDchainAoxiangDeliveryStatusUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_delivery_status_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 启用/停用配资源出参
	DeliveryStatusUpdateResponse *DeliveryStatusUpdateResponse `json:"delivery_status_update_response,omitempty" xml:"delivery_status_update_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangDeliveryStatusUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DeliveryStatusUpdateResponse = nil
}

var poolAlibabaDchainAoxiangDeliveryStatusUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangDeliveryStatusUpdateAPIResponse)
	},
}

// GetAlibabaDchainAoxiangDeliveryStatusUpdateAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangDeliveryStatusUpdateAPIResponse
func GetAlibabaDchainAoxiangDeliveryStatusUpdateAPIResponse() *AlibabaDchainAoxiangDeliveryStatusUpdateAPIResponse {
	return poolAlibabaDchainAoxiangDeliveryStatusUpdateAPIResponse.Get().(*AlibabaDchainAoxiangDeliveryStatusUpdateAPIResponse)
}

// ReleaseAlibabaDchainAoxiangDeliveryStatusUpdateAPIResponse 将 AlibabaDchainAoxiangDeliveryStatusUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangDeliveryStatusUpdateAPIResponse(v *AlibabaDchainAoxiangDeliveryStatusUpdateAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangDeliveryStatusUpdateAPIResponse.Put(v)
}
