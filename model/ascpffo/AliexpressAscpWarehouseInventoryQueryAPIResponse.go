package ascpffo

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAscpWarehouseInventoryQueryAPIResponse AliExpress在仓库存查询API API返回值
// aliexpress.ascp.warehouse.inventory.query
//
// AliExpress在仓库存查询查询API
type AliexpressAscpWarehouseInventoryQueryAPIResponse struct {
	model.CommonResponse
	AliexpressAscpWarehouseInventoryQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressAscpWarehouseInventoryQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressAscpWarehouseInventoryQueryAPIResponseModel).Reset()
}

// AliexpressAscpWarehouseInventoryQueryAPIResponseModel is AliExpress在仓库存查询API 成功返回结果
type AliexpressAscpWarehouseInventoryQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_ascp_warehouse_inventory_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分页查询结果
	Result *PageQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressAscpWarehouseInventoryQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressAscpWarehouseInventoryQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressAscpWarehouseInventoryQueryAPIResponse)
	},
}

// GetAliexpressAscpWarehouseInventoryQueryAPIResponse 从 sync.Pool 获取 AliexpressAscpWarehouseInventoryQueryAPIResponse
func GetAliexpressAscpWarehouseInventoryQueryAPIResponse() *AliexpressAscpWarehouseInventoryQueryAPIResponse {
	return poolAliexpressAscpWarehouseInventoryQueryAPIResponse.Get().(*AliexpressAscpWarehouseInventoryQueryAPIResponse)
}

// ReleaseAliexpressAscpWarehouseInventoryQueryAPIResponse 将 AliexpressAscpWarehouseInventoryQueryAPIResponse 保存到 sync.Pool
func ReleaseAliexpressAscpWarehouseInventoryQueryAPIResponse(v *AliexpressAscpWarehouseInventoryQueryAPIResponse) {
	v.Reset()
	poolAliexpressAscpWarehouseInventoryQueryAPIResponse.Put(v)
}
