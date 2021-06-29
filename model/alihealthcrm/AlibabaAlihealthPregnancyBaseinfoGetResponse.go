package alihealthcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
拉取备孕初始化信息 APIResponse
alibaba.alihealth.pregnancy.baseinfo.get

拉取备孕初始化信息
*/
type AlibabaAlihealthPregnancyBaseinfoGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthPregnancyBaseinfoGetResponse
}

type AlibabaAlihealthPregnancyBaseinfoGetResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_pregnancy_baseinfo_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果集
    
    Result   *AlibabaAlihealthPregnancyBaseinfoGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
