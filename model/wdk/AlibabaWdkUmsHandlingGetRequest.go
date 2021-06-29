package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
加工单-回流单（新接口） API请求
alibaba.wdk.ums.handling.get

加工单-回流单（新接口）
*/
type AlibabaWdkUmsHandlingGetRequest struct {
    model.Params
    // 仓库编码
    _warehouseCode   string
}

// 初始化AlibabaWdkUmsHandlingGetRequest对象
func NewAlibabaWdkUmsHandlingGetRequest() *AlibabaWdkUmsHandlingGetRequest{
    return &AlibabaWdkUmsHandlingGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsHandlingGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.ums.handling.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsHandlingGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarehouseCode Setter
// 仓库编码
func (r *AlibabaWdkUmsHandlingGetRequest) SetWarehouseCode(_warehouseCode string) error {
    r._warehouseCode = _warehouseCode
    r.Set("warehouse_code", _warehouseCode)
    return nil
}

// WarehouseCode Getter
func (r AlibabaWdkUmsHandlingGetRequest) GetWarehouseCode() string {
    return r._warehouseCode
}
