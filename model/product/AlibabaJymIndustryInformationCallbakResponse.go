package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
VMOS回调行业信息系统 APIResponse
alibaba.jym.industry.information.callbak

VMOS回调交易猫行业信息系统
*/
type AlibabaJymIndustryInformationCallbakAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_jym_industry_information_callbak_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 状态码
    
    StateCode   string `json:"state_code,omitempty" xml:"