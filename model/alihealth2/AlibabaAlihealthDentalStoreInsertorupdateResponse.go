package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV新增/修改口腔门店 APIResponse
alibaba.alihealth.dental.store.insertorupdate

ISV新增/修改口腔门店
*/
type AlibabaAlihealthDentalStoreInsertorupdateAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDentalStoreInsertorupdateResponse
}

type AlibabaAlihealthDentalStoreInsertorupdateResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_dental_store_insertorupdate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *AlibabaAlihealthDentalStoreInsertorupdateMtopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
