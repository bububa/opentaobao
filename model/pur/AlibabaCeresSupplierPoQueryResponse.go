package pur

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
采购供应商订单查询接口 APIResponse
alibaba.ceres.supplier.po.query

采购供应商订单查询接口
*/
type AlibabaCeresSupplierPoQueryAPIResponse struct {
    model.CommonResponse
    AlibabaCeresSupplierPoQueryResponse
}

type AlibabaCeresSupplierPoQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_ceres_supplier_po_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回消息体
    
    Result   *AlibabaCeresSupplierPoQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
