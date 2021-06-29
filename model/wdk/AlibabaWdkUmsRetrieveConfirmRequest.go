package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
回流单－外部对已拉取到的UMS单据进行确认 API请求
alibaba.wdk.ums.retrieve.confirm

回流单－外部对已拉取到的UMS单据进行确认
*/
type AlibabaWdkUmsRetrieveConfirmRequest struct {
    model.Params
    // 店仓code，指的是作业对象，对应一个物理店或仓编码
    _warehouseCode   string
    // 唯一识别码
    _uuid   string
}

// 初始化AlibabaWdkUmsRetrieveConfirmRequest对象
func NewAlibabaWdkUmsRetrieveConfirmRequest() *AlibabaWdkUmsRetrieveConfirmRequest{
    return &AlibabaWdkUmsRetrieveConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsRetrieveConfirmRequest) GetApiMethodName() string {
    return "alibaba.wdk.ums.retrieve.confirm"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsRetrieveConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarehouseCode Setter
// 店仓code，指的是作业对象，对应一个物理店或仓编码
func (r *AlibabaWdkUmsRetrieveConfirmRequest) SetWarehouseCode(_warehouseCode string) error {
    r._warehouseCode = _warehouseCode
    r.Set("warehouse_code", _warehouseCode)
    return nil
}

// WarehouseCode Getter
func (r AlibabaWdkUmsRetrieveConfirmRequest) GetWarehouseCode() string {
    return r._warehouseCode
}
// Uuid Setter
// 唯一识别码
func (r *AlibabaWdkUmsRetrieveConfirmRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r AlibabaWdkUmsRetrieveConfirmRequest) GetUuid() string {
    return r._uuid
}
