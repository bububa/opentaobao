package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
绑定门店信息，商品信息 APIResponse
alibaba.alihealth.reserve.dental.bindshopanditem

绑定门店信息，商品信息
*/
type AlibabaAlihealthReserveDentalBindshopanditemAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthReserveDentalBindshopanditemResponse
}

type AlibabaAlihealthReserveDentalBindshopanditemResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_reserve_dental_bindshopanditem_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaAlihealthReserveDentalBindshopanditemResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
