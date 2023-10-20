package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenDeliveryorderBatchcreateAPIResponse 发货单创建批量接口 API返回值
// taobao.qimen.deliveryorder.batchcreate
//
// taobao.qimen.deliveryorder.batchcreate
type TaobaoQimenDeliveryorderBatchcreateAPIResponse struct {
	model.CommonResponse
	TaobaoQimenDeliveryorderBatchcreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenDeliveryorderBatchcreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenDeliveryorderBatchcreateAPIResponseModel).Reset()
}

// TaobaoQimenDeliveryorderBatchcreateAPIResponseModel is 发货单创建批量接口 成功返回结果
type TaobaoQimenDeliveryorderBatchcreateAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_deliveryorder_batchcreate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *DeliveryOrderBatchCreateResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenDeliveryorderBatchcreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenDeliveryorderBatchcreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenDeliveryorderBatchcreateAPIResponse)
	},
}

// GetTaobaoQimenDeliveryorderBatchcreateAPIResponse 从 sync.Pool 获取 TaobaoQimenDeliveryorderBatchcreateAPIResponse
func GetTaobaoQimenDeliveryorderBatchcreateAPIResponse() *TaobaoQimenDeliveryorderBatchcreateAPIResponse {
	return poolTaobaoQimenDeliveryorderBatchcreateAPIResponse.Get().(*TaobaoQimenDeliveryorderBatchcreateAPIResponse)
}

// ReleaseTaobaoQimenDeliveryorderBatchcreateAPIResponse 将 TaobaoQimenDeliveryorderBatchcreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenDeliveryorderBatchcreateAPIResponse(v *TaobaoQimenDeliveryorderBatchcreateAPIResponse) {
	v.Reset()
	poolTaobaoQimenDeliveryorderBatchcreateAPIResponse.Put(v)
}
