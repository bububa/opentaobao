package pur

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
po反馈创建 APIResponse
alibaba.pur.supplier.porespcreate

PO反馈接口
*/
type AlibabaPurSupplierPorespcreateAPIResponse struct {
    model.CommonResponse
    AlibabaPurSupplierPorespcreateResponse
}

type AlibabaPurSupplierPorespcreateResponse struct {
    XMLName xml.Name `xml:"alibaba_pur_supplier_porespcreate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 获取url的出参
    
    Result   *ActionResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
