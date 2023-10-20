package tblogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIResponse 同城配送在线下单获取配送单 API返回值
// alibaba.ascp.logistics.instantsonline.deliveryorder.get
//
// 同城配送在线下单获取配送单
type AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIResponse struct {
	model.CommonResponse
	AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIResponseModel).Reset()
}

// AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIResponseModel is 同城配送在线下单获取配送单 成功返回结果
type AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_logistics_instantsonline_deliveryorder_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *AlibabaAscpLogisticsInstantsonlineDeliveryorderGetTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIResponse)
	},
}

// GetAlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIResponse 从 sync.Pool 获取 AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIResponse
func GetAlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIResponse() *AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIResponse {
	return poolAlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIResponse.Get().(*AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIResponse)
}

// ReleaseAlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIResponse 将 AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIResponse(v *AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIResponse) {
	v.Reset()
	poolAlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIResponse.Put(v)
}
