package tmallcarenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
EPC车辆版本信息与底盘信息库关系绑定 APIRequest
tmall.carcenter.vehicle.cvmapping.insert

EPC车辆版本信息与底盘信息库关系绑定
*/
type TmallCarcenterVehicleCvmappingInsertRequest struct {
    model.Params

    // 状态
    status   int64 

    // 版本ID
    supplierVersionCid   string 

    // 底盘ID
    supplierChassisCid   string 

}

func NewTmallCarcenterVehicleCvmappingInsertRequest() *TmallCarcenterVehicleCvmappingInsertRequest{
    return &TmallCarcenterVehicleCvmappingInsertRequest{
        Params: model.NewParams(),
    }
}

func (r TmallCarcenterVehicleCvmappingInsertRequest) GetApiMethodName() string {
    return "tmall.carcenter.vehicle.cvmapping.insert"
}

func (r TmallCarcenterVehicleCvmappingInsertRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallCarcenterVehicleCvmappingInsertRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TmallCarcenterVehicleCvmappingInsertRequest) GetStatus() int64 {
    return r.status
}

func (r *TmallCarcenterVehicleCvmappingInsertRequest) SetSupplierVersionCid(supplierVersionCid string) error {
    r.supplierVersionCid = supplierVersionCid
    r.Set("supplier_version_cid", supplierVersionCid)
    return nil
}

func (r TmallCarcenterVehicleCvmappingInsertRequest) GetSupplierVersionCid() string {
    return r.supplierVersionCid
}

func (r *TmallCarcenterVehicleCvmappingInsertRequest) SetSupplierChassisCid(supplierChassisCid string) error {
    r.supplierChassisCid = supplierChassisCid
    r.Set("supplier_chassis_cid", supplierChassisCid)
    return nil
}

func (r TmallCarcenterVehicleCvmappingInsertRequest) GetSupplierChassisCid() string {
    return r.supplierChassisCid
}

