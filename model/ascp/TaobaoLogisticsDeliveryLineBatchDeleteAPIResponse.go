package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsDeliveryLineBatchDeleteAPIResponse 线路能力删除 API返回值
// taobao.logistics.delivery.line.batch.delete
//
// 线路能力删除
type TaobaoLogisticsDeliveryLineBatchDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsDeliveryLineBatchDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsDeliveryLineBatchDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsDeliveryLineBatchDeleteAPIResponseModel).Reset()
}

// TaobaoLogisticsDeliveryLineBatchDeleteAPIResponseModel is 线路能力删除 成功返回结果
type TaobaoLogisticsDeliveryLineBatchDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_delivery_line_batch_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 线路能力删除出参
	DeliveryLineBatchDeleteResponse *DeliveryLineBatchDeleteResponse `json:"delivery_line_batch_delete_response,omitempty" xml:"delivery_line_batch_delete_response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsDeliveryLineBatchDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DeliveryLineBatchDeleteResponse = nil
}

var poolTaobaoLogisticsDeliveryLineBatchDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsDeliveryLineBatchDeleteAPIResponse)
	},
}

// GetTaobaoLogisticsDeliveryLineBatchDeleteAPIResponse 从 sync.Pool 获取 TaobaoLogisticsDeliveryLineBatchDeleteAPIResponse
func GetTaobaoLogisticsDeliveryLineBatchDeleteAPIResponse() *TaobaoLogisticsDeliveryLineBatchDeleteAPIResponse {
	return poolTaobaoLogisticsDeliveryLineBatchDeleteAPIResponse.Get().(*TaobaoLogisticsDeliveryLineBatchDeleteAPIResponse)
}

// ReleaseTaobaoLogisticsDeliveryLineBatchDeleteAPIResponse 将 TaobaoLogisticsDeliveryLineBatchDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsDeliveryLineBatchDeleteAPIResponse(v *TaobaoLogisticsDeliveryLineBatchDeleteAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsDeliveryLineBatchDeleteAPIResponse.Put(v)
}
