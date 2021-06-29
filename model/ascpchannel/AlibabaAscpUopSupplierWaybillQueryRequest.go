package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ERP调用打印面单取号接口 APIRequest
alibaba.ascp.uop.supplier.waybill.query

ERP调用打印面单取号接口
*/
type AlibabaAscpUopSupplierWaybillQueryRequest struct {
    model.Params

    // 查询面单请求参数
    waybillQueryRequest   *Waybillqueryrequest 

}

func NewAlibabaAscpUopSupplierWaybillQueryRequest() *AlibabaAscpUopSupplierWaybillQueryRequest{
    return &AlibabaAscpUopSupplierWaybillQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpUopSupplierWaybillQueryRequest) GetApiMethodName() string {
    return "alibaba.ascp.uop.supplier.waybill.query"
}

func (r AlibabaAscpUopSupplierWaybillQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpUopSupplierWaybillQueryRequest) SetWaybillQueryRequest(waybillQueryRequest *Waybillqueryrequest) error {
    r.waybillQueryRequest = waybillQueryRequest
    r.Set("waybill_query_request", waybillQueryRequest)
    return nil
}

func (r AlibabaAscpUopSupplierWaybillQueryRequest) GetWaybillQueryRequest() *Waybillqueryrequest {
    return r.waybillQueryRequest
}

