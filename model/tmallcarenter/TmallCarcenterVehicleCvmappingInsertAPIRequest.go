package tmallcarenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallCarcenterVehicleCvmappingInsertAPIRequest EPC车辆版本信息与底盘信息库关系绑定 API请求
// tmall.carcenter.vehicle.cvmapping.insert
//
// EPC车辆版本信息与底盘信息库关系绑定
type TmallCarcenterVehicleCvmappingInsertAPIRequest struct {
	model.Params
	// 状态
	_status int64
	// 版本ID
	_supplierVersionCid string
	// 底盘ID
	_supplierChassisCid string
}

// NewTmallCarcenterVehicleCvmappingInsertRequest 初始化TmallCarcenterVehicleCvmappingInsertAPIRequest对象
func NewTmallCarcenterVehicleCvmappingInsertRequest() *TmallCarcenterVehicleCvmappingInsertAPIRequest {
	return &TmallCarcenterVehicleCvmappingInsertAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarcenterVehicleCvmappingInsertAPIRequest) GetApiMethodName() string {
	return "tmall.carcenter.vehicle.cvmapping.insert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarcenterVehicleCvmappingInsertAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetStatus is Status Setter
// 状态
func (r *TmallCarcenterVehicleCvmappingInsertAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TmallCarcenterVehicleCvmappingInsertAPIRequest) GetStatus() int64 {
	return r._status
}

// SetSupplierVersionCid is SupplierVersionCid Setter
// 版本ID
func (r *TmallCarcenterVehicleCvmappingInsertAPIRequest) SetSupplierVersionCid(_supplierVersionCid string) error {
	r._supplierVersionCid = _supplierVersionCid
	r.Set("supplier_version_cid", _supplierVersionCid)
	return nil
}

// GetSupplierVersionCid SupplierVersionCid Getter
func (r TmallCarcenterVehicleCvmappingInsertAPIRequest) GetSupplierVersionCid() string {
	return r._supplierVersionCid
}

// SetSupplierChassisCid is SupplierChassisCid Setter
// 底盘ID
func (r *TmallCarcenterVehicleCvmappingInsertAPIRequest) SetSupplierChassisCid(_supplierChassisCid string) error {
	r._supplierChassisCid = _supplierChassisCid
	r.Set("supplier_chassis_cid", _supplierChassisCid)
	return nil
}

// GetSupplierChassisCid SupplierChassisCid Getter
func (r TmallCarcenterVehicleCvmappingInsertAPIRequest) GetSupplierChassisCid() string {
	return r._supplierChassisCid
}
