package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卡巴数据同步 API请求
alibaba.campus.guard.data.sync

数据同步门禁系统
*/
type AlibabaCampusGuardDataSyncRequest struct {
    model.Params
    // 1-刷卡流水
    _dataType   string
    // 供应商名称
    _supplierName   string
    // json串
    _data   string
}

// 初始化AlibabaCampusGuardDataSyncRequest对象
func NewAlibabaCampusGuardDataSyncRequest() *AlibabaCampusGuardDataSyncRequest{
    return &AlibabaCampusGuardDataSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusGuardDataSyncRequest) GetApiMethodName() string {
    return "alibaba.campus.guard.data.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusGuardDataSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DataType Setter
// 1-刷卡流水
func (r *AlibabaCampusGuardDataSyncRequest) SetDataType(_dataType string) error {
    r._dataType = _dataType
    r.Set("data_type", _dataType)
    return nil
}

// DataType Getter
func (r AlibabaCampusGuardDataSyncRequest) GetDataType() string {
    return r._dataType
}
// SupplierName Setter
// 供应商名称
func (r *AlibabaCampusGuardDataSyncRequest) SetSupplierName(_supplierName string) error {
    r._supplierName = _supplierName
    r.Set("supplier_name", _supplierName)
    return nil
}

// SupplierName Getter
func (r AlibabaCampusGuardDataSyncRequest) GetSupplierName() string {
    return r._supplierName
}
// Data Setter
// json串
func (r *AlibabaCampusGuardDataSyncRequest) SetData(_data string) error {
    r._data = _data
    r.Set("data", _data)
    return nil
}

// Data Getter
func (r AlibabaCampusGuardDataSyncRequest) GetData() string {
    return r._data
}
