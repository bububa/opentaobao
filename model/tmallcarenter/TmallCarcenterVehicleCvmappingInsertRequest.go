package tmallcarenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
EPC车辆版本信息与底盘信息库关系绑定 API请求
tmall.carcenter.vehicle.cvmapping.insert

EPC车辆版本信息与底盘信息库关系绑定
*/
type TmallCarcenterVehicleCvmappingInsertRequest struct {
    model.Params
    // 状态
    _status   int64
    // 版本ID
    _supplierVersionCid   string
    // 底盘ID
    _supplierChassisCid   string
}

// 初始化TmallCarcenterVehicleCvmappingInsertRequest对象
func NewTmallCarcenterVehicleCvmappingInsertRequest() *TmallCarcenterVehicleCvmappingInsertRequest{
    return &TmallCarcenterVehicleCvmappingInsertRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarcenterVehicleCvmappingInsertRequest) GetApiMethodName() string {
    return "tmall.carcenter.vehicle.cvmapping.insert"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarcenterVehicleCvmappingInsertRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Status Setter
// 状态
func (r *TmallCarcenterVehicleCvmappingInsertRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TmallCarcenterVehicleCvmappingInsertRequest) GetStatus() int64 {
    return r._status
}
// SupplierVersionCid Setter
// 版本ID
func (r *TmallCarcenterVehicleCvmappingInsertRequest) SetSupplierVersionCid(_supplierVersionCid string) error {
    r._supplierVersionCid = _supplierVersionCid
    r.Set("supplier_version_cid", _supplierVersionCid)
    return nil
}

// SupplierVersionCid Getter
func (r TmallCarcenterVehicleCvmappingInsertRequest) GetSupplierVersionCid() string {
    return r._supplierVersionCid
}
// SupplierChassisCid Setter
// 底盘ID
func (r *TmallCarcenterVehicleCvmappingInsertRequest) SetSupplierChassisCid(_supplierChassisCid string) error {
    r._supplierChassisCid = _supplierChassisCid
    r.Set("supplier_chassis_cid", _supplierChassisCid)
    return nil
}

// SupplierChassisCid Getter
func (r TmallCarcenterVehicleCvmappingInsertRequest) GetSupplierChassisCid() string {
    return r._supplierChassisCid
}
