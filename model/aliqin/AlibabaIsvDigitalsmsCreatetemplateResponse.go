package aliqin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
数字短信模板创建 APIResponse
alibaba.isv.digitalsms.createtemplate

数字短信模板创建，给聚石塔，类型：2
*/
type AlibabaIsvDigitalsmsCreatetemplateAPIResponse struct {
    model.CommonResponse
    AlibabaIsvDigitalsmsCreatetemplateResponse
}

type AlibabaIsvDigitalsmsCreatetemplateResponse struct {
    XMLName xml.Name `xml:"alibaba_isv_digitalsms_createtemplate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回值
    
    Result   *AlibabaIsvDigitalsmsCreatetemplateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
