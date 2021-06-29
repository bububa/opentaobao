package alidoc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新药店 APIResponse
alibaba.alihealth.alidoc.drug.store.update

药店信息更新接口
*/
type AlibabaAlihealthAlidocDrugStoreUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthAlidocDrugStoreUpdateResponse
}

type AlibabaAlihealthAlidocDrugStoreUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_alidoc_drug_store_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // errMessage
    
    ErrMessage   string `json:"err_message,omitempty" xml:"err_message,omitempty"`

    
    // errCode
    
    ErrKode   string `json:"err_kode,omitempty" xml:"err_kode,omitempty"`

    
    // success
    
    Successed   bool `json:"successed,omitempty" xml:"successed,omitempty"`

    
}
