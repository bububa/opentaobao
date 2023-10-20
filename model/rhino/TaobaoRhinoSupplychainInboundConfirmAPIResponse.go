package rhino

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRhinoSupplychainInboundConfirmAPIResponse WMS003成衣入库确认 API返回值
// taobao.rhino.supplychain.inbound.confirm
//
// 【WMS003】【同步成衣入库完成信息】
type TaobaoRhinoSupplychainInboundConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoRhinoSupplychainInboundConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRhinoSupplychainInboundConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRhinoSupplychainInboundConfirmAPIResponseModel).Reset()
}

// TaobaoRhinoSupplychainInboundConfirmAPIResponseModel is WMS003成衣入库确认 成功返回结果
type TaobaoRhinoSupplychainInboundConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"rhino_supplychain_inbound_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRhinoSupplychainInboundConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.RetCode = 0
	m.IsSuccess = false
}

var poolTaobaoRhinoSupplychainInboundConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRhinoSupplychainInboundConfirmAPIResponse)
	},
}

// GetTaobaoRhinoSupplychainInboundConfirmAPIResponse 从 sync.Pool 获取 TaobaoRhinoSupplychainInboundConfirmAPIResponse
func GetTaobaoRhinoSupplychainInboundConfirmAPIResponse() *TaobaoRhinoSupplychainInboundConfirmAPIResponse {
	return poolTaobaoRhinoSupplychainInboundConfirmAPIResponse.Get().(*TaobaoRhinoSupplychainInboundConfirmAPIResponse)
}

// ReleaseTaobaoRhinoSupplychainInboundConfirmAPIResponse 将 TaobaoRhinoSupplychainInboundConfirmAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRhinoSupplychainInboundConfirmAPIResponse(v *TaobaoRhinoSupplychainInboundConfirmAPIResponse) {
	v.Reset()
	poolTaobaoRhinoSupplychainInboundConfirmAPIResponse.Put(v)
}
