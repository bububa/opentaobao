package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卡巴数据同步 APIRequest
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

func NewAlibabaCampusGuardDataSyncRequest() *AlibabaCampusGuardDataSyncRequest{
    return &AlibabaCampusGuardDataSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusGuardDataSyncRequest) GetApiMethodName() string {
    return "alibaba.campus.guard.data.sync"
}

func (r AlibabaCampusGuardDataSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusGuardDataSyncRequest) SetDataType(dataType string) error {
    r.dataType = dataType
    r.Set("data_type", dataType)
    return nil
}

func (r AlibabaCampusGuardDataSyncRequest) GetDataType() string {
    return r.dataType
}

func (r *AlibabaCampusGuardDataSyncRequest) SetSupplierName(supplierName string) error {
    r.supplierName = supplierName
    r.Set("supplier_name", supplierName)
    return nil
}

func (r AlibabaCampusGuardDataSyncRequest) GetSupplierName() string {
    return r.supplierName
}

func (r *AlibabaCampusGuardDataSyncRequest) SetData(data string) error {
    r.data = data
    r.Set("data", data)
    return nil
}

func (r AlibabaCampusGuardDataSyncRequest) GetData() string {
    return r.data
}

