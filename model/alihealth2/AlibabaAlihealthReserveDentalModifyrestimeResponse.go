package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改预约时间 APIResponse
alibaba.alihealth.reserve.dental.modifyrestime

修改预约时间
*/
type AlibabaAlihealthReserveDentalModifyrestimeAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthReserveDentalModifyrestimeResponse
}

type AlibabaAlihealthReserveDentalModifyrestimeResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_reserve_dental_modifyrestime_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaAlihealthReserveDentalModifyrestimeResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
