package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
盘点结果单-回流单 APIRequest
alibaba.wdk.ums.inventory.check.get

盘点结果单-回流单
*/
type AlibabaWdkUmsInventoryCheckGetRequest struct {
    model.Params

    // 店仓code，指的是库调对象，对应一个物理店或仓编码
    warehouseCode   string 

}

func NewAlibabaWdkUmsInventoryCheckGetRequest() *AlibabaWdkUmsInventoryCheckGetRequest{
    return &AlibabaWdkUmsInventoryCheckGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkUmsInventoryCheckGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.ums.inventory.check.get"
}

func (r AlibabaWdkUmsInventoryCheckGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkUmsInventoryCheckGetRequest) SetWarehouseCode(warehouseCode string) error {
    r.warehouseCode = warehouseCode
    r.Set("warehouse_code", warehouseCode)
    return nil
}

func (r AlibabaWdkUmsInventoryCheckGetRequest) GetWarehouseCode() string {
    return r.warehouseCode
}

