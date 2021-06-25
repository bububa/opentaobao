package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
回流单－外部对已拉取到的UMS单据进行确认 APIRequest
alibaba.wdk.ums.retrieve.confirm

回流单－外部对已拉取到的UMS单据进行确认
*/
type AlibabaWdkUmsRetrieveConfirmRequest struct {
    model.Params

    // 店仓code，指的是作业对象，对应一个物理店或仓编码
    warehouseCode   string 

    // 唯一识别码
    uuid   string 

}

func NewAlibabaWdkUmsRetrieveConfirmRequest() *AlibabaWdkUmsRetrieveConfirmRequest{
    return &AlibabaWdkUmsRetrieveConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkUmsRetrieveConfirmRequest) GetApiMethodName() string {
    return "alibaba.wdk.ums.retrieve.confirm"
}

func (r AlibabaWdkUmsRetrieveConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkUmsRetrieveConfirmRequest) SetWarehouseCode(warehouseCode string) error {
    r.warehouseCode = warehouseCode
    r.Set("warehouse_code", warehouseCode)
    return nil
}

func (r AlibabaWdkUmsRetrieveConfirmRequest) GetWarehouseCode() string {
    return r.warehouseCode
}

func (r *AlibabaWdkUmsRetrieveConfirmRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r AlibabaWdkUmsRetrieveConfirmRequest) GetUuid() string {
    return r.uuid
}

