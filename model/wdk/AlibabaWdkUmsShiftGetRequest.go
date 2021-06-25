package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
移库单获取 APIRequest
alibaba.wdk.ums.shift.get

移库单获取
*/
type AlibabaWdkUmsShiftGetRequest struct {
    model.Params

    // 店仓code，指的是库调对象，对应一个物理店或仓编码
    warehouseCode   string 

}

func NewAlibabaWdkUmsShiftGetRequest() *AlibabaWdkUmsShiftGetRequest{
    return &AlibabaWdkUmsShiftGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkUmsShiftGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.ums.shift.get"
}

func (r AlibabaWdkUmsShiftGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkUmsShiftGetRequest) SetWarehouseCode(warehouseCode string) error {
    r.warehouseCode = warehouseCode
    r.Set("warehouse_code", warehouseCode)
    return nil
}

func (r AlibabaWdkUmsShiftGetRequest) GetWarehouseCode() string {
    return r.warehouseCode
}

