package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
库调单-回流单 APIRequest
alibaba.wdk.ums.inventory.adjust.get

库调单-回流单
*/
type AlibabaWdkUmsInventoryAdjustGetRequest struct {
    model.Params

    // 店仓code，指的是库调对象，对应一个物理店或仓编码
    warehouseCode   string 

}

func NewAlibabaWdkUmsInventoryAdjustGetRequest() *AlibabaWdkUmsInventoryAdjustGetRequest{
    return &AlibabaWdkUmsInventoryAdjustGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkUmsInventoryAdjustGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.ums.inventory.adjust.get"
}

func (r AlibabaWdkUmsInventoryAdjustGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkUmsInventoryAdjustGetRequest) SetWarehouseCode(warehouseCode string) error {
    r.warehouseCode = warehouseCode
    r.Set("warehouse_code", warehouseCode)
    return nil
}

func (r AlibabaWdkUmsInventoryAdjustGetRequest) GetWarehouseCode() string {
    return r.warehouseCode
}

