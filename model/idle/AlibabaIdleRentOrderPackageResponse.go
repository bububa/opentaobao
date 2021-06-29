package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
确认揽收商品 APIResponse
alibaba.idle.rent.order.package

确认揽收
*/
type AlibabaIdleRentOrderPackageAPIResponse struct {
    model.CommonResponse
    AlibabaIdleRentOrderPackageResponse
}

type AlibabaIdleRentOrderPackageResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_rent_order_package_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 系统自动生成
    
    Result   *TopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
