package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangWarehouseCreateUpdateAPIRequest 仓库信息同步 API请求
// alibaba.dchain.aoxiang.warehouse.create.update
//
// 仓库信息同步
type AlibabaDchainAoxiangWarehouseCreateUpdateAPIRequest struct {
	model.Params
	// 仓信息同步入参
	_warehouseUpsertRequest *WarehouseUpsertRequest
}

// NewAlibabaDchainAoxiangWarehouseCreateUpdateRequest 初始化AlibabaDchainAoxiangWarehouseCreateUpdateAPIRequest对象
func NewAlibabaDchainAoxiangWarehouseCreateUpdateRequest() *AlibabaDchainAoxiangWarehouseCreateUpdateAPIRequest {
	return &AlibabaDchainAoxiangWarehouseCreateUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangWarehouseCreateUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.warehouse.create.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangWarehouseCreateUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetWarehouseUpsertRequest is WarehouseUpsertRequest Setter
// 仓信息同步入参
func (r *AlibabaDchainAoxiangWarehouseCreateUpdateAPIRequest) SetWarehouseUpsertRequest(_warehouseUpsertRequest *WarehouseUpsertRequest) error {
	r._warehouseUpsertRequest = _warehouseUpsertRequest
	r.Set("warehouse_upsert_request", _warehouseUpsertRequest)
	return nil
}

// GetWarehouseUpsertRequest WarehouseUpsertRequest Getter
func (r AlibabaDchainAoxiangWarehouseCreateUpdateAPIRequest) GetWarehouseUpsertRequest() *WarehouseUpsertRequest {
	return r._warehouseUpsertRequest
}
