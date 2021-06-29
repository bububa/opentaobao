package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
移库单获取 API请求
alibaba.wdk.ums.shift.get

移库单获取
*/
type AlibabaWdkUmsShiftGetRequest struct {
    model.Params
    // 店仓code，指的是库调对象，对应一个物理店或仓编码
    warehouseCode   string
}

// 初始化AlibabaWdkUmsShiftGetRequest对象
func NewAlibabaWdkUmsShiftGetRequest() *AlibabaWdkUmsShiftGetRequest{
    return &AlibabaWdkUmsShiftGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsShiftGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.ums.shift.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsShiftGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarehouseCode Setter
// 店仓code，指的是库调对象，对应一个物理店或仓编码
func (r *AlibabaWdkUmsShiftGetRequest) SetWarehouseCode(warehouseCode string) error {
    r.warehouseCode = warehouseCode
    r.Set("warehouse_code", warehouseCode)
    return nil
}

// WarehouseCode Getter
func (r AlibabaWdkUmsShiftGetRequest) GetWarehouseCode() string {
    return r.warehouseCode
}
