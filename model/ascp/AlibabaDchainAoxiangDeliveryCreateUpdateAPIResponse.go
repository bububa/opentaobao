package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangDeliveryCreateUpdateAPIResponse 新建/更新配资源 API返回值
// alibaba.dchain.aoxiang.delivery.create.update
//
// 新建/更新配资源
type AlibabaDchainAoxiangDeliveryCreateUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangDeliveryCreateUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangDeliveryCreateUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangDeliveryCreateUpdateAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangDeliveryCreateUpdateAPIResponseModel is 新建/更新配资源 成功返回结果
type AlibabaDchainAoxiangDeliveryCreateUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_delivery_create_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 新建/更新配资源出参
	DeliveryUpsertResponse *DeliveryUpsertResponse `json:"delivery_upsert_response,omitempty" xml:"delivery_upsert_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangDeliveryCreateUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DeliveryUpsertResponse = nil
}

var poolAlibabaDchainAoxiangDeliveryCreateUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangDeliveryCreateUpdateAPIResponse)
	},
}

// GetAlibabaDchainAoxiangDeliveryCreateUpdateAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangDeliveryCreateUpdateAPIResponse
func GetAlibabaDchainAoxiangDeliveryCreateUpdateAPIResponse() *AlibabaDchainAoxiangDeliveryCreateUpdateAPIResponse {
	return poolAlibabaDchainAoxiangDeliveryCreateUpdateAPIResponse.Get().(*AlibabaDchainAoxiangDeliveryCreateUpdateAPIResponse)
}

// ReleaseAlibabaDchainAoxiangDeliveryCreateUpdateAPIResponse 将 AlibabaDchainAoxiangDeliveryCreateUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangDeliveryCreateUpdateAPIResponse(v *AlibabaDchainAoxiangDeliveryCreateUpdateAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangDeliveryCreateUpdateAPIResponse.Put(v)
}
