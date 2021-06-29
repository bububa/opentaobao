package drug

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康医保支付信息获取 APIResponse
alibaba.alihealth.nr.trade.medical.insurance.get

阿里健康医保支付信息获取
*/
type AlibabaAlihealthNrTradeMedicalInsuranceGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthNrTradeMedicalInsuranceGetResponse
}

type AlibabaAlihealthNrTradeMedicalInsuranceGetResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_nr_trade_medical_insurance_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回值总
    
    Result   *ResponseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
