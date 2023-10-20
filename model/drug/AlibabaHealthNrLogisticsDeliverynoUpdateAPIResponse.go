package drug

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthNrLogisticsDeliverynoUpdateAPIResponse 上传订单同城快递单号 API返回值
// alibaba.health.nr.logistics.deliveryno.update
//
// 上传订单同城快递单号
type AlibabaHealthNrLogisticsDeliverynoUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaHealthNrLogisticsDeliverynoUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHealthNrLogisticsDeliverynoUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHealthNrLogisticsDeliverynoUpdateAPIResponseModel).Reset()
}

// AlibabaHealthNrLogisticsDeliverynoUpdateAPIResponseModel is 上传订单同城快递单号 成功返回结果
type AlibabaHealthNrLogisticsDeliverynoUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_health_nr_logistics_deliveryno_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 更新成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHealthNrLogisticsDeliverynoUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAlibabaHealthNrLogisticsDeliverynoUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHealthNrLogisticsDeliverynoUpdateAPIResponse)
	},
}

// GetAlibabaHealthNrLogisticsDeliverynoUpdateAPIResponse 从 sync.Pool 获取 AlibabaHealthNrLogisticsDeliverynoUpdateAPIResponse
func GetAlibabaHealthNrLogisticsDeliverynoUpdateAPIResponse() *AlibabaHealthNrLogisticsDeliverynoUpdateAPIResponse {
	return poolAlibabaHealthNrLogisticsDeliverynoUpdateAPIResponse.Get().(*AlibabaHealthNrLogisticsDeliverynoUpdateAPIResponse)
}

// ReleaseAlibabaHealthNrLogisticsDeliverynoUpdateAPIResponse 将 AlibabaHealthNrLogisticsDeliverynoUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHealthNrLogisticsDeliverynoUpdateAPIResponse(v *AlibabaHealthNrLogisticsDeliverynoUpdateAPIResponse) {
	v.Reset()
	poolAlibabaHealthNrLogisticsDeliverynoUpdateAPIResponse.Put(v)
}
