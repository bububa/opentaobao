package aliqin

import (
    "github.com/bububa/opentaobao/model"
)

/* 
数字短信模板创建 APIResponse
alibaba.isv.digitalsms.createtemplate

数字短信模板创建，给聚石塔，类型：2
*/
type AlibabaIsvDigitalsmsCreatetemplateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaIsvDigitalsmsCreatetemplateResponse `json:"alibaba_isv_digitalsms_createtemplate_response,omitempty"` 
    AlibabaIsvDigitalsmsCreatetemplateResponse
}

/* model for simplify = false
type AlibabaIsvDigitalsmsCreatetemplateResponse struct {

    // 返回值
    
    Result  *struct {
        AlibabaIsvDigitalsmsCreatetemplateResult  *AlibabaIsvDigitalsmsCreatetemplateResult `json:"alibaba_isv_digitalsms_createtemplate_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaIsvDigitalsmsCreatetemplateResponse struct {

    // 返回值
    Result   *AlibabaIsvDigitalsmsCreatetemplateResult `json:"result,omitempty"`

}
