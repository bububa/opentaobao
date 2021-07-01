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
type AlibabaWdkUmsHandlingGetAPIRequest struct {
    model.Params
    // 仓库编码
    _warehouseCode   string
}

// 初始化AlibabaWdkUmsHandlingGetAPIRequest对象
func NewAlibabaWdkUmsHandlingGetRequest() *AlibabaWdkUmsHandlingGetAPIRequest{
    return &AlibabaWdkUmsHandlingGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsHandlingGetAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.ums.handling.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsHandlingGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarehouseCode Setter
// 仓库编码
func (r *AlibabaWdkUmsHandlingGetAPIRequest) SetWarehouseCode(_warehouseCode string) error {
    r._warehouseCode = _warehouseCode
    r.Set("warehouse_code", _warehouseCode)
    return nil
}

// WarehouseCode Getter
func (r AlibabaWdkUmsHandlingGetAPIRequest) GetWarehouseCode() string {
    return r._warehouseCode
}
