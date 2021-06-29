package legalcase

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取通用枚举值接口 APIResponse
alibaba.legal.case.common.enumdata

获取通用枚举值接口
*/
type AlibabaLegalCaseCommonEnumdataAPIResponse struct {
    model.CommonResponse
    AlibabaLegalCaseCommonEnumdataResponse
}

type AlibabaLegalCaseCommonEnumdataResponse struct {
    XMLName xml.Name `xml:"alibaba_legal_case_common_enumdata_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
