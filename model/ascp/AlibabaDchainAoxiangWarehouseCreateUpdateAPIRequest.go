package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangwarehousecreateupdateAPIRequest 仓库信息同步 API请求
// alibaba.dchain.aoxiang.warehouse.create.update
//
// 仓库信息同步
type AlibabadchainaoxiangwarehousecreateupdateAPIRequest struct {
	model.Params
	// 仓信息同步入参
	_warehouseUpsertRequest *WarehouseUpsertRequest
}

// NewAlibabadchainaoxiangwarehousecreateupdateRequest 初始化AlibabadchainaoxiangwarehousecreateupdateAPIRequest对象
func NewAlibabadchainaoxiangwarehousecreateupdateRequest() *AlibabadchainaoxiangwarehousecreateupdateAPIRequest {
	return &AlibabadchainaoxiangwarehousecreateupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangwarehousecreateupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.warehouse.create.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangwarehousecreateupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangwarehousecreateupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseUpsertRequest is WarehouseUpsertRequest Setter
// 仓信息同步入参
func (r *AlibabadchainaoxiangwarehousecreateupdateAPIRequest) SetWarehouseUpsertRequest(_warehouseUpsertRequest *WarehouseUpsertRequest) error {
	r._warehouseUpsertRequest = _warehouseUpsertRequest
	r.Set("warehouse_upsert_request", _warehouseUpsertRequest)
	return nil
}

// GetWarehouseUpsertRequest WarehouseUpsertRequest Getter
func (r AlibabadchainaoxiangwarehousecreateupdateAPIRequest) GetWarehouseUpsertRequest() *WarehouseUpsertRequest {
	return r._warehouseUpsertRequest
}
