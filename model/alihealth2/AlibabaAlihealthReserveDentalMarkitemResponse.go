package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
标记商品是否可预约 APIResponse
alibaba.alihealth.reserve.dental.markitem

标记商品是否可预约
*/
type AlibabaAlihealthReserveDentalMarkitemAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthReserveDentalMarkitemResponse
}

type AlibabaAlihealthReserveDentalMarkitemResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_reserve_dental_markitem_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
