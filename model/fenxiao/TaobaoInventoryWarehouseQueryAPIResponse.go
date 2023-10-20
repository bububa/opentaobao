package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoInventoryWarehouseQueryAPIResponse 分页查询商家仓信息 API返回值
// taobao.inventory.warehouse.query
//
// 分页查询商家仓信息
type TaobaoInventoryWarehouseQueryAPIResponse struct {
	model.CommonResponse
	TaobaoInventoryWarehouseQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoInventoryWarehouseQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoInventoryWarehouseQueryAPIResponseModel).Reset()
}

// TaobaoInventoryWarehouseQueryAPIResponseModel is 分页查询商家仓信息 成功返回结果
type TaobaoInventoryWarehouseQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"inventory_warehouse_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PaginationResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoInventoryWarehouseQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoInventoryWarehouseQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoInventoryWarehouseQueryAPIResponse)
	},
}

// GetTaobaoInventoryWarehouseQueryAPIResponse 从 sync.Pool 获取 TaobaoInventoryWarehouseQueryAPIResponse
func GetTaobaoInventoryWarehouseQueryAPIResponse() *TaobaoInventoryWarehouseQueryAPIResponse {
	return poolTaobaoInventoryWarehouseQueryAPIResponse.Get().(*TaobaoInventoryWarehouseQueryAPIResponse)
}

// ReleaseTaobaoInventoryWarehouseQueryAPIResponse 将 TaobaoInventoryWarehouseQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoInventoryWarehouseQueryAPIResponse(v *TaobaoInventoryWarehouseQueryAPIResponse) {
	v.Reset()
	poolTaobaoInventoryWarehouseQueryAPIResponse.Put(v)
}
