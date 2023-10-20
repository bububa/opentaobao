package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenInventoryQueryAPIResponse 库存查询接口（多商品） API返回值
// taobao.qimen.inventory.query
//
// taobao.qimen.inventory.query
type TaobaoQimenInventoryQueryAPIResponse struct {
	model.CommonResponse
	TaobaoQimenInventoryQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenInventoryQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenInventoryQueryAPIResponseModel).Reset()
}

// TaobaoQimenInventoryQueryAPIResponseModel is 库存查询接口（多商品） 成功返回结果
type TaobaoQimenInventoryQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_inventory_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *InventoryQueryResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenInventoryQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenInventoryQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenInventoryQueryAPIResponse)
	},
}

// GetTaobaoQimenInventoryQueryAPIResponse 从 sync.Pool 获取 TaobaoQimenInventoryQueryAPIResponse
func GetTaobaoQimenInventoryQueryAPIResponse() *TaobaoQimenInventoryQueryAPIResponse {
	return poolTaobaoQimenInventoryQueryAPIResponse.Get().(*TaobaoQimenInventoryQueryAPIResponse)
}

// ReleaseTaobaoQimenInventoryQueryAPIResponse 将 TaobaoQimenInventoryQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenInventoryQueryAPIResponse(v *TaobaoQimenInventoryQueryAPIResponse) {
	v.Reset()
	poolTaobaoQimenInventoryQueryAPIResponse.Put(v)
}
