package tmallcarenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcarcentervehiclecvmappinginsertAPIRequest EPC车辆版本信息与底盘信息库关系绑定 API请求
// tmall.carcenter.vehicle.cvmapping.insert
//
// EPC车辆版本信息与底盘信息库关系绑定
type TmallcarcentervehiclecvmappinginsertAPIRequest struct {
	model.Params
	// 版本ID
	_supplierVersionCid string
	// 底盘ID
	_supplierChassisCid string
	// 状态
	_status int64
}

// NewTmallcarcentervehiclecvmappinginsertRequest 初始化TmallcarcentervehiclecvmappinginsertAPIRequest对象
func NewTmallcarcentervehiclecvmappinginsertRequest() *TmallcarcentervehiclecvmappinginsertAPIRequest {
	return &TmallcarcentervehiclecvmappinginsertAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcarcentervehiclecvmappinginsertAPIRequest) GetApiMethodName() string {
	return "tmall.carcenter.vehicle.cvmapping.insert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcarcentervehiclecvmappinginsertAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcarcentervehiclecvmappinginsertAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSupplierVersionCid is SupplierVersionCid Setter
// 版本ID
func (r *TmallcarcentervehiclecvmappinginsertAPIRequest) SetSupplierVersionCid(_supplierVersionCid string) error {
	r._supplierVersionCid = _supplierVersionCid
	r.Set("supplier_version_cid", _supplierVersionCid)
	return nil
}

// GetSupplierVersionCid SupplierVersionCid Getter
func (r TmallcarcentervehiclecvmappinginsertAPIRequest) GetSupplierVersionCid() string {
	return r._supplierVersionCid
}

// SetSupplierChassisCid is SupplierChassisCid Setter
// 底盘ID
func (r *TmallcarcentervehiclecvmappinginsertAPIRequest) SetSupplierChassisCid(_supplierChassisCid string) error {
	r._supplierChassisCid = _supplierChassisCid
	r.Set("supplier_chassis_cid", _supplierChassisCid)
	return nil
}

// GetSupplierChassisCid SupplierChassisCid Getter
func (r TmallcarcentervehiclecvmappinginsertAPIRequest) GetSupplierChassisCid() string {
	return r._supplierChassisCid
}

// SetStatus is Status Setter
// 状态
func (r *TmallcarcentervehiclecvmappinginsertAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TmallcarcentervehiclecvmappinginsertAPIRequest) GetStatus() int64 {
	return r._status
}
