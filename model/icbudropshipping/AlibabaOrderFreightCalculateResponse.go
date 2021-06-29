package icbudropshipping

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里巴巴下单场景运费方案计算 APIResponse
alibaba.order.freight.calculate

icbu开展 drop shipping 业务，阿里巴巴下单场景运费方案计算
alibaba Create order scenario freight calculation
*/
type AlibabaOrderFreightCalculateAPIResponse struct {
    model.CommonResponse
    AlibabaOrderFreightCalculateResponse
}

type AlibabaOrderFreightCalculateResponse struct {
    XMLName xml.Name `xml:"alibaba_order_freight_calculate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // Logistics Solution List
    
    Value   []LogisticsSolution `json:"value,omitempty" xml:"value>logistics_solution,omitempty"`
    
    
}
