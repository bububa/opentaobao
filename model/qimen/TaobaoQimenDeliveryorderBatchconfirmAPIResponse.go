package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenDeliveryorderBatchconfirmAPIResponse 发货单确认接口 API返回值
// taobao.qimen.deliveryorder.batchconfirm
//
// taobao.qimen.deliveryorder.batchconfirm
type TaobaoQimenDeliveryorderBatchconfirmAPIResponse struct {
	model.CommonResponse
	TaobaoQimenDeliveryorderBatchconfirmAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenDeliveryorderBatchconfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenDeliveryorderBatchconfirmAPIResponseModel).Reset()
}

// TaobaoQimenDeliveryorderBatchconfirmAPIResponseModel is 发货单确认接口 成功返回结果
type TaobaoQimenDeliveryorderBatchconfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_deliveryorder_batchconfirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenDeliveryorderBatchconfirmResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenDeliveryorderBatchconfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenDeliveryorderBatchconfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenDeliveryorderBatchconfirmAPIResponse)
	},
}

// GetTaobaoQimenDeliveryorderBatchconfirmAPIResponse 从 sync.Pool 获取 TaobaoQimenDeliveryorderBatchconfirmAPIResponse
func GetTaobaoQimenDeliveryorderBatchconfirmAPIResponse() *TaobaoQimenDeliveryorderBatchconfirmAPIResponse {
	return poolTaobaoQimenDeliveryorderBatchconfirmAPIResponse.Get().(*TaobaoQimenDeliveryorderBatchconfirmAPIResponse)
}

// ReleaseTaobaoQimenDeliveryorderBatchconfirmAPIResponse 将 TaobaoQimenDeliveryorderBatchconfirmAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenDeliveryorderBatchconfirmAPIResponse(v *TaobaoQimenDeliveryorderBatchconfirmAPIResponse) {
	v.Reset()
	poolTaobaoQimenDeliveryorderBatchconfirmAPIResponse.Put(v)
}
