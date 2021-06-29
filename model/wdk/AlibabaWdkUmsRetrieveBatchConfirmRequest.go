package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量消息确认 API请求
alibaba.wdk.ums.retrieve.batch.confirm

批量消息确认
*/
type AlibabaWdkUmsRetrieveBatchConfirmRequest struct {
    model.Params
    // warehouse_code
    warehouseCode   string
    // warehouse_code
    uuids   []string
}

// 初始化AlibabaWdkUmsRetrieveBatchConfirmRequest对象
func NewAlibabaWdkUmsRetrieveBatchConfirmRequest() *AlibabaWdkUmsRetrieveBatchConfirmRequest{
    return &AlibabaWdkUmsRetrieveBatchConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsRetrieveBatchConfirmRequest) GetApiMethodName() string {
    return "alibaba.wdk.ums.retrieve.batch.confirm"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsRetrieveBatchConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarehouseCode Setter
// warehouse_code
func (r *AlibabaWdkUmsRetrieveBatchConfirmRequest) SetWarehouseCode(warehouseCode string) error {
    r.warehouseCode = warehouseCode
    r.Set("warehouse_code", warehouseCode)
    return nil
}

// WarehouseCode Getter
func (r AlibabaWdkUmsRetrieveBatchConfirmRequest) GetWarehouseCode() string {
    return r.warehouseCode
}
// Uuids Setter
// warehouse_code
func (r *AlibabaWdkUmsRetrieveBatchConfirmRequest) SetUuids(uuids []string) error {
    r.uuids = uuids
    r.Set("uuids", uuids)
    return nil
}

// Uuids Getter
func (r AlibabaWdkUmsRetrieveBatchConfirmRequest) GetUuids() []string {
    return r.uuids
}
