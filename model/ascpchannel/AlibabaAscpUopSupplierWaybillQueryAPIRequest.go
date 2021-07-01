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
type AlibabaAscpUopSupplierWaybillQueryAPIRequest struct {
    model.Params
    // 查询面单请求参数
    _waybillQueryRequest   *Waybillqueryrequest
}

// 初始化AlibabaAscpUopSupplierWaybillQueryAPIRequest对象
func NewAlibabaAscpUopSupplierWaybillQueryRequest() *AlibabaAscpUopSupplierWaybillQueryAPIRequest{
    return &AlibabaAscpUopSupplierWaybillQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopSupplierWaybillQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.ascp.uop.supplier.waybill.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopSupplierWaybillQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WaybillQueryRequest Setter
// 查询面单请求参数
func (r *AlibabaAscpUopSupplierWaybillQueryAPIRequest) SetWaybillQueryRequest(_waybillQueryRequest *Waybillqueryrequest) error {
    r._waybillQueryRequest = _waybillQueryRequest
    r.Set("waybill_query_request", _waybillQueryRequest)
    return nil
}

// WaybillQueryRequest Getter
func (r AlibabaAscpUopSupplierWaybillQueryAPIRequest) GetWaybillQueryRequest() *Waybillqueryrequest {
    return r._waybillQueryRequest
}
