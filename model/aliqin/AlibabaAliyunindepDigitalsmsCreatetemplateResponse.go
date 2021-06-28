package aliqin

import (
    "github.com/bububa/opentaobao/model"
)

/* 
数字短信模板创建 APIResponse
alibaba.aliyunindep.digitalsms.createtemplate

数字短信模板创建，给阿里云一方产品使用，类型：9
*/
type AlibabaAliyunindepDigitalsmsCreatetemplateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAliyunindepDigitalsmsCreatetemplateResponse `json:"alibaba_aliyunindep_digitalsms_createtemplate_response,omitempty"` 
    AlibabaAliyunindepDigitalsmsCreatetemplateResponse
}

/* model for simplify = false
type AlibabaAliyunindepDigitalsmsCreatetemplateResponse struct {

    // 返回值
    
    Result  *struct {
        AlibabaAliyunindepDigitalsmsCreatetemplateResult  *AlibabaAliyunindepDigitalsmsCreatetemplateResult `json:"alibaba_aliyunindep_digitalsms_createtemplate_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAliyunindepDigitalsmsCreatetemplateResponse struct {

    // 返回值
    Result   *AlibabaAliyunindepDigitalsmsCreatetemplateResult `json:"result,omitempty"`

}
