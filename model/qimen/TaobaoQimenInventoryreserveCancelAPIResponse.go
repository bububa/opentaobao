package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenInventoryreserveCancelAPIResponse 库存预占取消接口 API返回值
// taobao.qimen.inventoryreserve.cancel
//
// 库存预占取消
type TaobaoQimenInventoryreserveCancelAPIResponse struct {
	model.CommonResponse
	TaobaoQimenInventoryreserveCancelAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenInventoryreserveCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenInventoryreserveCancelAPIResponseModel).Reset()
}

// TaobaoQimenInventoryreserveCancelAPIResponseModel is 库存预占取消接口 成功返回结果
type TaobaoQimenInventoryreserveCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_inventoryreserve_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenInventoryreserveCancelResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenInventoryreserveCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenInventoryreserveCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenInventoryreserveCancelAPIResponse)
	},
}

// GetTaobaoQimenInventoryreserveCancelAPIResponse 从 sync.Pool 获取 TaobaoQimenInventoryreserveCancelAPIResponse
func GetTaobaoQimenInventoryreserveCancelAPIResponse() *TaobaoQimenInventoryreserveCancelAPIResponse {
	return poolTaobaoQimenInventoryreserveCancelAPIResponse.Get().(*TaobaoQimenInventoryreserveCancelAPIResponse)
}

// ReleaseTaobaoQimenInventoryreserveCancelAPIResponse 将 TaobaoQimenInventoryreserveCancelAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenInventoryreserveCancelAPIResponse(v *TaobaoQimenInventoryreserveCancelAPIResponse) {
	v.Reset()
	poolTaobaoQimenInventoryreserveCancelAPIResponse.Put(v)
}
