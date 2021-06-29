package aedropshiper

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提供给买家使用的运费计算接口 APIResponse
aliexpress.logistics.buyer.freight.calculate

提供给买家使用的运费计算接口
*/
type AliexpressLogisticsBuyerFreightCalculateAPIResponse struct {
    model.CommonResponse
    AliexpressLogisticsBuyerFreightCalculateResponse
}

type AliexpressLogisticsBuyerFreightCalculateResponse struct {
    XMLName xml.Name `xml:"aliexpress_logistics_buyer_freight_calculate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AeopFreightCalculateResultListResponseForBuyer `json:"result,omitempty" xml:"result,omitempty"`

    
}
