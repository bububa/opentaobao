package ascpffo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressascpinventorylogqueryAPIRequest AliExpress库存流水查询API API请求
// aliexpress.ascp.inventory.log.query
//
// AliExpress库存流水查询API
type AliexpressascpinventorylogqueryAPIRequest struct {
	model.Params
	// 查询DTO
	_warehouseInventoryLogQueryDto *WarehouseInventoryLogQueryDto
}

// NewAliexpressascpinventorylogqueryRequest 初始化AliexpressascpinventorylogqueryAPIRequest对象
func NewAliexpressascpinventorylogqueryRequest() *AliexpressascpinventorylogqueryAPIRequest {
	return &AliexpressascpinventorylogqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressascpinventorylogqueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.ascp.inventory.log.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressascpinventorylogqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressascpinventorylogqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseInventoryLogQueryDto is WarehouseInventoryLogQueryDto Setter
// 查询DTO
func (r *AliexpressascpinventorylogqueryAPIRequest) SetWarehouseInventoryLogQueryDto(_warehouseInventoryLogQueryDto *WarehouseInventoryLogQueryDto) error {
	r._warehouseInventoryLogQueryDto = _warehouseInventoryLogQueryDto
	r.Set("warehouse_inventory_log_query_dto", _warehouseInventoryLogQueryDto)
	return nil
}

// GetWarehouseInventoryLogQueryDto WarehouseInventoryLogQueryDto Getter
func (r AliexpressascpinventorylogqueryAPIRequest) GetWarehouseInventoryLogQueryDto() *WarehouseInventoryLogQueryDto {
	return r._warehouseInventoryLogQueryDto
}
