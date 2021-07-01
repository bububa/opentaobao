package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV新增/修改口腔门店 API返回值 
alibaba.alihealth.dental.store.insertorupdate

ISV新增/修改口腔门店
*/
type AlibabaAlihealthDentalStoreInsertorupdateAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDentalStoreInsertorupdateAPIResponseModel
}

// ISV新增/修改口腔门店 成功返回结果
type AlibabaAlihealthDentalStoreInsertorupdateAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_dental_store_insertorupdate_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Result   *AlibabaAlihealthDentalStoreInsertorupdateMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}
