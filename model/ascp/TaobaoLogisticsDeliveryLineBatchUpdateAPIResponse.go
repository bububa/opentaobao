package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsDeliveryLineBatchUpdateAPIResponse 线路能力创建/更新 API返回值
// taobao.logistics.delivery.line.batch.update
//
// 线路能力创建/更新
type TaobaoLogisticsDeliveryLineBatchUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsDeliveryLineBatchUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsDeliveryLineBatchUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsDeliveryLineBatchUpdateAPIResponseModel).Reset()
}

// TaobaoLogisticsDeliveryLineBatchUpdateAPIResponseModel is 线路能力创建/更新 成功返回结果
type TaobaoLogisticsDeliveryLineBatchUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_delivery_line_batch_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 线路能力创建/更新出参
	DeliveryLineBatchUpdateResponse *DeliveryLineBatchUpdateResponse `json:"delivery_line_batch_update_response,omitempty" xml:"delivery_line_batch_update_response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsDeliveryLineBatchUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DeliveryLineBatchUpdateResponse = nil
}

var poolTaobaoLogisticsDeliveryLineBatchUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsDeliveryLineBatchUpdateAPIResponse)
	},
}

// GetTaobaoLogisticsDeliveryLineBatchUpdateAPIResponse 从 sync.Pool 获取 TaobaoLogisticsDeliveryLineBatchUpdateAPIResponse
func GetTaobaoLogisticsDeliveryLineBatchUpdateAPIResponse() *TaobaoLogisticsDeliveryLineBatchUpdateAPIResponse {
	return poolTaobaoLogisticsDeliveryLineBatchUpdateAPIResponse.Get().(*TaobaoLogisticsDeliveryLineBatchUpdateAPIResponse)
}

// ReleaseTaobaoLogisticsDeliveryLineBatchUpdateAPIResponse 将 TaobaoLogisticsDeliveryLineBatchUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsDeliveryLineBatchUpdateAPIResponse(v *TaobaoLogisticsDeliveryLineBatchUpdateAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsDeliveryLineBatchUpdateAPIResponse.Put(v)
}
