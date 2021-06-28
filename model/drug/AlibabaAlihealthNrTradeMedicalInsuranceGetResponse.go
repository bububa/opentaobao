package drug

import (
    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康医保支付信息获取 APIResponse
alibaba.alihealth.nr.trade.medical.insurance.get

阿里健康医保支付信息获取
*/
type AlibabaAlihealthNrTradeMedicalInsuranceGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlihealthNrTradeMedicalInsuranceGetResponse `json:"alibaba_alihealth_nr_trade_medical_insurance_get_response,omitempty"` 
    AlibabaAlihealthNrTradeMedicalInsuranceGetResponse
}

/* model for simplify = false
type AlibabaAlihealthNrTradeMedicalInsuranceGetResponse struct {

    // 返回值总
    
    Result  *struct {
        ResponseResult  *ResponseResult `json:"response_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlihealthNrTradeMedicalInsuranceGetResponse struct {

    // 返回值总
    Result   *ResponseResult `json:"result,omitempty"`

}
