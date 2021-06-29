package pur

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
asn创建 APIResponse
alibaba.pur.supplier.asncreate

asn创建
*/
type AlibabaPurSupplierAsncreateAPIResponse struct {
    model.CommonResponse
    AlibabaPurSupplierAsncreateResponse
}

type AlibabaPurSupplierAsncreateResponse struct {
    XMLName xml.Name `xml:"alibaba_pur_supplier_asncreate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 获取url的出参
    
    Result   *ActionResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
