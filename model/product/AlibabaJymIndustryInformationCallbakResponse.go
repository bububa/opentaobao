package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
VMOS回调行业信息系统 APIResponse
alibaba.jym.industry.information.callbak

VMOS回调交易猫行业信息系统
*/
type AlibabaJymIndustryInformationCallbakAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaJymIndustryInformationCallbakResponse `json:"alibaba_jym_industry_information_callbak_response,omitempty"` 
    AlibabaJymIndustryInformationCallbakResponse
}

/* model for simplify = false
type AlibabaJymIndustryInformationCallbakResponse struct {

    // 状态码
    
    StateCode   string `json:"state_code,omitempty"`
    

    // 扩展错误信息
    
    ExtraErrMsg   string `json:"extra_err_msg,omitempty"`
    

}
*/

type AlibabaJymIndustryInformationCallbakResponse struct {

    // 状态码
    StateCode   string `json:"state_code,omitempty"`

    // 扩展错误信息
    ExtraErrMsg   string `json:"extra_err_msg,omitempty"`

}
