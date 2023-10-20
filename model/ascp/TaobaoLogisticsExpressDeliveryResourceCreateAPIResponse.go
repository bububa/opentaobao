package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressDeliveryResourceCreateAPIResponse 新建/更新配资源 API返回值
// taobao.logistics.express.delivery.resource.create
//
// 新建/更新配资源
type TaobaoLogisticsExpressDeliveryResourceCreateAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsExpressDeliveryResourceCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressDeliveryResourceCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsExpressDeliveryResourceCreateAPIResponseModel).Reset()
}

// TaobaoLogisticsExpressDeliveryResourceCreateAPIResponseModel is 新建/更新配资源 成功返回结果
type TaobaoLogisticsExpressDeliveryResourceCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_delivery_resource_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	DeliveryResourceCreateResponse *DeliveryResourceCreateResponse `json:"delivery_resource_create_response,omitempty" xml:"delivery_resource_create_response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressDeliveryResourceCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DeliveryResourceCreateResponse = nil
}

var poolTaobaoLogisticsExpressDeliveryResourceCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsExpressDeliveryResourceCreateAPIResponse)
	},
}

// GetTaobaoLogisticsExpressDeliveryResourceCreateAPIResponse 从 sync.Pool 获取 TaobaoLogisticsExpressDeliveryResourceCreateAPIResponse
func GetTaobaoLogisticsExpressDeliveryResourceCreateAPIResponse() *TaobaoLogisticsExpressDeliveryResourceCreateAPIResponse {
	return poolTaobaoLogisticsExpressDeliveryResourceCreateAPIResponse.Get().(*TaobaoLogisticsExpressDeliveryResourceCreateAPIResponse)
}

// ReleaseTaobaoLogisticsExpressDeliveryResourceCreateAPIResponse 将 TaobaoLogisticsExpressDeliveryResourceCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsExpressDeliveryResourceCreateAPIResponse(v *TaobaoLogisticsExpressDeliveryResourceCreateAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsExpressDeliveryResourceCreateAPIResponse.Put(v)
}
