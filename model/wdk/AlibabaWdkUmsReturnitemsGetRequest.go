package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退货库位商品查询（退货出库辅助）-回流单 APIRequest
alibaba.wdk.ums.returnitems.get

退货库位商品查询（退货出库辅助）-回流单
*/
type AlibabaWdkUmsReturnitemsGetRequest struct {
    model.Params

    // 店仓code，指的是作业对象，对应一个物理店或仓编码
    warehouseCode   string 

}

func NewAlibabaWdkUmsReturnitemsGetRequest() *AlibabaWdkUmsReturnitemsGetRequest{
    return &AlibabaWdkUmsReturnitemsGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkUmsReturnitemsGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.ums.returnitems.get"
}

func (r AlibabaWdkUmsReturnitemsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkUmsReturnitemsGetRequest) SetWarehouseCode(warehouseCode string) error {
    r.warehouseCode = warehouseCode
    r.Set("warehouse_code", warehouseCode)
    return nil
}

func (r AlibabaWdkUmsReturnitemsGetRequest) GetWarehouseCode() string {
    return r.warehouseCode
}

