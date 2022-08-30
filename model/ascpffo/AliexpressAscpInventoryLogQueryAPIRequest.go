package ascpffo

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressAscpInventoryLogQueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.ascp.inventory.log.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressAscpInventoryLogQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
