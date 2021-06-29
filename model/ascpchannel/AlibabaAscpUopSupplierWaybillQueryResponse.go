package ascpchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ERP调用打印面单取号接口 API返回值 
alibaba.ascp.uop.supplier.waybill.query

ERP调用打印面单取号接口
*/
type AlibabaAscpUopSupplierWaybillQueryAPIResponse struct {
    model.CommonResponse
    AlibabaAscpUopSupplierWaybillQueryResponse
}

// ERP调用打印面单取号接口 成功返回结果
type AlibabaAscpUopSupplierWaybillQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_ascp_uop_supplier_waybill_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回值包装,result为返回具体消息内容
    WaybillQueryResponse   *ResultWrapper `json:"waybill_query_response,omitempty" xml:"waybill_query_response,omitempty"`
}
