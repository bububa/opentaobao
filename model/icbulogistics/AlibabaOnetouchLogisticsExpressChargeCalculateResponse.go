package icbulogistics

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
计算快递运费&下单参数校验 APIResponse
alibaba.onetouch.logistics.express.charge.calculate

计算快递运费、下单参数校验
*/
type AlibabaOnetouchLogisticsExpressChargeCalculateAPIResponse struct {
    model.CommonResponse
    AlibabaOnetouchLogisticsExpressChargeCalculateResponse
}

type AlibabaOnetouchLogisticsExpressChargeCalculateResponse struct {
    XMLName xml.Name `xml:"alibaba_onetouch_logistics_express_charge_calculate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaOnetouchLogisticsExpressChargeCalculateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
