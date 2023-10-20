package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangWarehouseStatusUpdateAPIRequest 启用/停用仓资源 API请求
// alibaba.dchain.aoxiang.warehouse.status.update
//
// 启用/停用仓资源
type AlibabaDchainAoxiangWarehouseStatusUpdateAPIRequest struct {
	model.Params
	// 启用/停用仓资源入参
	_warehouseStatusUpdateRequest *WarehouseStatusUpdateRequest
}

// NewAlibabaDchainAoxiangWarehouseStatusUpdateRequest 初始化AlibabaDchainAoxiangWarehouseStatusUpdateAPIRequest对象
func NewAlibabaDchainAoxiangWarehouseStatusUpdateRequest() *AlibabaDchainAoxiangWarehouseStatusUpdateAPIRequest {
	return &AlibabaDchainAoxiangWarehouseStatusUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangWarehouseStatusUpdateAPIRequest) Reset() {
	r._warehouseStatusUpdateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangWarehouseStatusUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.warehouse.status.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangWarehouseStatusUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangWarehouseStatusUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseStatusUpdateRequest is WarehouseStatusUpdateRequest Setter
// 启用/停用仓资源入参
func (r *AlibabaDchainAoxiangWarehouseStatusUpdateAPIRequest) SetWarehouseStatusUpdateRequest(_warehouseStatusUpdateRequest *WarehouseStatusUpdateRequest) error {
	r._warehouseStatusUpdateRequest = _warehouseStatusUpdateRequest
	r.Set("warehouse_status_update_request", _warehouseStatusUpdateRequest)
	return nil
}

// GetWarehouseStatusUpdateRequest WarehouseStatusUpdateRequest Getter
func (r AlibabaDchainAoxiangWarehouseStatusUpdateAPIRequest) GetWarehouseStatusUpdateRequest() *WarehouseStatusUpdateRequest {
	return r._warehouseStatusUpdateRequest
}

var poolAlibabaDchainAoxiangWarehouseStatusUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangWarehouseStatusUpdateRequest()
	},
}

// GetAlibabaDchainAoxiangWarehouseStatusUpdateRequest 从 sync.Pool 获取 AlibabaDchainAoxiangWarehouseStatusUpdateAPIRequest
func GetAlibabaDchainAoxiangWarehouseStatusUpdateAPIRequest() *AlibabaDchainAoxiangWarehouseStatusUpdateAPIRequest {
	return poolAlibabaDchainAoxiangWarehouseStatusUpdateAPIRequest.Get().(*AlibabaDchainAoxiangWarehouseStatusUpdateAPIRequest)
}

// ReleaseAlibabaDchainAoxiangWarehouseStatusUpdateAPIRequest 将 AlibabaDchainAoxiangWarehouseStatusUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangWarehouseStatusUpdateAPIRequest(v *AlibabaDchainAoxiangWarehouseStatusUpdateAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangWarehouseStatusUpdateAPIRequest.Put(v)
}
