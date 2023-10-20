package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenDeliveryorderBatchcreateAnswerAPIResponse 发货单创建结果通知接口(批量) API返回值
// taobao.qimen.deliveryorder.batchcreate.answer
//
// WMS调用接口，用于异步化的批量发货单创建结果通知。（如菜鸟发货单批量创建结果的返回）
type TaobaoQimenDeliveryorderBatchcreateAnswerAPIResponse struct {
	model.CommonResponse
	TaobaoQimenDeliveryorderBatchcreateAnswerAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenDeliveryorderBatchcreateAnswerAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenDeliveryorderBatchcreateAnswerAPIResponseModel).Reset()
}

// TaobaoQimenDeliveryorderBatchcreateAnswerAPIResponseModel is 发货单创建结果通知接口(批量) 成功返回结果
type TaobaoQimenDeliveryorderBatchcreateAnswerAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_deliveryorder_batchcreate_answer_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenDeliveryorderBatchcreateAnswerResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenDeliveryorderBatchcreateAnswerAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenDeliveryorderBatchcreateAnswerAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenDeliveryorderBatchcreateAnswerAPIResponse)
	},
}

// GetTaobaoQimenDeliveryorderBatchcreateAnswerAPIResponse 从 sync.Pool 获取 TaobaoQimenDeliveryorderBatchcreateAnswerAPIResponse
func GetTaobaoQimenDeliveryorderBatchcreateAnswerAPIResponse() *TaobaoQimenDeliveryorderBatchcreateAnswerAPIResponse {
	return poolTaobaoQimenDeliveryorderBatchcreateAnswerAPIResponse.Get().(*TaobaoQimenDeliveryorderBatchcreateAnswerAPIResponse)
}

// ReleaseTaobaoQimenDeliveryorderBatchcreateAnswerAPIResponse 将 TaobaoQimenDeliveryorderBatchcreateAnswerAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenDeliveryorderBatchcreateAnswerAPIResponse(v *TaobaoQimenDeliveryorderBatchcreateAnswerAPIResponse) {
	v.Reset()
	poolTaobaoQimenDeliveryorderBatchcreateAnswerAPIResponse.Put(v)
}
