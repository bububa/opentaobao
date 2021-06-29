package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ERP调用打印面单取号接口 API请求
alibaba.ascp.uop.supplier.waybill.query

ERP调用打印面单取号接口
*/
type AlibabaAscpUopSupplierWaybillQueryRequest struct {
    model.Params
    // 查询面单请求参数
    waybillQueryRequest   *Waybillqueryrequest
}

// 初始化AlibabaAscpUopSupplierWaybillQueryRequest对象
func NewAlibabaAscpUopSupplierWaybillQueryRequest() *AlibabaAscpUopSupplierWaybillQueryRequest{
    return &AlibabaAscpUopSupplierWaybillQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopSupplierWaybillQueryRequest) GetApiMethodName() string {
    return "alibaba.ascp.uop.supplier.waybill.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopSupplierWaybillQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WaybillQueryRequest Setter
// 查询面单请求参数
func (r *AlibabaAscpUopSupplierWaybillQueryRequest) SetWaybillQueryRequest(waybillQueryRequest *Waybillqueryrequest) error {
    r.waybillQueryRequest = waybillQueryRequest
    r.Set("waybill_query_request", waybillQueryRequest)
    return nil
}

// WaybillQueryRequest Getter
func (r AlibabaAscpUopSupplierWaybillQueryRequest) GetWaybillQueryRequest() *Waybillqueryrequest {
    return r.waybillQueryRequest
}
