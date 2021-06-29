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
    dataType   string
    // 供应商名称
    supplierName   string
    // json串
    data   string
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
func (r *AlibabaCampusGuardDataSyncRequest) SetDataType(dataType string) error {
    r.dataType = dataType
    r.Set("data_type", dataType)
    return nil
}

// DataType Getter
func (r AlibabaCampusGuardDataSyncRequest) GetDataType() string {
    return r.dataType
}
// SupplierName Setter
// 供应商名称
func (r *AlibabaCampusGuardDataSyncRequest) SetSupplierName(supplierName string) error {
    r.supplierName = supplierName
    r.Set("supplier_name", supplierName)
    return nil
}

// SupplierName Getter
func (r AlibabaCampusGuardDataSyncRequest) GetSupplierName() string {
    return r.supplierName
}
// Data Setter
// json串
func (r *AlibabaCampusGuardDataSyncRequest) SetData(data string) error {
    r.data = data
    r.Set("data", data)
    return nil
}

// Data Getter
func (r AlibabaCampusGuardDataSyncRequest) GetData() string {
    return r.data
}
