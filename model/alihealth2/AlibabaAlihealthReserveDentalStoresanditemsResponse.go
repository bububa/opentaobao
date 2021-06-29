package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商户门店，商品列表 APIResponse
alibaba.alihealth.reserve.dental.storesanditems

查询商户门店，商品列表
*/
type AlibabaAlihealthReserveDentalStoresanditemsAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthReserveDentalStoresanditemsResponse
}

type AlibabaAlihealthReserveDentalStoresanditemsResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_reserve_dental_storesanditems_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
