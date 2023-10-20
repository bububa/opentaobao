package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenDeliveryorderConfirmAPIResponse 发货单确认接口 API返回值
// taobao.qimen.deliveryorder.confirm
//
// taobao.qimen.deliveryorder.confirm
type TaobaoQimenDeliveryorderConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoQimenDeliveryorderConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenDeliveryorderConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenDeliveryorderConfirmAPIResponseModel).Reset()
}

// TaobaoQimenDeliveryorderConfirmAPIResponseModel is 发货单确认接口 成功返回结果
type TaobaoQimenDeliveryorderConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_deliveryorder_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenDeliveryorderConfirmResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenDeliveryorderConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenDeliveryorderConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenDeliveryorderConfirmAPIResponse)
	},
}

// GetTaobaoQimenDeliveryorderConfirmAPIResponse 从 sync.Pool 获取 TaobaoQimenDeliveryorderConfirmAPIResponse
func GetTaobaoQimenDeliveryorderConfirmAPIResponse() *TaobaoQimenDeliveryorderConfirmAPIResponse {
	return poolTaobaoQimenDeliveryorderConfirmAPIResponse.Get().(*TaobaoQimenDeliveryorderConfirmAPIResponse)
}

// ReleaseTaobaoQimenDeliveryorderConfirmAPIResponse 将 TaobaoQimenDeliveryorderConfirmAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenDeliveryorderConfirmAPIResponse(v *TaobaoQimenDeliveryorderConfirmAPIResponse) {
	v.Reset()
	poolTaobaoQimenDeliveryorderConfirmAPIResponse.Put(v)
}
