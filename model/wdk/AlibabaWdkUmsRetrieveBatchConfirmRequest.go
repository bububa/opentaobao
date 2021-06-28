package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量消息确认 APIRequest
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

func NewAlibabaWdkUmsRetrieveBatchConfirmRequest() *AlibabaWdkUmsRetrieveBatchConfirmRequest{
    return &AlibabaWdkUmsRetrieveBatchConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkUmsRetrieveBatchConfirmRequest) GetApiMethodName() string {
    return "alibaba.wdk.ums.retrieve.batch.confirm"
}

func (r AlibabaWdkUmsRetrieveBatchConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkUmsRetrieveBatchConfirmRequest) SetWarehouseCode(warehouseCode string) error {
    r.warehouseCode = warehouseCode
    r.Set("warehouse_code", warehouseCode)
    return nil
}

func (r AlibabaWdkUmsRetrieveBatchConfirmRequest) GetWarehouseCode() string {
    return r.warehouseCode
}

func (r *AlibabaWdkUmsRetrieveBatchConfirmRequest) SetUuids(uuids []string) error {
    r.uuids = uuids
    r.Set("uuids", uuids)
    return nil
}

func (r AlibabaWdkUmsRetrieveBatchConfirmRequest) GetUuids() []string {
    return r.uuids
}

