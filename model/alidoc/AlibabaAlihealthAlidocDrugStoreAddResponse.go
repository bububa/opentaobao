package alidoc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
gsk新增药店 APIResponse
alibaba.alihealth.alidoc.drug.store.add

GSK上传药店信息
*/
type AlibabaAlihealthAlidocDrugStoreAddAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthAlidocDrugStoreAddResponse
}

type AlibabaAlihealthAlidocDrugStoreAddResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_alidoc_drug_store_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // success
    
    Successed   bool `json:"successed,omitempty" xml:"successed,omitempty"`

    
    // errCode
    
    ErrorKode   string `json:"error_kode,omitempty" xml:"error_kode,omitempty"`

    
    // errMessage
    
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`

    
}
