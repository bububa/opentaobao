package ascp

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangWarehouseCreateUpdateAPIRequest) Reset() {
	r._warehouseUpsertRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangWarehouseCreateUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.warehouse.create.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangWarehouseCreateUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangWarehouseCreateUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaDchainAoxiangWarehouseCreateUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangWarehouseCreateUpdateRequest()
	},
}

// GetAlibabaDchainAoxiangWarehouseCreateUpdateRequest 从 sync.Pool 获取 AlibabaDchainAoxiangWarehouseCreateUpdateAPIRequest
func GetAlibabaDchainAoxiangWarehouseCreateUpdateAPIRequest() *AlibabaDchainAoxiangWarehouseCreateUpdateAPIRequest {
	return poolAlibabaDchainAoxiangWarehouseCreateUpdateAPIRequest.Get().(*AlibabaDchainAoxiangWarehouseCreateUpdateAPIRequest)
}

// ReleaseAlibabaDchainAoxiangWarehouseCreateUpdateAPIRequest 将 AlibabaDchainAoxiangWarehouseCreateUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangWarehouseCreateUpdateAPIRequest(v *AlibabaDchainAoxiangWarehouseCreateUpdateAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangWarehouseCreateUpdateAPIRequest.Put(v)
}
