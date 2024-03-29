package tmallcarenter

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarcenterVehicleCvmappingInsertAPIRequest EPC车辆版本信息与底盘信息库关系绑定 API请求
// tmall.carcenter.vehicle.cvmapping.insert
//
// EPC车辆版本信息与底盘信息库关系绑定
type TmallCarcenterVehicleCvmappingInsertAPIRequest struct {
	model.Params
	// 版本ID
	_supplierVersionCid string
	// 底盘ID
	_supplierChassisCid string
	// 状态
	_status int64
}

// NewTmallCarcenterVehicleCvmappingInsertRequest 初始化TmallCarcenterVehicleCvmappingInsertAPIRequest对象
func NewTmallCarcenterVehicleCvmappingInsertRequest() *TmallCarcenterVehicleCvmappingInsertAPIRequest {
	return &TmallCarcenterVehicleCvmappingInsertAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallCarcenterVehicleCvmappingInsertAPIRequest) Reset() {
	r._supplierVersionCid = ""
	r._supplierChassisCid = ""
	r._status = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarcenterVehicleCvmappingInsertAPIRequest) GetApiMethodName() string {
	return "tmall.carcenter.vehicle.cvmapping.insert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarcenterVehicleCvmappingInsertAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallCarcenterVehicleCvmappingInsertAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTmallCarcenterVehicleCvmappingInsertAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallCarcenterVehicleCvmappingInsertRequest()
	},
}

// GetTmallCarcenterVehicleCvmappingInsertRequest 从 sync.Pool 获取 TmallCarcenterVehicleCvmappingInsertAPIRequest
func GetTmallCarcenterVehicleCvmappingInsertAPIRequest() *TmallCarcenterVehicleCvmappingInsertAPIRequest {
	return poolTmallCarcenterVehicleCvmappingInsertAPIRequest.Get().(*TmallCarcenterVehicleCvmappingInsertAPIRequest)
}

// ReleaseTmallCarcenterVehicleCvmappingInsertAPIRequest 将 TmallCarcenterVehicleCvmappingInsertAPIRequest 放入 sync.Pool
func ReleaseTmallCarcenterVehicleCvmappingInsertAPIRequest(v *TmallCarcenterVehicleCvmappingInsertAPIRequest) {
	v.Reset()
	poolTmallCarcenterVehicleCvmappingInsertAPIRequest.Put(v)
}
