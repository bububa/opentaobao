package rhino

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRhinoSupplychainOutboundPickingcompleteAPIResponse 【WMS005】接收成衣捡配完成通知 API返回值
// taobao.rhino.supplychain.outbound.pickingcomplete
//
// 接收成衣捡配完成通知,WMS005
type TaobaoRhinoSupplychainOutboundPickingcompleteAPIResponse struct {
	model.CommonResponse
	TaobaoRhinoSupplychainOutboundPickingcompleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRhinoSupplychainOutboundPickingcompleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRhinoSupplychainOutboundPickingcompleteAPIResponseModel).Reset()
}

// TaobaoRhinoSupplychainOutboundPickingcompleteAPIResponseModel is 【WMS005】接收成衣捡配完成通知 成功返回结果
type TaobaoRhinoSupplychainOutboundPickingcompleteAPIResponseModel struct {
	XMLName xml.Name `xml:"rhino_supplychain_outbound_pickingcomplete_response"`
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
func (m *TaobaoRhinoSupplychainOutboundPickingcompleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.RetCode = 0
	m.IsSuccess = false
}

var poolTaobaoRhinoSupplychainOutboundPickingcompleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRhinoSupplychainOutboundPickingcompleteAPIResponse)
	},
}

// GetTaobaoRhinoSupplychainOutboundPickingcompleteAPIResponse 从 sync.Pool 获取 TaobaoRhinoSupplychainOutboundPickingcompleteAPIResponse
func GetTaobaoRhinoSupplychainOutboundPickingcompleteAPIResponse() *TaobaoRhinoSupplychainOutboundPickingcompleteAPIResponse {
	return poolTaobaoRhinoSupplychainOutboundPickingcompleteAPIResponse.Get().(*TaobaoRhinoSupplychainOutboundPickingcompleteAPIResponse)
}

// ReleaseTaobaoRhinoSupplychainOutboundPickingcompleteAPIResponse 将 TaobaoRhinoSupplychainOutboundPickingcompleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRhinoSupplychainOutboundPickingcompleteAPIResponse(v *TaobaoRhinoSupplychainOutboundPickingcompleteAPIResponse) {
	v.Reset()
	poolTaobaoRhinoSupplychainOutboundPickingcompleteAPIResponse.Put(v)
}
