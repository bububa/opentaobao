package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退货库位商品查询（退货出库辅助）-回流单 API请求
alibaba.wdk.ums.returnitems.get

退货库位商品查询（退货出库辅助）-回流单
*/
type AlibabaWdkUmsReturnitemsGetRequest struct {
    model.Params
    // 店仓code，指的是作业对象，对应一个物理店或仓编码
    _warehouseCode   string
}

// 初始化AlibabaWdkUmsReturnitemsGetRequest对象
func NewAlibabaWdkUmsReturnitemsGetRequest() *AlibabaWdkUmsReturnitemsGetRequest{
    return &AlibabaWdkUmsReturnitemsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsReturnitemsGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.ums.returnitems.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsReturnitemsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarehouseCode Setter
// 店仓code，指的是作业对象，对应一个物理店或仓编码
func (r *AlibabaWdkUmsReturnitemsGetRequest) SetWarehouseCode(_warehouseCode string) error {
    r._warehouseCode = _warehouseCode
    r.Set("warehouse_code", _warehouseCode)
    return nil
}

// WarehouseCode Getter
func (r AlibabaWdkUmsReturnitemsGetRequest) GetWarehouseCode() string {
    return r._warehouseCode
}
