package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
VMOS回调行业信息系统 API返回值 
alibaba.jym.industry.information.callbak

VMOS回调交易猫行业信息系统
*/
type AlibabaJymIndustryInformationCallbakAPIResponse struct {
    model.CommonResponse
    AlibabaJymIndustryInformationCallbakAPIResponseModel
}

// VMOS回调行业信息系统 成功返回结果
type AlibabaJymIndustryInformationCallbakAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_jym_industry_information_callbak_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 状态码
    StateCode   string `json:"state_code,omitempty" xml:"state_code,omitempty"`
    // 扩展错误信息
    ExtraErrMsg   string `json:"extra_err_msg,omitempty" xml:"extra_err_msg,omitempty"`
}
