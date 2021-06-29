package pur

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
采购供应商订单明细查询接口 APIResponse
alibaba.ceres.supplier.po.querydetail

采购供应商订单明细查询接口
*/
type AlibabaCeresSupplierPoQuerydetailAPIResponse struct {
    model.CommonResponse
    AlibabaCeresSupplierPoQuerydetailResponse
}

type AlibabaCeresSupplierPoQuerydetailResponse struct {
    XMLName xml.Name `xml:"alibaba_ceres_supplier_po_querydetail_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回消息体
    
    Result   *SupplierPoDetailDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
