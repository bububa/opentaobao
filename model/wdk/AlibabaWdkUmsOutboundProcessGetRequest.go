package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
出库业务UMS异步处理结果返回 API请求
alibaba.wdk.ums.outbound.process.get

出库业务UMS异步处理结果返回
*/
type AlibabaWdkUmsOutboundProcessGetRequest struct {
    model.Params
    // 店仓code，指的是作业对象，对应一个物理店或仓编码
    warehouseCode   string
}

// 初始化AlibabaWdkUmsOutboundProcessGetRequest对象
func NewAlibabaWdkUmsOutboundProcessGetRequest() *AlibabaWdkUmsOutboundProcessGetRequest{
    return &AlibabaWdkUmsOutboundProcessGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsOutboundProcessGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.ums.outbound.process.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsOutboundProcessGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarehouseCode Setter
// 店仓code，指的是作业对象，对应一个物理店或仓编码
func (r *AlibabaWdkUmsOutboundProcessGetRequest) SetWarehouseCode(warehouseCode string) error {
    r.warehouseCode = warehouseCode
    r.Set("warehouse_code", warehouseCode)
    return nil
}

// WarehouseCode Getter
func (r AlibabaWdkUmsOutboundProcessGetRequest) GetWarehouseCode() string {
    return r.warehouseCode
}
