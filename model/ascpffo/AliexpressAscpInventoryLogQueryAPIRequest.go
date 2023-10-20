package ascpffo

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAscpInventoryLogQueryAPIRequest AliExpress库存流水查询API API请求
// aliexpress.ascp.inventory.log.query
//
// AliExpress库存流水查询API
type AliexpressAscpInventoryLogQueryAPIRequest struct {
	model.Params
	// 查询DTO
	_warehouseInventoryLogQueryDto *WarehouseInventoryLogQueryDto
}

// NewAliexpressAscpInventoryLogQueryRequest 初始化AliexpressAscpInventoryLogQueryAPIRequest对象
func NewAliexpressAscpInventoryLogQueryRequest() *AliexpressAscpInventoryLogQueryAPIRequest {
	return &AliexpressAscpInventoryLogQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressAscpInventoryLogQueryAPIRequest) Reset() {
	r._warehouseInventoryLogQueryDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressAscpInventoryLogQueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.ascp.inventory.log.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressAscpInventoryLogQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressAscpInventoryLogQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseInventoryLogQueryDto is WarehouseInventoryLogQueryDto Setter
// 查询DTO
func (r *AliexpressAscpInventoryLogQueryAPIRequest) SetWarehouseInventoryLogQueryDto(_warehouseInventoryLogQueryDto *WarehouseInventoryLogQueryDto) error {
	r._warehouseInventoryLogQueryDto = _warehouseInventoryLogQueryDto
	r.Set("warehouse_inventory_log_query_dto", _warehouseInventoryLogQueryDto)
	return nil
}

// GetWarehouseInventoryLogQueryDto WarehouseInventoryLogQueryDto Getter
func (r AliexpressAscpInventoryLogQueryAPIRequest) GetWarehouseInventoryLogQueryDto() *WarehouseInventoryLogQueryDto {
	return r._warehouseInventoryLogQueryDto
}

var poolAliexpressAscpInventoryLogQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressAscpInventoryLogQueryRequest()
	},
}

// GetAliexpressAscpInventoryLogQueryRequest 从 sync.Pool 获取 AliexpressAscpInventoryLogQueryAPIRequest
func GetAliexpressAscpInventoryLogQueryAPIRequest() *AliexpressAscpInventoryLogQueryAPIRequest {
	return poolAliexpressAscpInventoryLogQueryAPIRequest.Get().(*AliexpressAscpInventoryLogQueryAPIRequest)
}

// ReleaseAliexpressAscpInventoryLogQueryAPIRequest 将 AliexpressAscpInventoryLogQueryAPIRequest 放入 sync.Pool
func ReleaseAliexpressAscpInventoryLogQueryAPIRequest(v *AliexpressAscpInventoryLogQueryAPIRequest) {
	v.Reset()
	poolAliexpressAscpInventoryLogQueryAPIRequest.Put(v)
}
